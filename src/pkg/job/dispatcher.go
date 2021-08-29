package job

import (
	"github.com/jasonlvhit/gocron"
	"github.com/margostino/singularity/pkg/config"
)

var WorkersPool []Worker

type Worker struct {
	Interval uint64
	Action   func()
}

var ActionStorage = map[string]func(){
	"air_quality": UpdateAirQuality,
	"world_cycle": UpdateWorldCycle,
}

func LoadJobs() {
	jobs := config.GetJobsConfiguration()
	for _, job := range *jobs {
		worker := Worker{
			Interval: getInterval(job.Schedule),
			Action:   ActionStorage[job.Id],
		}
		WorkersPool = append(WorkersPool, worker)
	}
}

func Dispatch() {
	for _, worker := range WorkersPool {
		go executeCronJob(worker.Interval, worker.Action)
	}
}

func executeCronJob(interval uint64, action func()) {
	gocron.Every(interval).Second().Do(action)
	<-gocron.Start()
}

func getInterval(schedule string) uint64 {
	if schedule == "daily" {
		return 24
	}
	return 1
}
