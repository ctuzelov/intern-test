package routes

import (
	"auth-service/pkg/api/handlers"
	"auth-service/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/create-project", middleware.Authenticate(), handlers.CreateProject())
	incomingRoutes.PUT("/update-project/:id", middleware.IsAdmin(), handlers.UpdateProject())
	incomingRoutes.GET("/get-projects", handlers.GetAllProjects())
}
