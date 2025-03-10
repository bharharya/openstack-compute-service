package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bharharya/openstack-compute-service/config"
	"github.com/bharharya/openstack-compute-service/database"
	"github.com/bharharya/openstack-compute-service/routes"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Initialize database connection
	database.InitDB()

	// Setup routes using Gorilla Mux
	router := routes.SetupRoutes()

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Printf("Server running on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
