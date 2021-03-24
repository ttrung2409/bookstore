package domain

import "time"

const (
	Queued = iota
	Accepted = iota
	Receiving = iota
	StockFilled = iota
	Rejected = iota
)

type Order struct {
	id        string
	number    string
	createdAt time.Time
	status string
}