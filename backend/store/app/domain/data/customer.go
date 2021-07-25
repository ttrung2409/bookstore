package data

type Customer struct {
	Id              EntityId `gorm:"primaryKey"`
	Name            string
	Phone           string
	DeliveryAddress string
}

func (c *Customer) GetId() EntityId {
	return c.Id
}

func (c *Customer) SetId(id EntityId) {
	c.Id = id
}
