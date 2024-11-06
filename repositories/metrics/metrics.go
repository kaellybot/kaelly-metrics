package metrics

import "github.com/kaellybot/kaelly-metrics/utils/databases"

func New(db databases.InfluxDBConnection) *Impl {
	return &Impl{db: db}
}
