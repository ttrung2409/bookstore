package command

import (
	"store/app/order/command"
	"store/app/order/query"

	"github.com/sarulabs/di"
)

func Install(builder *di.Builder) {
	query.Install(builder)
	command.Install(builder)
}
