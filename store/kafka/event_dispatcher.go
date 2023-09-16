package kafka

import (
	"context"
	"encoding/json"
	"store/app/domain"

	"github.com/thoas/go-funk"
)

type kafkaEvent struct {
	domain.Event
	key string
}

func (e kafkaEvent) Key() string {
	return e.key
}

func (e kafkaEvent) Value() []byte {
	serialized, _ := json.Marshal(e)

	return serialized
}

type eventDispatcher struct {
	producers map[string]Producer
}

func (d *eventDispatcher) Dispatch(topic string, key string, events ...domain.Event) error {
	if len(events) == 0 {
		return nil
	}

	producer, ok := d.producers[topic]
	if !ok {
		producer = NewProducer(topic)
		d.producers[topic] = producer
	}

	kafkaEvents := funk.Map(events, func(event domain.Event) kafkaEvent {
		return kafkaEvent{event, key}
	}).([]Message)

	if err := producer.Send(context.Background(), kafkaEvents...); err != nil {
		return err
	}

	return nil
}

func (d *eventDispatcher) Dispose() error {
	if d.producers == nil {
		return nil
	}

	for _, producer := range d.producers {
		if error := producer.Dispose(); error != nil {
			return error
		}
	}

	return nil
}
