package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db *bolt.DB
)

func main() {

	// Init Database
	db = openDB("job-data.db")
	initBucket(db, "jobs")

	// Init Jobs list
	cronJobs := Jobs{0: []Job{}}

	// import any job data from db
	importDB(db, cronJobs)

	// test job data
	// testJob := Job{
	// 	CronPattern: "*/1 * * * *",
	// 	ImageName:   "jpweber/crontest:0.1.0",
	// 	State:       ENABLED,
	// }
	// testJob.encodeHash()
	// testJob2 := Job{
	// 	CronPattern: "*/5 * * * *",
	// 	ImageName:   "jpweber/crontest2:0.1.0",
	// 	State:       DISABLED,
	// }
	// testJob2.encodeHash()

	// add job to list of jobs
	// cronJobs.Add(testJob)
	// writeToDB(db, testJob, "jobs")
	// cronJobs.Add(testJob2)
	// writeToDB(db, testJob2, "jobs")

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
