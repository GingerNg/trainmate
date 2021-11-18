package router

import (
	Controller "trainmate/controller"

	"github.com/gin-gonic/gin"
)

func InitExperimentRouter(Router *gin.RouterGroup) {
	JobRouter := Router.Group("experiment") // .Use(middleware.OperationRecord())
	{
		JobRouter.POST("insert", Controller.CreateExp)

		// JobRouter.POST("update", Controller.UpdateJob)

		JobRouter.GET("query", Controller.QueryExps)

	}
}
