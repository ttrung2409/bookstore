package domain

import (
	"github.com/thoas/go-funk"
)

type StockData map[string]StockItemData

type StockItemData struct {
	BookId      string
	OnhandQty   int
	ReservedQty int
}

func (stock StockData) Clone() StockData {
	return funk.Map(stock, func(key string, value StockItemData) (string, StockItemData) {
		return key, StockItemData{
			BookId:      value.BookId,
			OnhandQty:   value.OnhandQty,
			ReservedQty: value.ReservedQty,
		}
	}).(StockData)
}

type Stock struct {
	state StockData
}

func (Stock) New(stock StockData) *Stock {
	return &Stock{state: stock.Clone()}
}

func (stock *Stock) State() StockData {
	return stock.state.Clone()
}

func (stock *Stock) clone() *Stock {
	return &Stock{state: stock.state.Clone()}
}

func (stock *Stock) enoughForOrder(order Order) bool {
	for _, item := range order.state.Items {
		if stockItem, ok := stock.state[item.BookId]; ok {
			if item.Qty > stockItem.OnhandQty-stock.state[item.BookId].ReservedQty {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func (stock *Stock) decreaseByOrder(order Order) *Stock {
	newStock := stock.clone()

	for _, item := range order.state.Items {
		if _, ok := newStock.state[item.BookId]; ok {
			newStock.state[item.BookId] = StockItemData{
				BookId:      item.BookId,
				OnhandQty:   newStock.state[item.BookId].OnhandQty - item.Qty,
				ReservedQty: newStock.state[item.BookId].ReservedQty,
			}
		}
	}

	return newStock
}

func (stock *Stock) reserveForOrder(order Order) *Stock {
	newStock := stock.clone()

	for _, item := range order.state.Items {
		if _, ok := newStock.state[item.BookId]; ok {
			newStock.state[item.BookId] = StockItemData{
				BookId:      item.BookId,
				OnhandQty:   newStock.state[item.BookId].OnhandQty,
				ReservedQty: newStock.state[item.BookId].ReservedQty + item.Qty,
			}
		}
	}

	return newStock
}

func (stock *Stock) releaseReservation(order Order) *Stock {
	newStock := stock.clone()

	for _, item := range order.state.Items {
		if _, ok := stock.state[item.BookId]; ok {
			newStock.state[item.BookId] = StockItemData{
				BookId:      item.BookId,
				OnhandQty:   newStock.state[item.BookId].OnhandQty,
				ReservedQty: newStock.state[item.BookId].ReservedQty - item.Qty,
			}
		}
	}

	return newStock
}
