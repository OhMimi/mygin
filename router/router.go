package router

import (
	"mygin/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../view/*")
	} else {
		router.LoadHTMLGlob("view/*")
	}

	// static resources

	router.Static("/assetPath", "asset")

	// router table

	pingRouting := router.Group("/ping")
	{
		pingRouting.GET("", handler.WsPing)
	}

	indexRouting := router.Group("/")
	{
		indexRouting.GET("", handler.GetIndex)
	}

	fileRouting := router.Group("/file")
	{
		fileRouting.POST("/upload", handler.UploadSingleIndex)
	}

	helloRouting := router.Group("/hello")
	{
		helloRouting.GET("", handler.GetHello)
		helloRouting.DELETE("/:id", handler.DeleteHello)
	}

	userRouting := router.Group("/user")
	{
		userRouting.GET("/:uid", handler.GetUser)
		userRouting.POST("/login", handler.UserLogin)
	}

	return router
}
