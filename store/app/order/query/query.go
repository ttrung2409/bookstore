package query

import (
	"store/app/domain"
	repo "store/repository"
	"time"

	"github.com/thoas/go-funk"
)

type OrderListItem = struct {
	Id        string
	CreatedAt time.Time
	Customer  Customer
	Status    domain.OrderStatus
}

type Query interface {
	FindDeliverableOrders() ([]OrderListItem, error)
	GetOrderDetails(id string) (Order, error)
}

func New() Query {
	return &query{}
}

type query struct{}

func (*query) FindDeliverableOrders() ([]OrderListItem, error) {
	orders, err := repo.Query[domain.OrderData]{}.New().
		Select(
			"id",
			"created_at",
			"customer_id",
			"customer_name",
			"customer_phone",
			"customer_delivery_address",
			"status").
		Where("status = ?", domain.OrderStatusAccepted).
		Find()

	if err != nil {
		return nil, err
	}

	return funk.Map(orders, func(order domain.OrderData) OrderListItem {
		return OrderListItem{
			Id:        order.Id,
			CreatedAt: order.CreatedAt,
			Customer: Customer{
				Name:            order.CustomerName,
				Phone:           order.CustomerPhone,
				DeliveryAddress: order.CustomerDeliveryAddress,
			},
			Status: order.Status,
		}
	}).([]OrderListItem), nil
}

func (*query) GetOrderDetails(id string) (Order, error) {
	record, err := repo.Query[domain.OrderData]{}.New().
		Join("Customer").
		Preload("Items").
		Join("Items.Book").
		Where("id = ?", id).
		FindOne()

	if err != nil {
		return Order{}, err
	}

	order := Order{}.fromDataObject(record)

	return order, nil
}
