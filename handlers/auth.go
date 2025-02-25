package handlers

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "net/http"
)

type User struct {
    gorm.Model
    Username string `json:"username" gorm:"unique"`
    Password string `json:"-"`
    Credits  int    `json:"credits"`
}

func RegisterUser(c *gin.Context) {
    var input User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    user := User{Username: input.Username, Password: string(hashedPassword), Credits: 100}

    database.DB.Create(&user)
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "credits": user.Credits})
}
