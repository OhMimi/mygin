package main

import "mygin/router"

func main() {
	router := router.SetupRouter()
	router.Run()
}
