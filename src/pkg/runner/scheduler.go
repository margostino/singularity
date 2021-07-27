package options

import (
	"sort"
	"time"
)

type Scheduler struct {
	jobs []*Job // Array store jobs
}

var (
	defaultScheduler = NewScheduler()
)

func NewScheduler() *Scheduler {
	return &Scheduler{
		***REMOVED*** []*Job{},
	}
}

func (s *Scheduler) Jobs() []*Job {
	return s.jobs
}

func (s *Scheduler) Every(interval uint64) *Job {
	job := NewJob(interval)
	s.jobs[len(s.jobs)] = job
	return job
}

func Every(interval uint64) *Job {
	return defaultScheduler.Every(interval)
}

func (s *Scheduler) getRunnableJobs() (runningJobs []*Job, n int) {
	runnableJobs := []*Job{}
	n = 0
	//sort.Sort(s)
	for _, job := range s.jobs {
		if job.shouldRun() {
			runnableJobs[n] = job
			n++
		} else {
			break
		}
	}
	return runnableJobs, n
}

func (s *Scheduler) Start() chan bool {
	stopped := make(chan bool, 1)
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				s.RunPending()
			case <-stopped:
				ticker.Stop()
				return
			}
		}
	}()

	return stopped
}

func (s *Scheduler) RunPending() {
	runnableJobs, n := s.getRunnableJobs()

	if n != 0 {
		for i := 0; i < n; i++ {
			go runnableJobs[i].run()
			runnableJobs[i].lastRun = time.Now()
			runnableJobs[i].scheduleNextRun()
		}
	}
}
