package data

type Stock map[EntityId]StockItem

type StockItem struct {
	BookId       EntityId
	OnhandQty    int
	PreservedQty int
}

type StockRepository interface {
	GetByOrderItems(items []OrderItem, tx Transaction) (Stock, error)
}
