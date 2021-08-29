package action

import (
	"math/rand"
	"github.com/margostino/singularity/pkg/context"
	"github.com/margostino/singularity/pkg/db"
	"github.com/margostino/singularity/pkg/job"
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
