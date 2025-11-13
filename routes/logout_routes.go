package routes

import (
	"dinsos_kuburaya/controllers"

	"github.com/gin-gonic/gin"
)

func LogoutRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/logout", controllers.Logout)
	}
}
