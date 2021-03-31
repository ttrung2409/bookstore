package domain

import (
	repo "store/repository/interface"
	"time"
)

type OrderStatus int

const (
	Queued OrderStatus = iota
	Accepted OrderStatus = iota
	Receiving OrderStatus = iota
	StockFilled OrderStatus = iota
	Rejected OrderStatus = iota 
)

type Order struct {
	Id        repo.EntityId
	Number    string
	CreatedAt time.Time
	Status OrderStatus
}