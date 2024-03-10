package handlers

import (
	"auth-service/pkg/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditPersonalData() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := controllers.EditPersonalData(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Personal data updated successfully"})
	}
}
