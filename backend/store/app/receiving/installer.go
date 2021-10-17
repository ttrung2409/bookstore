package command

import (
	"store/app/receiving/command"
	"store/app/receiving/query"

	"github.com/sarulabs/di"
)

func Install(builder *di.Builder) {
	query.Install(builder)
	command.Install(builder)
}
