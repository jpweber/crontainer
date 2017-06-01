package main

import (
	"encoding/json"
	"log"
)

// JobList - list of all cron jobs keyed via their hash digest
type JobList map[string]Job

// Add job to list of jobs
func (j *JobList) Add(job Job) {

	k := *j
	k[job.Hash] = job
	*j = k
}

// Del - delete a job from list of jobs
func (j *JobList) Del(key string) {
	k := *j
	job := k[key]
	delete(*j, key)
	// when we delete a job list we must always delete
	// the job from the schedule as well

	cronJobs.Del(job.NextRun, job.Hash)
}

// List all the values in the job list
func (j *JobList) List() []byte {
	encoded, err := json.Marshal(j)
	if err != nil {
		log.Println("Json encoding error in RunList:", err)
	}

	return encoded
}

// Get a single job from the job list
func (j JobList) Get(digest string) []byte {
	encoded, err := json.Marshal(j[digest])
	if err != nil {
		log.Println("Json encoding error in RunList:", err)
	}

	return encoded
}
