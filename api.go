package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// wrapper function for http logging
func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer log.Printf("%s - %s - %s", r.Host, r.Method, r.URL)
		fn(w, r)
	}
}

func get(url string, w http.ResponseWriter) {

	data := cronJobs.List()

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
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
	default:
		fmt.Println(r.Method + "HTTP method not implemented")
	}
}
