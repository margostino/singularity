package preload

import "org.gene/singularity/pkg/config"

func Preload() {
	if config.IsPreLoadEnabled() {
		preloadCountries()
		preloadPopulation()
		//preloadCompanies()
		//preloadClimate()
	}
}
