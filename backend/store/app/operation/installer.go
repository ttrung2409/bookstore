package operation

import (
	"store/utils"

	"github.com/sarulabs/di"
)

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*GoogleBookQuery)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return googleBookQuery{}, nil
			},
		},
	}...)
}
