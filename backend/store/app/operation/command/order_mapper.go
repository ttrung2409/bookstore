package operation

import "store/app/data"

func (Order) fromDataObject(order data.Order) Order {
	items := []OrderItem{}
	for _, dataItem := range order.Items {
		items = append(items, OrderItem{}.fromDataObject(dataItem))
	}

	return Order{
		Id:         order.Id.ToString(),
		Number:     order.Number,
		CreatedAt:  order.CreatedAt,
		CustomerId: order.CustomerId.ToString(),
		Status:     string(order.Status),
		Items:      items,
	}
}

func (OrderItem) fromDataObject(item data.OrderItem) OrderItem {
	return OrderItem{
		Book: Book{}.fromDataObject(item.Book),
		Qty:  item.Qty,
	}
}
