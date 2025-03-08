package database

import "time"

// User represents a registered user in the system
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // Omit password in JSON responses
	Credits   int       `json:"credits"`
	CreatedAt time.Time `json:"created_at"`
}

// Instance represents a virtual machine instance
type Instance struct {
	ID         string    `json:"id"`
	UserID     int       `json:"user_id"`
	Name       string    `json:"name"`
	Flavor     string    `json:"flavor"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
