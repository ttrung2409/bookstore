package command

import (
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

var BookRepository = container.Instance().Get(utils.Nameof((*repo.BookRepository)(nil))).(repo.BookRepository)

var OrderRepository = container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)

var BookReceiptRepository = container.Instance().Get(utils.Nameof((*repo.BookReceiptRepository)(nil))).(repo.BookReceiptRepository)

var TransactionFactory = container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
