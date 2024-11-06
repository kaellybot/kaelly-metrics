package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/models/constants"
	metricRepo "github.com/kaellybot/kaelly-metrics/repositories/metrics"
	"github.com/kaellybot/kaelly-metrics/services/metrics"
	"github.com/kaellybot/kaelly-metrics/utils/databases"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	db := databases.New()

	broker, err := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress),
		[]amqp.Binding{metrics.GetBinding()})
	if err != nil {
		return nil, err
	}

	// repositories
	metricRepo := metricRepo.New(db)

	// services
	metricService := metrics.New(broker, metricRepo)

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
