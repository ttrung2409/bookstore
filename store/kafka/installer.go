package kafka

import (
	"store/app/messaging"
	"store/utils"

	"github.com/sarulabs/di"
)

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name:  utils.Nameof((*messaging.EventDispatcher)(nil)),
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return &eventDispatcher{}, nil
			},
		},
	}...)
}
