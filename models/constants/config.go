package constants

import "github.com/rs/zerolog"

const (
	ConfigFileName = ".env"

	// InfluxDB URL.
	InfluxDBURL = "INFLUXDB_URL"

	// InfluxDB Token, generated from WebUI > Load Data > API Tokens.
	InfluxDBToken = "INFLUXDB_TOKEN"

	// InfluxDB Org.
	InfluxDBOrg = "INFLUXDB_ORG"

	// InfluxDB Bucket.
	InfluxDBBucket = "INFLUXDB_BUCKET"

	// RabbitMQ address.
	RabbitMQAddress = "RABBITMQ_ADDRESS"

	// Metric port.
	MetricPort = "METRIC_PORT"

	// Zerolog values from [trace, debug, info, warn, error, fatal, panic].
	LogLevel = "LOG_LEVEL"

	// Boolean; used to register commands at development guild level or globally.
	Production = "PRODUCTION"

	defaultInfluxDBURL     = "http://localhost:8086"
	defaultInfluxDBToken   = ""
	defaultInfluxDBOrg     = "kaellybot"
	defaultInfluxDBBucket  = "kaellybot"
	defaultRabbitMQAddress = "amqp://localhost:5672"
	defaultMetricPort      = 2112
	defaultLogLevel        = zerolog.InfoLevel
	defaultProduction      = false
)

func GetDefaultConfigValues() map[string]any {
	return map[string]any{
		InfluxDBURL:     defaultInfluxDBURL,
		InfluxDBToken:   defaultInfluxDBToken,
		InfluxDBOrg:     defaultInfluxDBOrg,
		InfluxDBBucket:  defaultInfluxDBBucket,
		RabbitMQAddress: defaultRabbitMQAddress,
		MetricPort:      defaultMetricPort,
		LogLevel:        defaultLogLevel.String(),
		Production:      defaultProduction,
	}
}
