package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

type education struct {
	University string `json:"university"`
	City       string `json:"city"`
	Field      string `json:"field"`
	Degree     string `json:"degree"`
}

type experience struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	City      string `json:"city"`
	Timeframe string `json:"timeframe"`
}

func getExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	experienceJSON, err := json.Marshal(loadExperience())
	if err != nil {
		log.Fatal(err)
	}

	w.Write(experienceJSON)
}
func getEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	educationJSON, err := json.Marshal(loadEducation())
	if err != nil {
		log.Fatal(err)
	}

	w.Write(educationJSON)
}
func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

func loadEducation() education {
	education := education{
		"University of applied science",
		"Leipzig, Germany",
		"Information technology",
		"Bachelor of Science",
	}

	return education
}

func loadExperience() []experience {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var expirences []experience

	iter := client.Collection("resume").Doc("JrEvSIoWiSgTgXQRIC6I").Collection("experience").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var e experience
		doc.DataTo(&e)

		expirences = append(expirences, e)
	}

	return expirences
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/resume/v1").Subrouter()
	api.HandleFunc("/experience/", getExperience).Methods(http.MethodGet)
	api.HandleFunc("/education/", getEducation).Methods(http.MethodGet)
	api.HandleFunc("", notImplemented)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "karrlein" // os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
