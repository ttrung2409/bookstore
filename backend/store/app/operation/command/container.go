package operation

import (
	module "store"
	"store/app/data"
	"store/utils"
)

var BookRepository = module.Container().Get(utils.Nameof((*data.BookRepository)(nil))).(data.BookRepository)

var OrderRepository = module.Container().Get(utils.Nameof((*data.OrderRepository)(nil))).(data.OrderRepository)

var BookReceiptRepository = module.Container().Get(utils.Nameof((*data.BookReceiptRepository)(nil))).(data.BookReceiptRepository)

var TransactionFactory = module.Container().Get(utils.Nameof((*data.TransactionFactory)(nil))).(data.TransactionFactory)
