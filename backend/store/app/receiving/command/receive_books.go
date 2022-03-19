package command

import (
	"store/app/domain"
	"store/app/domain/data"
	"store/app/messaging"
	"store/app/messaging/event"
	repo "store/app/repository"
	"store/container"
	"store/utils"

	funk "github.com/thoas/go-funk"
)

type ReceivingBook struct {
	Book
	Qty int
	id  string
}

type ReceiveBooksRequest struct {
	Items []*ReceivingBook
}

func (*command) Receive(request ReceiveBooksRequest) error {
	bookRepository := container.Instance().Get(utils.Nameof((*repo.BookRepository)(nil))).(repo.BookRepository)
	bookReceiptRepository := container.Instance().Get(utils.Nameof((*repo.BookReceiptRepository)(nil))).(repo.BookReceiptRepository)
	transactionFactory := container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
	producer := container.Instance().Get(utils.Nameof((*messaging.Producer)(nil))).(messaging.Producer)

	events := make([]interface{}, 0)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			// create books if not exists
			for _, item := range request.Items {
				bookId, isNew, err := bookRepository.CreateIfNotExists(
					domain.Book{}.New(item.Book.toDataObject()),
					tx,
				)

				if err != nil {
					return nil, err
				}

				item.id = bookId

				if isNew {
					events = append(
						events,
						&event.BookCreated{
							Id:            bookId,
							Title:         item.Title,
							Subtitle:      item.Subtitle,
							Description:   item.Description,
							Authors:       item.Authors,
							Publisher:     item.Publisher,
							PublishedDate: item.PublishedDate,
							AverageRating: item.AverageRating,
							RatingsCount:  item.RatingsCount,
							ThumbnailUrl:  item.ThumbnailUrl,
							PreviewUrl:    item.PreviewUrl,
						},
					)
				}
			}

			events = append(
				events,
				&event.StockChanged{
					Items: funk.Map(request.Items, func(item *ReceivingBook) *event.StockChangedItem {
						return &event.StockChangedItem{BookId: item.id}
					}).([]*event.StockChangedItem),
				},
			)

			// create book receipt
			newReceipt := domain.BookReceipt{}.NewFromReceivingBooks(funk.Map(
				request.Items,
				func(item *ReceivingBook) *domain.ReceivingBook {
					return &domain.ReceivingBook{Book: item.toDataObject(), ReceivingQty: item.Qty}
				},
			).([]*domain.ReceivingBook))

			receiptId, err := bookReceiptRepository.Create(newReceipt, tx)
			if err != nil {
				return nil, err
			}

			return receiptId, err
		})

	if err == nil {
		producer.Send(events...)
	}

	channel := make(chan error)

	go updateOrdersToStockFilled(channel)

	return err
}

// Update order status to StockFilled for any orders
// that can be fulfilled by the new stock
func updateOrdersToStockFilled(channel chan error) {
	defer func() {
		close(channel)
	}()

	orderRepository := container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)
	producer := container.Instance().Get(utils.Nameof((*messaging.Producer)(nil))).(messaging.Producer)

	orders, err := orderRepository.GetReceivingOrders(nil)
	if err != nil {
		channel <- err
		return
	}

	for _, order := range orders {
		if ok, _ := order.UpdateToStockFilled(); ok {
			err = orderRepository.Update(order, nil)

			if err == nil {
				producer.Send(
					&event.OrderStatusChanged{OrderId: order.State().Id, Status: data.OrderStatusStockFilled},
				)
			}
		}
	}
}
