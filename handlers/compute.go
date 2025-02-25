package handlers

import (
	"net/http"

	"github.com/bharharya/openstack-compute-service/openstack"
	"github.com/gin-gonic/gin"
)

func CreateInstanceHandler(c *gin.Context) {
	var request struct {
		Name      string `json:"name"`
		FlavorID  string `json:"flavor_id"`
		ImageID   string `json:"image_id"`
		NetworkID string `json:"network_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	provider, _ := openstack.GetOpenStackClient("AUTH_URL", "USERNAME", "PASSWORD")

	instance, err := openstack.CreateInstance(provider, request.Name, request.FlavorID, request.ImageID, request.NetworkID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Instance created", "instance_id": instance.ID})
}
