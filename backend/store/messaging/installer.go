package messaging

import (
	"store/app/messaging"
	"store/utils"

	"github.com/sarulabs/di"
)

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*messaging.Producer)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return nil, nil
			},
		},
	}...)
}
