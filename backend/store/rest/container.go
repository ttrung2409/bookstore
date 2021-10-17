package rest

import (
	OrderCommand "store/app/order/command"
	OrderQuery "store/app/order/query"
	ReceivingCommand "store/app/receiving/command"
	ReceivingQuery "store/app/receiving/query"
	"store/container"
	"store/utils"
)

var receivingCommand = container.Instance().Get(utils.Nameof((*ReceivingCommand.Command)(nil))).(ReceivingCommand.Command)

var receivingQuery = container.Instance().Get(utils.Nameof((*ReceivingQuery.Query)(nil))).(ReceivingQuery.Query)

var orderQuery = container.Instance().Get(utils.Nameof((*OrderQuery.Query)(nil))).(OrderQuery.Query)

var orderCommand = container.Instance().Get(utils.Nameof((*OrderCommand.Command)(nil))).(OrderCommand.Command)
