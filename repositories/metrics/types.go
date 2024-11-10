package metrics

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/utils/databases"
)

type Repository interface {
	Write(message *amqp.RabbitMQMessage, correlationID, replyTo string,
		timestamp time.Time)
}

type Impl struct {
	db databases.InfluxDBConnection
}
