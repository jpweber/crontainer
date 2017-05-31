package main

import (
	"testing"
)

// TestAdd - test adding a job to a fresh list of jobs
func TestAdd(t *testing.T) {

	// Init Jobs list
	cronJobs := Jobs{0: []Job{}}

	testJob := Job{
		CronPattern: "*/1 * * * *",
		ImageName:   "jpweber/crontest:0.1.0",
		State:       ENABLED,
	}

	cronJobs.Add(testJob)

	if len(cronJobs) != 2 {
		t.Error("Expected 2 jobs in jobs list got ", len(cronJobs), "instead")
	}
}

// TestEncodeHash - test job hashing method
func TestEncodeHash(t *testing.T) {

	testJob := Job{
		CronPattern: "*/1 * * * *",
		ImageName:   "jpweber/crontest:0.1.0",
		State:       ENABLED,
	}

	testJob.encodeHash()

	if testJob.Hash != "0b9d40033bbc6b5290341d389e59287de04d7637" {
		t.Error("Job Hash is not correct. Got", testJob.Hash, "Expected 0b9d40033bbc6b5290341d389e59287de04d7637")
	}
}

// TestDel - test deleting a job from list of jobs
func TestDel(t *testing.T) {

	// Init Jobs list
	cronJobs := Jobs{0: []Job{}}

	testJob := Job{
		CronPattern: "*/1 * * * *",
		ImageName:   "jpweber/crontest:0.1.0",
		State:       ENABLED,
	}
	testJob.encodeHash()

	cronJobs.Add(testJob)

	cronJobs.Del(0)
	testData := cronJobs[0]
	if len(testData) != 0 {
		t.Error("Job was not deleted.")
	}
}
