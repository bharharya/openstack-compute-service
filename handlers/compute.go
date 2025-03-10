package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// ListInstances retrieves all instances
func ListInstances(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "List of instances"})
}

// CreateInstance creates a new instance
func CreateInstance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Instance created"})
}

// StopInstance stops a running instance
func StopInstance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(map[string]string{"message": "Instance stopped", "id": params["id"]})
}

// DeleteInstance deletes an instance
func DeleteInstance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(map[string]string{"message": "Instance deleted", "id": params["id"]})
}
