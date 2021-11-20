package router

import (
	Controller "trainmate/controller"

	"github.com/gin-gonic/gin"
)

func InitTaskRouter(Router *gin.RouterGroup) {
	TaskRouter := Router.Group("task") // .Use(middleware.OperationRecord())
	{
		TaskRouter.POST("insert", Controller.CreateTask)
		TaskRouter.GET("query", Controller.QueryTask)
		TaskRouter.GET("queries", Controller.QueryTasks)
		TaskRouter.POST("delete", Controller.DeleteTask)
	}
}
