package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"strings"
)

// JobState type for job state enum
type JobState uint8

// job states
const (
	DISABLED JobState = iota
	ENABLED  JobState = iota
)

// Job - Struct for the Job Model
type Job struct {
	CronPattern string
	ImageName   string
	RunCommand  []string
	NextRun     int64
	State       JobState
	Hash        string
}

func (j *Job) encodeHash() {
	var hBuf bytes.Buffer
	// create new Sha1 hash
	h := sha1.New()
	// concat struct values into a string
	hBuf.WriteString(j.ImageName)
	hBuf.WriteString(strings.Join(j.RunCommand, ","))
	hString := hBuf.String()
	// create byte slice of struct parts string
	h.Write([]byte(hString))
	// save sha1 sum
	bs := h.Sum(nil)
	j.Hash = fmt.Sprintf("%x", bs)

}
