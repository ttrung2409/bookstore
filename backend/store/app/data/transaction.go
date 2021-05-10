package data

type Transaction interface {
	Commit() error
	Rollback() error
}

type TransactionFactory interface {
	New() Transaction
}
