package data

type Transaction interface {
	Commit() error
	Rollback() error
}

type TransactionalFunc func(tx Transaction) (interface{}, error)

type TransactionFactory interface {
	New() Transaction
	RunInTransaction(fn TransactionalFunc, ambientTx Transaction) (interface{}, error)
}
