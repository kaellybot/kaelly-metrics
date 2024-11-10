package metrics

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/utils/databases"
)

func New(db databases.InfluxDBConnection) *Impl {
	return &Impl{db: db}
}

func (impl *Impl) Write(message *amqp.RabbitMQMessage, correlationID, replyTo string,
	timestamp time.Time) {
	p := influxdb2.NewPoint(
		"request",
		map[string]string{
			"game":          message.Game.String(),
			"language":      message.Language.String(),
			"type":          message.Type.String(),
			"userID":        message.UserID,
			"shard":         replyTo,
			"correlationID": correlationID,
		},
		map[string]interface{}{
			"requestNumber": 1,
		},
		timestamp)

	// write asynchronously
	impl.db.WriteAPI().WritePoint(p)
}
