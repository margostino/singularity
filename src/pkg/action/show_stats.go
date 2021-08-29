package action

import (
	"fmt"
	"github.com/margostino/singularity/pkg/db"
)

func ExecuteShowStats() {
	var avg float64
	total := len(db.Countries)
	sum := 0.0
	var unit string
	for _, country := range db.Countries {
		for _, metric := range country.WarmingMetrics {
			if metric.Key == "co2_emissions" {
				sum += metric.Value
				unit = metric.Unit
			}
		}
	}
	avg = sum/float64(total)
	fmt.Printf("Global CO2 Emissions: %f %s\n", avg, unit)
}
