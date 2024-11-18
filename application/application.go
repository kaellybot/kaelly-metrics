package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/models/constants"
	metricRepo "github.com/kaellybot/kaelly-metrics/repositories/metrics"
	"github.com/kaellybot/kaelly-metrics/services/metrics"
	"github.com/kaellybot/kaelly-metrics/utils/databases"
	"github.com/kaellybot/kaelly-metrics/utils/insights"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	broker := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress),
		amqp.WithBindings(metrics.GetBinding()))
	db := databases.New()
	probes := insights.NewProbes(broker.IsConnected, db.IsConnected)
	prom := insights.NewPrometheusMetrics()

	// repositories
	metricRepo := metricRepo.New(db)

	// services
	metricService := metrics.New(broker, metricRepo)

	return &Impl{
		metricService: metricService,
		broker:        broker,
		db:            db,
		probes:        probes,
		prom:          prom,
	}, nil
}

func (app *Impl) Run() error {
	app.probes.ListenAndServe()
	app.prom.ListenAndServe()

	if err := app.broker.Run(); err != nil {
		return err
	}

	app.metricService.Consume()
	return nil
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	app.db.Shutdown()
	app.prom.Shutdown()
	app.probes.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
