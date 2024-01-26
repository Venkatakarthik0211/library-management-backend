package routes

import (
	controller "github.com/Venkatakarthik0211/library-management-backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incoming *gin.Engine) {
	incoming.POST("users/signup", controller.Signup())
	incoming.POST("users/login", controller.Login())

}
