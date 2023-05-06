package domain

import "github.com/thoas/go-funk"

type StockAdjustmentData []StockAdjustmentItemData

func (adjustment StockAdjustmentData) Clone() StockAdjustmentData {
	return funk.Map(adjustment, func(item StockAdjustmentItemData) StockAdjustmentItemData {
		return StockAdjustmentItemData{BookId: item.BookId, Qty: item.Qty}
	}).(StockAdjustmentData)
}

type StockAdjustmentItemData struct {
	BookId string
	Qty    int
}
