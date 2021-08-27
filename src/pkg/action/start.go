package action

import (
	"math/rand"
	"org.gene/singularity/pkg/context"
	"org.gene/singularity/pkg/db"
	"org.gene/singularity/pkg/job"
	"time"
)

func ExecuteStart() {
	context.SetRunning()
	job.Dispatch()
}

func myTask() {
	value := rand.Float64()
	country := db.PickCountry()
	country.WarmingMetrics[0].Value = value
}

func SomeAPICallHandler() {
	time.Sleep(10000 * time.Millisecond)
}
