package data

type Transaction interface {
	Commit() error
	Rollback()
}

type TransactionFactory interface {
	New() *Transaction
}