package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// URLElements - struct for holding parts of the url
// we are going to reference in a tidy way
type URLElements struct {
	resourceType string
	resource     string
}

// wrapper function for http logging
func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer log.Printf("%s - %s - %s", r.Host, r.Method, r.URL)
		fn(w, r)
	}
}

func parseURL(url string) URLElements {
	urlParts := strings.Split(url, "/")
	elements := URLElements{}
	if len(urlParts) > 1 {
		elements.resourceType = urlParts[1]
	}
	if len(urlParts) > 2 {
		elements.resource = urlParts[2]
	}
	// ditch index 0 because its always blank
	return elements
}

func get(url string, w http.ResponseWriter) {

	urlElements := parseURL(url)
	data := []byte{}

	switch urlElements.resourceType {
	case "schedule":
		data = cronJobs.List()
	case "jobs":
		data = jobList.List()
	case "job":
		data = jobList.Get(urlElements.resource)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func httpDel(w http.ResponseWriter, r *http.Request) {

	// currently we only support deleting jobs
	urlElements := parseURL(r.URL.String())
	jobList.Del(urlElements.resource)
	delFromDB(db, urlElements.resource, "jobs")

}

func post(w http.ResponseWriter, r *http.Request) {
	// parse the form
	err := r.ParseForm()
	if err != nil {
		status := http.StatusInternalServerError
		log.Println("form parse problem " + strconv.Itoa(status))
	}
	payload := new(Job)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		log.Println(err)
	}

	payload.encodeHash()
	cronJobs.Add(*payload)
	writeToDB(db, *payload, "jobs")
	w.WriteHeader(http.StatusOK)
}

func serve(w http.ResponseWriter, r *http.Request) {

	// determine function based on http method
	switch r.Method {
	case "GET":
		get(r.URL.String(), w)
	case "POST":
		post(w, r)
	case "DELETE":
		httpDel(w, r)
	default:
		fmt.Println(r.Method + "HTTP method not implemented")
	}
}
