package router

import (
	Controller "trainmate/controller"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("vue-element-admin/user") // .Use(middleware.OperationRecord())
	{
		UserRouter.POST("refresh", Controller.UserLogin)
		UserRouter.POST("logout", Controller.UserLogout)
		UserRouter.POST("hello", Controller.SayHello)
		// ************
		UserRouter.POST("login", Controller.UserLogin)
		UserRouter.OPTIONS("login", Controller.UserLogin)
		UserRouter.GET("info", Controller.UserInfo)
		UserRouter.OPTIONS("info", Controller.UserInfo)

	}
}
