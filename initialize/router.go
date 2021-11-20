package initialize

import (
	Docs "trainmate/docs"
	"trainmate/middlewares"
	Router "trainmate/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	apiVer := "/api/v1"
	privateGroup := r.Group(apiVer)
	Router.InitUserRouter(privateGroup)
	// Router.InitDemoRouter(privateGroup)

	Router.InitTaskRouter(privateGroup)
	Router.InitDatasetRouter(privateGroup)
	Router.InitExperimentRouter(privateGroup)
	Router.InitJobRouter(privateGroup)

	Docs.SwaggerInfo.BasePath = apiVer

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
