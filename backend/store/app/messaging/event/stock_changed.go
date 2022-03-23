package event

import "store/app/messaging"

type StockChanged struct {
	*messaging.Message
	Items []*StockChangedItem
}

type StockChangedItem struct {
	BookId      string
	OnhandQty   int
	ReservedQty int
}
