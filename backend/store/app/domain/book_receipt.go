package domain

import (
	data "store/app/data"
)

type BookReceipt struct {
	data.BookReceipt
}

type ReceivingBook struct {
	data.Book
	ReceivingQty int
}

func (BookReceipt) NewFromReceivingBooks(books []ReceivingBook) *BookReceipt {
	receipt := data.BookReceipt{
		Id: data.NewEntityId(),
	}

	items := []data.BookReceiptItem{}

	for _, book := range books {
		item := data.BookReceiptItem{
			Id:            data.NewEntityId(),
			BookReceiptId: receipt.Id,
			BookId:        book.Id,
			Qty:           book.ReceivingQty,
		}

		items = append(items, item)
	}

	receipt.Items = items

	return &BookReceipt{BookReceipt: receipt}
}

func (BookReceipt) New(dataReceipt *data.BookReceipt) *BookReceipt {
	receipt := &BookReceipt{BookReceipt: *dataReceipt}
	return receipt
}

func (receipt *BookReceipt) IncreaseStock() Stock {
	stock := Stock{Stock: receipt.Stock}
	receipt.Stock = stock.IncreaseByReceipt(receipt).Stock
	return stock
}
