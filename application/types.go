package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/services/metrics"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	metricService metrics.Service
	broker        amqp.MessageBroker
}
