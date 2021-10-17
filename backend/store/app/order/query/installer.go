package query

import (
	"store/utils"

	"github.com/sarulabs/di"
)

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*Query)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return query{}, nil
			},
		},
	}...)
}
