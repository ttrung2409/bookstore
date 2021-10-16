package query

import (
	module "store"
	repo "store/app/repository"
	"store/utils"
)

var OrderRepository = module.Container().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)
