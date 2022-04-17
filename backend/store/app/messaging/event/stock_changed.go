package event

import (
	"github.com/google/uuid"
)

type StockChanged struct {
	Items []*StockChangedItem
}

type StockChangedItem struct {
	BookId      string
	OnhandQty   int
	ReservedQty int
}

func (event StockChanged) Key() string {
	return uuid.NewString()
}

func (event StockChanged) Topic() string {
	return "stock"
}
