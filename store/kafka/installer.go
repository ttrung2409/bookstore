package kafka

import (
	"store/app/kafka"
	"store/utils"

	"github.com/sarulabs/di"
)

const ClusterAddress = "localhost:9092"

var BrokerAddresses = []string{"localhost:9092", "localhost:9093", "localhost:9094"}

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*kafka.Factory)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return &factory{}, nil
			},
		},
	}...)
}
