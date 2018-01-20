package main

import (
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

// Gin Router, listen defined endpoints
func networkHandler(wardrobe *Wardrobe) *gin.Engine {
	engine := gin.Default()
	engine.MaxMultipartMemory = 8 << 20

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.POST("/v1/uploadTest", wardrobe.uploadTest)
	engine.POST("/v1/test", wardrobe.runTest)
	engine.GET("/v1/test", wardrobe.getTests)
	engine.GET("/v1/tools", wardrobe.getTools)

	return engine
}

func controllerNetworkHandler(context *ATQContext) *gin.Engine {
	engine := gin.Default()
	engine.MaxMultipartMemory = 8 << 20

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.POST("/v1/container", context.postRegisterNewContainer)
	engine.GET("/v1/containers/:group", context.GetContainersByGroup)
	engine.GET("/v1/container/:hostname", context.getContainer)

	return engine
}
