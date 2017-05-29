package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorhill/cronexpr"
)

// JobState type for job state enum
type JobState uint8

// job states
const (
	DISABLED JobState = iota
	ENABLED  JobState = iota
	RUNNING  JobState = iota
)

// Job - Struct for the Job Model
type Job struct {
	CronPattern string
	ImageName   string
	RunCommand  string
	NextRun     time.Time
	State       JobState
}

// Jobs - collection cron jobs
type Jobs map[int64][]Job

// Add job to list of jobs
func (j Jobs) Add(job Job) Jobs {

	// generate the next time this should run
	nextTime := cronexpr.MustParse(job.CronPattern).Next(time.Now())
	job.NextRun = nextTime
	timeStamp := nextTime.Unix()

	// need function to figure out next run time
	// hackingin time for now
	if len(j[timeStamp]) == 0 {
		j[timeStamp] = []Job{job}
	} else {
		j[timeStamp] = append(j[timeStamp], job)
	}

	return j
}

// Del  job to list of jobs
func (j *Jobs) Del(key int64) {
	delete(*j, key)
}

// RunList - Get jobs that should run now
func (j *Jobs) RunList(time string) {

}

// List jobs all that are stored
func (j Jobs) List() {
	for _, jobSet := range j {
		for _, job := range jobSet {
			fmt.Println("ImageName", job.ImageName)
			fmt.Println("RunCommand", job.RunCommand)
			fmt.Println("State", job.State)
			fmt.Println("CronPattern", job.CronPattern)
			fmt.Println("NextRun", job.NextRun)
		}
	}

}

// Poll - Polls the list of jobs if anything should be running at this time
func (j *Jobs) Poll(ticker <-chan time.Time) {

	for {
		select {
		case <-ticker:
			k := *j
			currentMinute := time.Now().Truncate(time.Minute).Unix()
			if len(k[currentMinute]) != 0 {
				for _, v := range k[currentMinute] {
					log.Println("Running", v.ImageName)
					// run the job

					// delete the job
					k.Del(currentMinute)
					// add next run to list
					k.Add(v)

				}
			} else {
				log.Println("No jobs to Run")
			}

		}
	}
}

// Run job in job list
func (j *Jobs) Run(job Job) {

}
