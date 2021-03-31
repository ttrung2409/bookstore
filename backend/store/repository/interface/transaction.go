package repository

type Transaction interface {
	Commit()
	Rollback()
}

type TransactionFactory interface {
	New() Transaction
}