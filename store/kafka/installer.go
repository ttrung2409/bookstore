package kafka

import (
	"store/app/messaging"
	"store/utils"

	"github.com/sarulabs/di"
)

func RegisterDependencies(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name:  utils.Nameof((*messaging.EventDispatcher)(nil)),
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return &eventDispatcher{producers: map[string]Producer{}}, nil
			},
			Close: func(obj interface{}) error {
				return obj.(*eventDispatcher).Dispose()
			},
		}}...)
}
