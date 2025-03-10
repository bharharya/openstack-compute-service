package database

import (
	"errors"
	"gorm.io/gorm"
)

// DeductCredits deducts credits from a user based on instance runtime
func DeductCredits(db *gorm.DB, userID uint, amount int) error {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Credits < amount {
		return errors.New("insufficient credits")
	}

	user.Credits -= amount
	return db.Save(&user).Error
}

// AddCredits adds credits to a user
func AddCredits(db *gorm.DB, userID uint, amount int) error {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	user.Credits += amount
	return db.Save(&user).Error
}
