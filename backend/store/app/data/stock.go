package data

type Stock map[EntityId]StockItem

type StockItem struct {
	BookId       EntityId
	OnhandQty    int
	PreservedQty int
}

func (stock Stock) Clone() Stock {
	clone := Stock{}
	for key, value := range stock {
		clone[key] = value
	}

	return clone
}

func (stock Stock) Enough(items []OrderItem) bool {
	enoughStock := true
	for _, item := range items {
		if item.Qty > stock[item.BookId].OnhandQty {
			enoughStock = false
		}
	}

	return enoughStock
}

func (stock Stock) Issue(items []OrderItem) Stock {
	newStock := stock.Clone()
	for _, item := range items {
		newStock[item.BookId] = StockItem{
			BookId:       item.BookId,
			OnhandQty:    newStock[item.BookId].OnhandQty - item.Qty,
			PreservedQty: newStock[item.BookId].PreservedQty,
		}
	}

	return newStock
}

type StockRepository interface {
	GetByOrderItems(items []OrderItem, tx Transaction) (Stock, error)
}
