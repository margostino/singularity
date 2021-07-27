package options

import (
	"time"
)

// Job struct keeping information about job
type Job struct {
	interval uint64                   // pause interval * unit between runs
	jobFunc  string                   // the job jobFunc to run, func[jobFunc]
	unit     int                 // time units, ,e.g. 'minutes', 'hours'...
	atTime   time.Duration            // optional time at which this job runs
	err      error                    // error related to job
	lastRun  time.Time                // datetime of last run
	nextRun  time.Time                // datetime of next run
	startDay time.Weekday             // Specific day of the week to start on
	funcs    map[string]interface{}   // Map for the function task store
	fparams  map[string][]interface{} // Map for function and  params of function
	lock     bool                     // lock the job from running at same time form multiple instances
	tags     []string                 // allow the user to tag jobs with certain labels
}

func NewJob(interval uint64) *Job {
	return &Job{
		interval: interval,
		lastRun:  time.Unix(0, 0),
		nextRun:  time.Unix(0, 0),
		startDay: time.Sunday,
		funcs:    make(map[string]interface{}),
		fparams:  make(map[string][]interface{}),
		tags:     []string{},
	}
}

func (j *Job) shouldRun() bool {
	return time.Now().Unix() >= j.nextRun.Unix()
}