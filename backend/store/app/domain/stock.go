package domain

import "store/app/data"

type Stock struct {
	data.Stock
}

func (stock Stock) Clone() Stock {
	clone := data.Stock{}
	for key, value := range stock.Stock {
		clone[key] = data.StockItem{
			BookId:      value.BookId,
			OnhandQty:   value.OnhandQty,
			ReservedQty: value.ReservedQty,
		}
	}

	return Stock{Stock: clone}
}

func (stock Stock) EnoughForOrder(order *Order) bool {
	enoughStock := true
	for _, item := range order.Items {
		if item.Qty > stock.Stock[item.BookId].OnhandQty {
			enoughStock = false
		}
	}

	return enoughStock
}

func (stock Stock) IncreaseByReceipt(receipt *BookReceipt) Stock {
	newStock := stock.Clone()
	for _, item := range receipt.Items {
		newStock.Stock[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   newStock.Stock[item.BookId].OnhandQty + item.Qty,
			ReservedQty: newStock.Stock[item.BookId].ReservedQty,
		}
	}

	return newStock
}

func (stock Stock) DecreaseByOrder(order *Order) Stock {
	newStock := stock.Clone()
	for _, item := range order.Items {
		newStock.Stock[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   newStock.Stock[item.BookId].OnhandQty - item.Qty,
			ReservedQty: newStock.Stock[item.BookId].ReservedQty,
		}
	}

	return newStock
}

func (stock Stock) ReserveForOrder(order *Order) Stock {
	newStock := stock.Clone()
	for _, item := range order.Items {
		newStock.Stock[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   newStock.Stock[item.BookId].OnhandQty,
			ReservedQty: newStock.Stock[item.BookId].ReservedQty + item.Qty,
		}
	}

	return newStock
}

func (stock Stock) ReleaseReservation(order *Order) Stock {
	newStock := stock.Clone()
	for _, item := range order.Items {
		newStock.Stock[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   newStock.Stock[item.BookId].OnhandQty,
			ReservedQty: newStock.Stock[item.BookId].ReservedQty - item.Qty,
		}
	}

	return newStock
}
