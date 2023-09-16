package command

import "ecommerce/app/domain/data"

type Customer struct {
	Name            string
	Phone           string
	DeliveryAddress string
}

func (c *Customer) toDataObject() data.Customer {
	return data.Customer{Name: c.Name, Phone: c.Phone, DeliveryAddress: c.DeliveryAddress}
}
