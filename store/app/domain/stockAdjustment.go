package domain

import "github.com/thoas/go-funk"

type StockAdjustment []StockAdjustmentItem

type StockType string

const (
	StockTypeOnhand   StockType = "Onhand"
	StockTypeReserved StockType = "Reserved"
)

type StockAdjustmentItem struct {
	BookId    string
	Qty       int
	StockType StockType
}

func (adjustment StockAdjustment) Clone() StockAdjustment {
	return funk.Map(adjustment, func(item StockAdjustmentItem) StockAdjustmentItem {
		return StockAdjustmentItem{BookId: item.BookId, Qty: item.Qty}
	}).(StockAdjustment)
}
