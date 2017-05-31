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
