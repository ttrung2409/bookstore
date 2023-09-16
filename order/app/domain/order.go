package domain

import (
	"fmt"
	"time"

	"github.com/thoas/go-funk"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusAccepted  OrderStatus = "Accepted"
	OrderStatusRejected  OrderStatus = "Rejected"
	OrderStatusCancelled OrderStatus = "Cancelled"
	OrderStatusDelivered OrderStatus = "Delivered"
)

type OrderData struct {
	Id                      string
	CreatedAt               time.Time
	Status                  OrderStatus
	CustomerId              string
	CustomerName            string
	CustomerPhone           string
	CustomerDeliveryAddress string
	Items                   []OrderItem
}

type OrderItem struct {
	OrderId          string
	BookId           string
	BookTitle        string
	BookSubtitle     string
	BookDescription  string
	BookThumbnailUrl string
	Qty              int
}

type Order struct {
	EventSource
	order OrderData
}

func (Order) New(customer Customer, books []Book) *Order {
	customerData := customer.State()

	return &Order{
		order: OrderData{
			Status:                  OrderStatusPending,
			CreatedAt:               time.Now(),
			CustomerId:              customerData.Id,
			CustomerName:            customerData.Name,
			CustomerPhone:           customerData.Phone,
			CustomerDeliveryAddress: customerData.DeliveryAddress,
			Items: funk.Map(books, func(book Book) OrderItem {
				bookData := book.State()

				return OrderItem{
					BookId:           bookData.Id,
					BookTitle:        bookData.Title,
					BookDescription:  bookData.Description,
					BookSubtitle:     bookData.Subtitle,
					BookThumbnailUrl: bookData.ThumbnailUrl,
				}
			}).([]OrderItem),
		},
	}
}

func (order *Order) Cancel() error {
	if order.order.Status != OrderStatusPending &&
		order.order.Status != OrderStatusAccepted {
		return fmt.Errorf(
			"order status is '%s', no cancellation allowed",
			order.order.Status,
		)
	}

	return nil
}
