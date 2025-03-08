package handlers

import (
	"log"
	"net/http"

	"github.com/bharharya/openstack-compute-service/database"
	"github.com/gin-gonic/gin"
)

// CheckCreditsHandler checks the current credit balance of a user
func CheckCreditsHandler(c *gin.Context) {
	userID := c.Param("user_id")

	var credits int
	err := database.DB.QueryRow("SELECT credits FROM users WHERE id = $1", userID).Scan(&credits)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user credits"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "credits": credits})
}

// DeductCredits deducts credits based on instance usage
func DeductCredits(userID int, amount int) error {
	_, err := database.DB.Exec("UPDATE users SET credits = credits - $1 WHERE id = $2 AND credits >= $1", amount, userID)
	if err != nil {
		log.Println("Failed to deduct credits:", err)
		return err
	}
	return nil
}

// AddCreditsHandler allows an admin to add credits to a user
func AddCreditsHandler(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id"`
		Amount int `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err := database.DB.Exec("UPDATE users SET credits = credits + $1 WHERE id = $2", req.Amount, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add credits"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Credits added successfully"})
}

// RefundCredits refunds credits if an instance is deleted
func RefundCredits(userID int, amount int) error {
	_, err := database.DB.Exec("UPDATE users SET credits = credits + $1 WHERE id = $2", amount, userID)
	if err != nil {
		log.Println("Failed to refund credits:", err)
		return err
	}
	return nil
}
