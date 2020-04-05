package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
	user := os.Getenv("DBUSERNAME")
	password := os.Getenv("DBPASSWORD")
	host := os.Getenv("DBHOST")
	dbname := os.Getenv("DBNAME")

	connection := user + ":" + password + "@(" + host + ":3306)/" + dbname + "?parseTime=true&charset=utf8"

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT role, company, city, timeframe FROM resume`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var expirences []experience
	for rows.Next() {
		var e experience
		err = rows.Scan(&e.Role, &e.Company, &e.City, &e.Timeframe)
		if err != nil {
			log.Fatal(err)
		}
		expirences = append(expirences, e)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
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
