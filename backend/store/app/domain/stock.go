package domain

import "store/app/data"

type Stock data.Stock

func (stock Stock) Clone() Stock {
	clone := Stock{}
	for key, value := range stock {
		clone[key] = value
	}

	return clone
}

func (stock Stock) Enough(items []data.OrderItem) bool {
	enoughStock := true
	for _, item := range items {
		if item.Qty > stock[item.BookId].OnhandQty {
			enoughStock = false
		}
	}

	return enoughStock
}

func (stock Stock) Issue(items []data.OrderItem) Stock {
	newStock := stock.Clone()
	for _, item := range items {
		newStock[item.BookId] = data.StockItem{
			BookId:       item.BookId,
			OnhandQty:    newStock[item.BookId].OnhandQty - item.Qty,
			PreservedQty: newStock[item.BookId].PreservedQty,
		}
	}

	return newStock
}
