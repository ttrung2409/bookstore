package domain

import repository "store/repository/interface"

type Customer struct {
	Id              repository.EntityId
	Name            string
	Phone           string
	DeliveryAddress string
}