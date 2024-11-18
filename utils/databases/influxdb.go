package databases

import (
	"context"

	"github.com/kaellybot/kaelly-metrics/models/constants"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDBConnection interface {
	WriteAPI() api.WriteAPI
	IsConnected() bool
	Shutdown()
}

type influxDBConnection struct {
	client   influxdb2.Client
	writeAPI api.WriteAPI
}

func New() InfluxDBConnection {
	url := viper.GetString(constants.InfluxDBURL)
	token := viper.GetString(constants.InfluxDBToken)
	org := viper.GetString(constants.InfluxDBOrg)
	bucket := viper.GetString(constants.InfluxDBBucket)

	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPI(org, bucket)

	return &influxDBConnection{
		client:   client,
		writeAPI: writeAPI,
	}
}

func (c *influxDBConnection) WriteAPI() api.WriteAPI {
	return c.writeAPI
}

func (c *influxDBConnection) IsConnected() bool {
	if c.client == nil {
		return false
	}

	_, err := c.client.Health(context.Background())
	return err == nil
}

func (c *influxDBConnection) Shutdown() {
	log.Info().Msgf("Shutdown connection to InfluxDB")

	// Since write is async, data is not flushed everytime.
	// To not lose data, flush while shutting down.
	c.writeAPI.Flush()
	c.client.Close()
}
