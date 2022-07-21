package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	type appStatus struct {
		Status      string `json:"status"`
		Environment string `json:"environment"`
	}

	currentStatus := appStatus{
		Status: os.Getenv("STATUS"),
		Environment: os.Getenv("ENV"),
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		log.Panicln("Error parsin")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
