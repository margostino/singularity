package preload

import "github.com/margostino/singularity/pkg/config"

func Preload() {
	if config.IsPreLoadEnabled() {
		preloadCountries()
		preloadPopulation()
		//preloadCompanies()
		//preloadGlobalWarming()
	}
}
