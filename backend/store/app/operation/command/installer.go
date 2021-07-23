package command

import (
	"store/utils"

	"github.com/sarulabs/di"
)

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*ReceiveBookCommand)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return receiveBookCommand{}, nil
			},
		},
		{
			Name: utils.Nameof((*AcceptOrderCommand)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return acceptOrderCommand{}, nil
			},
		},
		{
			Name: utils.Nameof((*placeAsBackOrderCommand)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return placeAsBackOrderCommand{}, nil
			},
		},
		{
			Name: utils.Nameof((*RejectOrderCommand)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return rejectOrderCommand{}, nil
			},
		},
	}...)
}
