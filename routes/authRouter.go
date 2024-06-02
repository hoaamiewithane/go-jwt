package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoaamiewithane/go-jwt/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("sign-up", controllers.Signup())
	incomingRoutes.POST("login", controllers.Login())
}
