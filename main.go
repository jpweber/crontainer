package main

import (
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db       *bolt.DB
	cronJobs Jobs
)

func main() {

	log.Println("Starting Crontainer")

	// Init Database
	db = openDB("job-data.db")
	initBucket(db, "jobs")

	// Init Jobs list
	cronJobs = Jobs{0: []Job{}}

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
	go cronJobs.Poll(tickChan)

	mux := http.NewServeMux()
	mux.HandleFunc("/", logger(serve))
	http.ListenAndServe(":8675", mux)
}
