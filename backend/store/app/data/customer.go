package data

type Customer struct {
	Id              EntityId `gorm:"primaryKey"`
	Name            string
	Phone           string
	DeliveryAddress string
}
