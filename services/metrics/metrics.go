package metrics

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

func New(broker amqp.MessageBroker) (*Impl, error) {
	return &Impl{
		broker: broker,
	}, nil
}

func (service *Impl) Consume() error {
	// TODO
	return nil
}
