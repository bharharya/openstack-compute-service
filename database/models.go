package database

import "gorm.io/gorm"

// User represents a registered user in the system
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Credits  int    `gorm:"default:100"` // Initial credits for new users
}

// Instance represents a virtual machine instance
type Instance struct {
	gorm.Model
	UserID     uint   `gorm:"not null"` // Foreign key to User
	Name       string `gorm:"not null"`
	InstanceID string `gorm:"unique;not null"`
	Status     string `gorm:"default:'pending'"`
}
