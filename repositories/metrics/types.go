package metrics

import (
	"github.com/kaellybot/kaelly-metrics/utils/databases"
)

type Repository interface {
}

type Impl struct {
	db databases.InfluxDBConnection
}
