package query

import (
	module "store"
	"store/app/data"
	"store/utils"
)

var OrderRepository = module.Container().Get(utils.Nameof((*data.OrderRepository)(nil))).(data.OrderRepository)
