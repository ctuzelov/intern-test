package pkg

import (
	"auth-service/pkg/api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

// Function that handles the initialization of the project
func Run() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.Default()

	routes.UserRoutes(router)
	routes.ProjectRoutes(router)

	router.Run(":" + port)
}
