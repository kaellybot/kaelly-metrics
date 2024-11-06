package databases

import (
	"github.com/kaellybot/kaelly-metrics/models/constants"
	"github.com/spf13/viper"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDBConnection interface {
	WriteAPI() api.WriteAPI
	Shutdown()
}

type InfluxDBConnectionImpl struct {
	client   influxdb2.Client
	writeAPI api.WriteAPI
}

func New() *InfluxDBConnectionImpl {
	url := viper.GetString(constants.InfluxDBURL)
	token := viper.GetString(constants.InfluxDBToken)
	org := viper.GetString(constants.InfluxDBOrg)
	bucket := viper.GetString(constants.InfluxDBBucket)

	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPI(org, bucket)

	return &InfluxDBConnectionImpl{
		client:   client,
		writeAPI: writeAPI,
	}
}

func (connection *InfluxDBConnectionImpl) WriteAPI() api.WriteAPI {
	return connection.writeAPI
}

func (connection *InfluxDBConnectionImpl) Shutdown() {
	// Since write is async, data is not flushed everytime.
	// To not lose data, flush while shutting down.
	connection.writeAPI.Flush()
	connection.client.Close()
}
