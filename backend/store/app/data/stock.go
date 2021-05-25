package data

import "github.com/thoas/go-funk"

type Stock map[EntityId]StockItem

type StockItem struct {
	BookId      EntityId
	OnhandQty   int
	ReservedQty int
}

func (stock Stock) Clone() Stock {
	return funk.Map(stock, func(key EntityId, value StockItem) (EntityId, StockItem) {
		return key, StockItem{
			BookId:      value.BookId,
			OnhandQty:   value.OnhandQty,
			ReservedQty: value.ReservedQty,
		}
	}).(Stock)
}

type StockAdjustment []StockAdjustmentItem

func (adjustment StockAdjustment) Clone() StockAdjustment {
	return funk.Map(adjustment, func(item StockAdjustmentItem) StockAdjustmentItem {
		return StockAdjustmentItem{BookId: item.BookId, Qty: item.Qty}
	}).(StockAdjustment)
}

type StockAdjustmentItem struct {
	BookId EntityId
	Qty    int
}
