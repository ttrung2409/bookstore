package domain

import "store/app/data"

type Stock struct {
	state data.Stock
}

func (Stock) New(stock data.Stock) Stock {
	return Stock{state: stock}
}

func (stock Stock) Clone() Stock {
	return Stock{state: stock.state.Clone()}
}

func (stock Stock) enoughForOrder(order *Order) bool {
	enoughStock := true
	for _, item := range order.state.Items {
		if item.Qty > stock.state[item.BookId].OnhandQty {
			enoughStock = false
		}
	}

	return enoughStock
}

func (stock Stock) decreaseByOrder(order *Order) Stock {
	newStock := stock.Clone()
	for _, item := range order.state.Items {
		newStock.state[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   newStock.state[item.BookId].OnhandQty - item.Qty,
			ReservedQty: newStock.state[item.BookId].ReservedQty,
		}
	}

	return newStock
}

func (stock Stock) reserveForOrder(order *Order) Stock {
	newStock := stock.Clone()
	for _, item := range order.state.Items {
		newStock.state[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   newStock.state[item.BookId].OnhandQty,
			ReservedQty: newStock.state[item.BookId].ReservedQty + item.Qty,
		}
	}

	return newStock
}

func (stock Stock) releaseReservation(order *Order) Stock {
	newStock := stock.Clone()
	for _, item := range order.state.Items {
		newStock.state[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   newStock.state[item.BookId].OnhandQty,
			ReservedQty: newStock.state[item.BookId].ReservedQty - item.Qty,
		}
	}

	return newStock
}
