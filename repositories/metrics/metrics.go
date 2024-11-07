package metrics

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/utils/databases"
)

func New(db databases.InfluxDBConnection) *Impl {
	return &Impl{db: db}
}

func (impl *Impl) Write(ctx context.Context, message *amqp.RabbitMQMessage,
	correlationID string) {
	p := influxdb2.NewPoint(
		"request",
		map[string]string{
			"game":     message.Game.String(),
			"language": message.Language.String(),
			"type":     message.Type.String(),
			"userID":   message.UserID,
		},
		map[string]interface{}{
			"correlationID": correlationID,
			"requestNumber": 1,
		},
		time.Now())

	// write asynchronously
	impl.db.WriteAPI().WritePoint(p)
}
