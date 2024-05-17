package main

import (
	"basic-crud/config"
	"basic-crud/router"
)

func main() {
	config.InitDB()

	e := router.New()

	e.Start(":9000")
}
