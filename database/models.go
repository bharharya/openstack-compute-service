package database

import (
    "gorm.io/gorm"
)

// User model
type User struct {
    gorm.Model
    Username string `json:"username" gorm:"unique"`
    Password string `json:"-"`
    Credits  int    `json:"credits"`
}

// Instance model (maps to OpenStack instances)
type Instance struct {
    gorm.Model
    UserID   uint   `json:"user_id"`
    Name     string `json:"name"`
    FlavorID string `json:"flavor_id"`
    ImageID  string `json:"image_id"`
    NetworkID string `json:"network_id"`
    Status   string `json:"status"` // Active, Shutoff, etc.
}

// Migrate models
func MigrateDB() {
    DB.AutoMigrate(&User{}, &Instance{})
}
