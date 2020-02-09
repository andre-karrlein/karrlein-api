package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Education struct {
	University string `json:"university"`
	City       string `json:"city"`
	Field      string `json:"field"`
	Degree     string `json:"degree"`
}

type Experience struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	City      string `json:"city"`
	Timeframe string `json:"timeframe"`
}

func getExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	experienceJson, _ := json.Marshal(loadExperience())

	w.Write(experienceJson)
}
func getEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	educationJson, _ := json.Marshal(loadEducation())

	w.Write(educationJson)
}
func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

func loadEducation() Education {
	education := Education{
		"University of applied science",
		"Leipzig, Germany",
		"Information technology",
		"Bachelor of Science",
	}

	return education
}

func loadExperience() []Experience {
	experience := Experience{
		"Software Developer",
		"FLYERALARM",
		"Wuerzburg",
		"August 2016 - October 2017",
	}
	experience2 := Experience{
		"Lead Developer",
		"FLYERALARM",
		"Wuerzburg",
		"October 2017 - June 2019",
	}
	experience3 := Experience{
		"Senior Software Engineer",
		"FLYERALARM",
		"Wuerzburg",
		"June 2019 - today",
	}
	resume := []Experience{
		experience,
		experience2,
		experience3,
	}

	return resume
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/resume/v1").Subrouter()
	api.HandleFunc("/experience/", getExperience).Methods(http.MethodGet)
	api.HandleFunc("/education/", getEducation).Methods(http.MethodGet)
	api.HandleFunc("", notImplemented)

	log.Fatal(http.ListenAndServe(":8080", r))
}
