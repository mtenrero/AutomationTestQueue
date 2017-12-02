package main

import "github.com/gin-gonic/gin"

var engine *gin.Engine

// Gin Router, listen defined endpoints
func networkHandler(wardrobe *Wardrobe) {
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
	engine.Run() // listen and serve on 0.0.0.0:8080 by default

}
