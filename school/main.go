package main

import (
	"school/initialize"
	"school/router"
)

func main() {
	initialize.DB()
	initialize.InitLogger()
	router.Router().Run(":8080")
}
