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

	broker := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress),
		amqp.WithBindings(metrics.GetBinding()))

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
	if err := app.broker.Run(); err != nil {
		return err
	}

	app.metricService.Consume()
	return nil
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
