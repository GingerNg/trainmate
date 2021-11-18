package router

import (
	Controller "trainmate/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("vue-element-admin/user") // .Use(middleware.OperationRecord())
	{
		UserRouter.POST("login", Controller.UserLogin)
		UserRouter.OPTIONS("login", Controller.UserLogin)

		// JobRouter.POST("update", Controller.UpdateJob)

		UserRouter.GET("info", Controller.UserInfo)
		UserRouter.OPTIONS("info", Controller.UserInfo)

	}
}
