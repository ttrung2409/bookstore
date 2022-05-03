package domain

import "store/app/domain/data"

type Stock struct {
	state data.Stock
}

func (Stock) New(stock data.Stock) Stock {
	return Stock{state: stock}
}

func (stock Stock) enoughForOrder(order Order) bool {
	for _, item := range order.state.Items {
		if item.Qty > stock.state[item.BookId].OnhandQty-stock.state[item.BookId].ReservedQty {
			return false
		}
	}

	return true
}

func (stock Stock) decreaseByOrder(order Order) Stock {
	for _, item := range order.state.Items {
		stock.state[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   stock.state[item.BookId].OnhandQty - item.Qty,
			ReservedQty: stock.state[item.BookId].ReservedQty,
		}
	}

	return stock
}

func (stock Stock) reserveForOrder(order Order) Stock {
	for _, item := range order.state.Items {
		stock.state[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   stock.state[item.BookId].OnhandQty,
			ReservedQty: stock.state[item.BookId].ReservedQty + item.Qty,
		}
	}

	return stock
}

func (stock Stock) releaseReservation(order Order) Stock {
	for _, item := range order.state.Items {
		stock.state[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   stock.state[item.BookId].OnhandQty,
			ReservedQty: stock.state[item.BookId].ReservedQty - item.Qty,
		}
	}

	return stock
}
