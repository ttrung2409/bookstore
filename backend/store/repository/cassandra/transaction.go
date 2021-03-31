package repository

type transaction struct {
	commands []Command
}

type transactionFactory struct{}

type Command struct {
	Statement string
	Args      []string
}

func (f *transactionFactory) New() *transaction {
	return &transaction{}
}