package main

import (
	"trainmate/drivers"
	"trainmate/initialize"
	Router "trainmate/router"
)

func main() {
	drivers.Init()
	initialize.InitConf()
	initialize.InitORM()
	initialize.InitHandler()
	initialize.InitService()
	router := Router.RegisterRouter()
	router.Run(":8666")
}
