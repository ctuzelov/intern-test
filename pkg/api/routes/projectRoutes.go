package routes

import (
	"auth-service/pkg/api/handlers"
	"auth-service/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/create-project", handlers.CreateProject()).Use(middleware.Authenticate())
	incomingRoutes.PUT("/update-project/:id", handlers.UpdateProject()).Use(middleware.IsAdmin())
}
