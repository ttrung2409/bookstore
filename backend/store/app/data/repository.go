package data

type repositoryBase interface {
	Get(id Identifier) (interface{}, error)
	Create(entity interface{}, transaction *Transaction) (Identifier, error)
	CreateIfNotExist(entity interface{}, transaction *Transaction) (Identifier, error)
	Update(id Identifier, entity interface{}, transaction *Transaction) error
}
