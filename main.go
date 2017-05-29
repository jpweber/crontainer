package main

import (
	"log"
	"time"
)

func checkJobs() {

}

func main() {

	// Init Jobs list
	cronJobs := Jobs{0: []Job{}}
	// test job data
	testJob := Job{
		CronPattern: "*/1 * * * *",
		ImageName:   "jpweber/crontest:0.1.0",
		State:       ENABLED,
	}
	testJob2 := Job{
		CronPattern: "*/5 * * * *",
		ImageName:   "jpweber/crontest2:0.1.0",
		State:       DISABLED,
	}

	// add job to list of jobs
	cronJobs.Add(testJob)
	// cronJobs.List()
	cronJobs.Add(testJob2)
	cronJobs.List()
	log.Println("Initial Jobs Loaded Starting Service")

	// start the 1 minute ticker that looks for jobs
	tickChan := time.NewTicker(time.Second * 1).C

	// make a channel to keep us running until exited
	done := make(chan bool, 1)
	go cronJobs.Poll(tickChan)

	// waiting until done = true
	<-done

}
