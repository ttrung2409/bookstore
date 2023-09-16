package repository

import (
	"context"
	"ecommerce/app/domain"
	"ecommerce/kafka"
	"encoding/json"

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

type EventDispatcher interface {
	Dispatch(topic string, key string, events ...domain.Event) error
	Dispose()
}

var dispatcher EventDispatcher

func GetEventDispatcher() EventDispatcher {
	if dispatcher == nil {
		dispatcher = &eventDispatcher{producers: map[string]kafka.Producer{}}
	}

	return dispatcher
}

type eventDispatcher struct {
	producers map[string]kafka.Producer
}

func (d *eventDispatcher) Dispatch(topic string, key string, events ...domain.Event) error {
	if len(events) == 0 {
		return nil
	}

	producer, ok := d.producers[topic]
	if !ok {
		producer = kafka.NewProducer(topic)
		d.producers[topic] = producer
	}

	kafkaEvents := funk.Map(events, func(event domain.Event) kafkaEvent {
		return kafkaEvent{event, key}
	}).([]kafka.Message)

	if err := producer.Send(context.Background(), kafkaEvents...); err != nil {
		return err
	}

	return nil
}

func (d *eventDispatcher) Dispose() {
	if d.producers == nil {
		return
	}

	for _, producer := range d.producers {
		producer.Dispose()
	}
}
