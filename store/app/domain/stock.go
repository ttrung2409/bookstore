package domain

import (
	"store/app/domain/data"
)

type Stock struct {
	state data.Stock
}

func (Stock) New(stock data.Stock) *Stock {
	return &Stock{state: stock.Clone()}
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
			newStock.state[item.BookId] = data.StockItem{
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
			newStock.state[item.BookId] = data.StockItem{
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
			newStock.state[item.BookId] = data.StockItem{
				BookId:      item.BookId,
				OnhandQty:   newStock.state[item.BookId].OnhandQty,
				ReservedQty: newStock.state[item.BookId].ReservedQty - item.Qty,
			}
		}
	}

	return newStock
}
