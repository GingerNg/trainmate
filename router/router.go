package router

import (
	Docs "trainmate/docs"
	"trainmate/middlewares"

	// Router "trainmate/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	apiVer := "/api/v1"
	privateGroup := r.Group(apiVer)
	RegisterUserRouter(privateGroup)
	RegisterDemoRouter(privateGroup)
	RegisterTaskRouter(privateGroup)
	RegisterDatasetRouter(privateGroup)
	RegisterExperimentRouter(privateGroup)
	RegisterJobRouter(privateGroup)

	Docs.SwaggerInfo.BasePath = apiVer

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
