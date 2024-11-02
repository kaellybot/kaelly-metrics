package metrics

import (
	amqp "github.com/kaellybot/kaelly-amqp"
)

const (
	routingkey = "requests.*"
)

type Service interface {
	Consume() error
}

type Impl struct {
	broker amqp.MessageBroker
}
