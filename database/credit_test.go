package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup in-memory database for testing
func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&Credit{})
	return db
}

func TestInitializeCredits(t *testing.T) {
	db = setupTestDB() // Use in-memory DB
	userID := uint(1)
	initialBalance := 100.0

	err := InitializeCredits(userID, initialBalance)
	assert.NoError(t, err)

	var credit Credit
	err = db.Where("user_id = ?", userID).First(&credit).Error
	assert.NoError(t, err)
	assert.Equal(t, initialBalance, credit.Balance)
}

func TestDeductCredits_Success(t *testing.T) {
	db = setupTestDB()
	userID := uint(1)
	_ = InitializeCredits(userID, 100.0)

	err := DeductCredits(userID, 20.0)
	assert.NoError(t, err)

	balance, _ := GetCredits(userID)
	assert.Equal(t, 80.0, balance)
}

func TestDeductCredits_InsufficientFunds(t *testing.T) {
	db = setupTestDB()
	userID := uint(1)
	_ = InitializeCredits(userID, 10.0)

	err := DeductCredits(userID, 20.0)
	assert.Error(t, err)
	assert.Equal(t, "insufficient credits", err.Error())
}

func TestRefundCredits(t *testing.T) {
	db = setupTestDB()
	userID := uint(1)
	_ = InitializeCredits(userID, 50.0)

	err := RefundCredits(userID, 20.0)
	assert.NoError(t, err)

	balance, _ := GetCredits(userID)
	assert.Equal(t, 70.0, balance)
}

func TestGetCredits(t *testing.T) {
	db = setupTestDB()
	userID := uint(1)
	_ = InitializeCredits(userID, 200.0)

	balance, err := GetCredits(userID)
	assert.NoError(t, err)
	assert.Equal(t, 200.0, balance)
}
