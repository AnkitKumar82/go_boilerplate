package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthGet(w http.ResponseWriter, req *http.Request) {
	type HealthCheck struct {
		Message string `json:"message"`
	}

	healthCheckResponse := HealthCheck{
		Message: "health check successfull",
	}

	jsonData, err := json.Marshal(healthCheckResponse)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonData)
}
