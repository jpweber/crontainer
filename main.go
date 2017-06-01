package main

import (
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db       *bolt.DB
	cronJobs Schedule
	jobList  JobList
)

func main() {

	log.Println("Starting Crontainer")

	// Init Database
	db = openDB("job-data.db")
	initBucket(db, "jobs")

	// Init Jobs list
	cronJobs = Schedule{0: []Job{}}
	jobList = make(JobList)

	// import any job data from db
	importDB(db, cronJobs)

	cronJobs.List()
	log.Println("Initial Jobs Loaded Starting Service")

	// start the 1 minute ticker that looks for jobs
	tickChan := time.NewTicker(time.Second * 1).C
	go cronJobs.Poll(tickChan)

	mux := http.NewServeMux()
	mux.HandleFunc("/", logger(serve))
	http.ListenAndServe(":8675", mux)
}
