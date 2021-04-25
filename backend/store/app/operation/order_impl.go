package operation

import (
	"store/app/domain"
)

func (Order) fromDomainObject(o domain.Order) Order {
	return Order{
		Id:         o.Id.ToString(),
		Number:     o.Number,
		CreatedAt:  o.CreatedAt,
		CustomerId: o.CustomerId.ToString(),
		Status:     int(o.Status),
	}
}
