package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/bharharya/openstack-compute-service/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/register", handlers.RegisterUser)
    r.POST("/login", handlers.Login)
    r.POST("/compute/create", handlers.CreateInstanceHandler)

    return r
}
