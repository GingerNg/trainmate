package main

import (
	"trainmate/drivers"
	initialize "trainmate/initialize"
)

func main() {
	drivers.Init()
	router := initialize.RegisterRouter()
	router.Run(":8000")
}
