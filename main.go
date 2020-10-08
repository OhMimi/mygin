package main

import (
	"fmt"
	"log"
	_ "mygin/docs"
	"mygin/router"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @contact.name hulk.hsu

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// schemes http

func main() {
	port := 8080
	router := router.SetupRouter(port)
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
