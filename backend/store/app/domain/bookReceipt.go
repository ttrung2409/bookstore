package domain

import data "store/data"

type BookReceipt struct {
	data.BookReceipt
}

func CreateBookReceipt(books []data.Book, transaction *data.Transaction) (*BookReceipt, error) {
}