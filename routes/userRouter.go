package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoaamiewithane/go-jwt/middleware"

	"github.com/hoaamiewithane/go-jwt/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("users", controllers.GetUsers())
	incomingRoutes.POST("users/:user_id", controllers.GetUser())
}
