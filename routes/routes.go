package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/bharharya/openstack-compute-service/handlers"
	"github.com/bharharya/openstack-compute-service/middleware"
)

// SetupRoutes initializes all API routes
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Authentication routes
	router.HandleFunc("/login", handlers.Login).Methods("POST")               // ✅ Corrected function name
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")     // ✅ Corrected function name

	// Middleware for protected routes
	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Compute instance routes
	protectedRoutes.HandleFunc("/instances", handlers.ListInstances).Methods("GET")        // ✅ Ensure ListInstances exists
	protectedRoutes.HandleFunc("/instances", handlers.CreateInstance).Methods("POST")      // ✅ Ensure CreateInstance exists
	protectedRoutes.HandleFunc("/instances/{id}", handlers.StopInstance).Methods("PUT")    // ✅ Ensure StopInstance exists
	protectedRoutes.HandleFunc("/instances/{id}", handlers.DeleteInstance).Methods("DELETE") // ✅ Ensure DeleteInstance exists

	// Credit system routes
	protectedRoutes.HandleFunc("/credits", handlers.GetCredit).Methods("GET")          // ✅ Ensure GetCredit exists
	protectedRoutes.HandleFunc("/credits", handlers.UpdateCredit).Methods("PUT")       // ✅ Ensure UpdateCredit exists

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	return router
}
