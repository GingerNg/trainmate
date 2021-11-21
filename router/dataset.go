package router

import (
	Controller "trainmate/controller"

	"github.com/gin-gonic/gin"
)

func RegisterDatasetRouter(Router *gin.RouterGroup) {
	TaskRouter := Router.Group("dataset") // .Use(middleware.OperationRecord())
	{
		TaskRouter.POST("insert", Controller.CreateDataset)
		TaskRouter.GET("query", Controller.QueryDataset)
		TaskRouter.GET("queries", Controller.QueryDatasets)
		TaskRouter.POST("delete", Controller.DeleteDataset)
	}
}
