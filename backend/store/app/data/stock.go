package data

type Stock map[EntityId]StockItem

type StockItem struct {
	BookId      EntityId
	OnhandQty   int
	ReservedQty int
}
