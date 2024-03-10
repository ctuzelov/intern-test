package routes

import (
	"auth-service/pkg/api/handlers"
	"auth-service/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", handlers.Signup())
	incomingRoutes.POST("/signin", handlers.Login())
	incomingRoutes.POST("/edit", handlers.UpdateUser()).Use(middleware.Authenticate())
	incomingRoutes.POST("/refresh-token", handlers.RefreshToken())
}
