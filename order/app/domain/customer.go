package domain

type CustomerData struct {
	Id              string
	Name            string
	Phone           string
	DeliveryAddress string
}

func (c *CustomerData) Clone() CustomerData {
	return CustomerData{Id: c.Id, Name: c.Name, Phone: c.Phone, DeliveryAddress: c.DeliveryAddress}
}

type Customer struct {
	customer CustomerData
}

func (Customer) New(customer CustomerData) *Customer {
	cloned := customer.Clone()
	if cloned.Id == "" {
		cloned.Id = NewId()
	}

	return &Customer{customer: cloned}
}

func (customer *Customer) State() CustomerData {
	return customer.customer.Clone()
}
