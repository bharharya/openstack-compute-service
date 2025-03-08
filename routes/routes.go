package routes

import (
	"github.com/bharharya/openstack-compute-service/handlers"
	"github.com/bharharya/openstack-compute-service/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes all routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Apply JWT authentication middleware to protected routes
	authMiddleware := middleware.AuthMiddleware()

	// Public Routes (Authentication)
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.Login)

	// Compute Service (Requires JWT Auth)
	compute := r.Group("/compute")
	compute.Use(authMiddleware)
	{
		compute.POST("/create", handlers.CreateInstanceHandler)                // Create a new VM
		compute.POST("/stop/:instance_id", handlers.StopInstanceHandler)       // Stop VM
		compute.DELETE("/delete/:instance_id", handlers.DeleteInstanceHandler) // Delete VM
		compute.GET("/list", handlers.ListUserInstancesHandler)                // Get user's VMs
	}

	// Credit System (Requires JWT Auth)
	credit := r.Group("/credit")
	credit.Use(authMiddleware)
	{
		credit.GET("/balance", handlers.GetCreditBalanceHandler) // Check credit balance
		credit.POST("/deduct", handlers.DeductCreditHandler)     // Deduct credits
		credit.POST("/refund", handlers.RefundCreditHandler)     // Refund credits
	}

	return r
}
