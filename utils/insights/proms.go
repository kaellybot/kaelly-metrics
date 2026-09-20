package insights

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaellybot/kaelly-metrics/models/constants"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type PrometheusMetrics interface {
	ListenAndServe()
	Shutdown()
	CountRejectedMessage(messageType string)
}

type prom struct {
	server   *http.Server
	rejected *prometheus.CounterVec
}

func NewPrometheusMetrics() PrometheusMetrics {
	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", promhttp.Handler())
	return &prom{
		server: &http.Server{
			Addr:              fmt.Sprintf(":%v", viper.GetInt(constants.MetricPort)),
			Handler:           metricsMux,
			ReadHeaderTimeout: 0,
		},
		rejected: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "kaelly_rejected_messages_total",
			Help: "Messages refused by the AMQP guards before reaching the consumer, by type.",
		}, []string{"type"}),
	}
}

// CountRejectedMessage records a message the AMQP guards refused, so a producer
// publishing without a game shows up on the dashboards and not only in the logs.
func (prom *prom) CountRejectedMessage(messageType string) {
	prom.rejected.WithLabelValues(messageType).Inc()
}

func (prom *prom) ListenAndServe() {
	go func() {
		log.Info().Msgf("Exposing Prometheus metrics...")
		err := prom.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msgf("Cannot listen and serve Prometheus metrics")
		}
	}()
}

func (prom *prom) Shutdown() {
	if prom.server != nil {
		if err := prom.server.Shutdown(context.Background()); err != nil {
			log.Error().Err(err).Msgf("Failed to shutdown Prometheus metrics server")
		}
	}
}
