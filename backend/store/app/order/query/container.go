package query

import (
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

var OrderRepository = container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)
