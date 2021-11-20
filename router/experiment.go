package router

import (
	Controller "trainmate/controller"

	"github.com/gin-gonic/gin"
)

func InitExperimentRouter(Router *gin.RouterGroup) {
	ExpRouter := Router.Group("experiment") // .Use(middleware.OperationRecord())
	{
		ExpRouter.POST("insert", Controller.CreateExp)
		// JobRouter.POST("update", Controller.UpdateJob)
		ExpRouter.GET("query", Controller.QueryExp)
		ExpRouter.GET("queries", Controller.QueryExps)
		ExpRouter.POST("delete", Controller.DeleteExp)

	}
}
