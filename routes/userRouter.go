package routes

import (
	controller "github.com/Venkatakarthik0211/library-management-backend/controllers"
	"github.com/Venkatakarthik0211/library-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incoming *gin.Engine) {
	incoming.Use(middleware.Authenticate())
	incoming.GET("users/", controller.GetUsers())
	incoming.POST("users/:user_id", controller.GetUsers())

}
