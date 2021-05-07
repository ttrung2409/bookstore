package operation

import "store/app/data"

func (Order) fromDataObject(o data.Order) Order {
	return Order{
		Id:         o.Id.ToString(),
		Number:     o.Number,
		CreatedAt:  o.CreatedAt,
		CustomerId: o.CustomerId.ToString(),
		Status:     string(o.Status),
	}
}
