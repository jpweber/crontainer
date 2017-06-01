package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gorhill/cronexpr"
)

// Schedule - collection cron jobs
type Schedule map[int64][]Job

// Add job to list of jobs
func (j Schedule) Add(job Job) Schedule {

	// generate the next time this should run
	nextTime := cronexpr.MustParse(job.CronPattern).Next(time.Now())
	timeStamp := nextTime.Unix()
	job.NextRun = timeStamp

	// add job to schedule.
	// we have to check if there are any jobs in this time slot or not
	if len(j[timeStamp]) == 0 {
		j[timeStamp] = []Job{job}
	} else {
		j[timeStamp] = append(j[timeStamp], job)
	}

	// add the job to the JobList
	jobList.Add(job)

	return j
}

// Del - delete job from the list
func (j *Schedule) Del(key int64, digest string) {
	// delete(*j, key)
	k := *j
	for idx, v := range k[key] {
		if v.Hash == digest {
			log.Println("Deleting Job", v.Hash, "From Timestamp", v.NextRun)
			// if we have a match, copy the value in the last position
			// to the value in our matched position
			// and truncate the slice. Order does not matter so we can do this
			k[key][idx] = k[key][len(k[key])-1]
			k[key] = k[key][:len(k[key])-1]
		}
	}

	// if our timestamp is empty now delete that key from the list
	if len(k[key]) == 0 {
		delete(k, key)
	}
	// set our schedule via the pointer the values
	// of our copy in the function
	*j = k
}

// List jobs all that are stored
func (j Schedule) List() []byte {

	encoded, err := json.Marshal(j)
	if err != nil {
		log.Println("Json encoding error in RunList:", err)
	}

	return encoded

}

// Poll - Polls the list of jobs if anything should be running at this time
func (j *Schedule) Poll(ticker <-chan time.Time) {

	for {
		select {
		case <-ticker:
			k := *j // can't iterate over pointer. Copy to value internal to function
			// ticker runs every second. Truncate to minute precision
			// effectively always round down.
			currentMinute := time.Now().Truncate(time.Minute).Unix()
			if len(k[currentMinute]) != 0 {
				for _, v := range k[currentMinute] {
					log.Println("Running", v.ImageName)
					// run the job
					runContainer(v)
					// delete the job
					k.Del(currentMinute, v.Hash)
					// add next run to list
					go k.Add(v)

				}
			}

		}
	}
}

// Run job in job list
func (j *Schedule) Run(job Job) {
	// a lot of this depends on environment
	// where the containers are being ran

}

func importDB(db *bolt.DB, jobs Schedule) {
	db.View(func(tx *bolt.Tx) error {
		job := Job{}
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("jobs"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := json.Unmarshal(v, &job); err != nil {
				panic(err)
			}
			jobs.Add(job)
		}

		return nil
	})
}
