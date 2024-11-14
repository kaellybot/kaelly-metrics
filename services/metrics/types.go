package metrics

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/repositories/metrics"
)

const (
	queueName  = "metrics-requests"
	routingkey = "requests.*"
)

type Service interface {
	Consume()
}

type Impl struct {
	broker     amqp.MessageBroker
	repository metrics.Repository
}
