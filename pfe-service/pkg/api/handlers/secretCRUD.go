package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pfe-service/pkg/shamir"
)

type SaveResponse struct {
	Message string `json:"message"`
}

type SaveRequest struct {
	Part shamir.ShamirPart `json:"part"`
}

var part = shamir.ShamirPart{}

func HandleSavePart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body := SaveRequest{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	part = body.Part
	// return ok
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SaveResponse{Message: "Part saved"})
}

func HandleGetSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SaveRequest{Part: part})
}
