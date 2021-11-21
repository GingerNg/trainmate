package main

import (
	"trainmate/drivers"
	Router "trainmate/router"
)

func main() {
	drivers.Init()
	router := Router.RegisterRouter()
	router.Run(":8666")
}
