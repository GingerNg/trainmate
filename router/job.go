package router

import (
	Controller "trainmate/controller"

	"github.com/gin-gonic/gin"
)

func InitJobRouter(Router *gin.RouterGroup) {
	JobRouter := Router.Group("job") // .Use(middleware.OperationRecord())
	{
		JobRouter.POST("insert", Controller.CreateJob)

		JobRouter.POST("update", Controller.UpdateJob)

		JobRouter.GET("query", Controller.QueryJob)

		// JobRouter.OPTIONS("querybyexpid", Controller.QueryJobByExp)
		JobRouter.GET("querybyexpid", Controller.QueryJobByExp)

		JobRouter.POST("delete", Controller.DeleteJob)

		// DemoRouter.GET("/user/:name/*action", Controller.UserNameAction)

		// DemoRouter.GET("/welcome", Controller.Welcome)

		// DemoRouter.POST(("/formpost"), Controller.FormPost)
	}
}
