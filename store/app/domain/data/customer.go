package data

type Customer struct {
	Id              string `gorm:"primaryKey"`
	Name            string
	Phone           string
	DeliveryAddress string
}

func (c *Customer) GetId() string {
	return c.Id
}

func (c *Customer) SetId(id string) {
	c.Id = id
}
