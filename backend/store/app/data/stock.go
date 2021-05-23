package data

type Stock map[EntityId]StockItem

type StockItem struct {
	BookId      EntityId
	OnhandQty   int
	ReservedQty int
}

func (stock Stock) Clone() Stock {
	clone := Stock{}
	for key, value := range stock {
		clone[key] = StockItem{
			BookId:      value.BookId,
			OnhandQty:   value.OnhandQty,
			ReservedQty: value.ReservedQty,
		}
	}

	return clone
}
