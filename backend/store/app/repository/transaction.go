package repository

type Transaction interface {
	Commit() error
	Rollback() error
}

type TransactionalFunc func(tx Transaction) (interface{}, error)

type TransactionFactory interface {
	New() Transaction
	RunInTransaction(fn TransactionalFunc) (interface{}, error)
}
