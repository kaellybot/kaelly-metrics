package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/services/metrics"
	"github.com/kaellybot/kaelly-metrics/utils/databases"
	"github.com/kaellybot/kaelly-metrics/utils/insights"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	metricService metrics.Service
	broker        amqp.MessageBroker
	db            databases.InfluxDBConnection
	probes        insights.Probes
	prom          insights.PrometheusMetrics
}
