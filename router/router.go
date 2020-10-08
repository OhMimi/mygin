package router

import (
	"fmt"
	"mygin/handler"

	swaggerFiles "github.com/swaggo/files"

	gin "github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(port int) *gin.Engine {
	router := gin.Default()

	if mode := gin.Mode(); mode == gin.DebugMode {
		config := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port))
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, config))
	}

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
