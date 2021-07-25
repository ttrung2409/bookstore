package command

import (
	module "store"
	repo "store/app/repository"
	"store/utils"
)

var BookRepository = module.Container().Get(utils.Nameof((*repo.BookRepository)(nil))).(repo.BookRepository)

var OrderRepository = module.Container().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)

var BookReceiptRepository = module.Container().Get(utils.Nameof((*repo.BookReceiptRepository)(nil))).(repo.BookReceiptRepository)

var TransactionFactory = module.Container().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
