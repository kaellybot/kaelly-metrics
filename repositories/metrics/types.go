package metrics

import (
	"context"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/utils/databases"
)

type Repository interface {
	Write(ctx context.Context, message *amqp.RabbitMQMessage, correlationID string)
}

type Impl struct {
	db databases.InfluxDBConnection
}
