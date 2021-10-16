package query

import (
	"store/utils"

	"github.com/sarulabs/di"
)

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*BookQuery)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return bookQuery{}, nil
			},
		},
	}...)
}
