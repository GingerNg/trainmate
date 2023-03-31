package router

import (
	Controller "trainmate/controller"

	"github.com/gin-gonic/gin"
)

func RegisterDemoRouter(Router *gin.RouterGroup) {
	DemoRouter := Router.Group("demo") // .Use(middleware.OperationRecord())
	{
		DemoRouter.GET("hello", Controller.Helloworld)

		DemoRouter.GET("/user/:name", Controller.UserName)

		DemoRouter.GET("/user/:name/*action", Controller.UserNameAction)

		DemoRouter.GET("/welcome", Controller.Welcome)

		DemoRouter.POST(("/formpost"), Controller.FormPost)
	}
}
