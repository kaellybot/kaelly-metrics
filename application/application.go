package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/models/constants"
	"github.com/kaellybot/kaelly-metrics/services/metrics"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	// TODO influx

	broker, err := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress), nil)
	if err != nil {
		return nil, err
	}

	// services
	metricService, err := metrics.New(broker)
	if err != nil {
		return nil, err
	}

	return &Impl{
		metricService: metricService,
		broker:        broker,
	}, nil
}

func (app *Impl) Run() error {
	return app.metricService.Consume()
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
