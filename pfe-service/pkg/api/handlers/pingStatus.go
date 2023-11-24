package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pfe-service/config"
	"github.com/pfe-service/pkg/models"
)

type pingStatusResponse struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

func HandlePingStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// return json map with status of the service
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := pingStatusResponse{
		Status: models.ServiceUp.String(),
		Name:   config.GetConfig().Name,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonResponse)
}
