package preload

import (
	"math/rand"
	"github.com/margostino/singularity/pkg/config"
	"github.com/margostino/singularity/pkg/db"
)

func InitializeWarmingMetrics() []db.WarmingMetric {
	var warmingMetrics []db.WarmingMetric
	metrics := *config.GetMetricsConfiguration()
	value := rand.Float64()
	for _, metric := range metrics {
		warmingMetric := db.WarmingMetric{
			Key:         metric.Id,
			Value:       value,
			Unit:        metric.Unit,
			Description: metric.Description,
		}
		warmingMetrics = append(warmingMetrics, warmingMetric)
	}
	return warmingMetrics
}
