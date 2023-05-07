package domain

import (
	"github.com/thoas/go-funk"
)

type Stock map[string]StockItem

type StockItem struct {
	BookId      string
	OnhandQty   int
	ReservedQty int
}

func (stock Stock) Clone() Stock {
	return funk.Map(stock, func(key string, value StockItem) (string, StockItem) {
		return key, StockItem{
			BookId:      value.BookId,
			OnhandQty:   value.OnhandQty,
			ReservedQty: value.ReservedQty,
		}
	}).(Stock)
}
