package action

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"math/rand"
	"org.gene/singularity/pkg/context"
	"org.gene/singularity/pkg/db"
	"org.gene/singularity/pkg/job"
	"time"
)

func ExecuteStart() {
	context.SetRunning()
	fmt.Println("start!")
	go runJobs()
}

func myTask() {
	//fmt.Println("This task will run periodically")
	value := rand.Float64()
	country := db.PickCountry()
	country.WarmingMetrics[0].Value = value
}

func executeCronJob() {
	gocron.Every(1).Second().Do(myTask)
	<-gocron.Start()
}

func SomeAPICallHandler() {
	time.Sleep(10000 * time.Millisecond)
}

func runJobs() {
	for _, country := range *db.GetCountries() {
		for index := range country.WarmingMetrics {
			if country.WarmingMetrics[index].Key == "air_quality" {
				value := job.GetAirQualityFor(country.Latitude, country.Longitude)
				country.WarmingMetrics[index].Value = value
			}
		}
	}
}
