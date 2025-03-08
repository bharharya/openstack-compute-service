package database

import (
	"errors"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

// Credit represents the credit balance for each user
type Credit struct {
	ID      uint    `gorm:"primaryKey"`
	UserID  uint    `gorm:"uniqueIndex"` // Ensure each user has only one credit record
	Balance float64 `gorm:"not null"`    // Available credit balance
}

// Mutex to prevent race conditions when modifying credits
var creditMutex sync.Mutex

// InitializeCredits initializes the credit entry for a new user
func InitializeCredits(db *gorm.DB, userID uint, initialBalance float64) error {
	credit := Credit{
		UserID:  userID,
		Balance: initialBalance,
	}
	if err := db.Create(&credit).Error; err != nil {
		return fmt.Errorf("failed to initialize credits: %v", err)
	}
	return nil
}

// GetCreditBalance retrieves the current credit balance for a user
func GetCreditBalance(db *gorm.DB, userID uint) (float64, error) {
	var credit Credit
	if err := db.Where("user_id = ?", userID).First(&credit).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch credit balance: %v", err)
	}
	return credit.Balance, nil
}

// DeductCredits deducts a given amount from the user's credit balance
func DeductCredits(db *gorm.DB, userID uint, amount float64) error {
	creditMutex.Lock()
	defer creditMutex.Unlock()

	var credit Credit
	if err := db.Where("user_id = ?", userID).First(&credit).Error; err != nil {
		return fmt.Errorf("failed to find credit record: %v", err)
	}

	if credit.Balance < amount {
		return errors.New("insufficient credits")
	}

	credit.Balance -= amount
	if err := db.Save(&credit).Error; err != nil {
		return fmt.Errorf("failed to deduct credits: %v", err)
	}
	return nil
}

// RefundCredits adds a specified amount back to the user's credit balance
func RefundCredits(db *gorm.DB, userID uint, amount float64) error {
	creditMutex.Lock()
	defer creditMutex.Unlock()

	var credit Credit
	if err := db.Where("user_id = ?", userID).First(&credit).Error; err != nil {
		return fmt.Errorf("failed to find credit record: %v", err)
	}

	credit.Balance += amount
	if err := db.Save(&credit).Error; err != nil {
		return fmt.Errorf("failed to refund credits: %v", err)
	}
	return nil
}
