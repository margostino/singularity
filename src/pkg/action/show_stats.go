package action

import (
	"fmt"
	"org.gene/singularity/pkg/db"
)

func ExecuteShowStats() {
	total := len(db.Countries)
	sum := 0
	var unit string
	for _, country := range db.Countries {
		for _, metric := range country.WarmingMetrics {
			if metric.Key == "co2_emissions" {
				sum += metric.Value
				unit = metric.Unit
			}
		}
	}
	fmt.Printf("Global CO2 Emissions: %v %s\n", sum/total, unit)
}
