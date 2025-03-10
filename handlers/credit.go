package handlers

import (
	"encoding/json"
	"net/http"
)

// GetCredit retrieves the user's current credits
func GetCredit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"credits": 100})
}

// UpdateCredit updates the user's credit balance
func UpdateCredit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Credits updated"})
}
