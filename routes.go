package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (wardrobe *Wardrobe) uploadTest(context *gin.Context) {
	nameParam := context.PostForm("name")
	toolAlias := context.PostForm("toolAlias")

	tool := wardrobe.tools.Find(toolAlias)

	if tool == nil {
		context.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{"error": "The Tool specified is not available", "tools": *wardrobe.tools})
		return
	}

	file, err := context.FormFile("file")
	if err != nil {
		context.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	if err := context.SaveUploadedFile(file, file.Filename); err != nil {
		context.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		context.JSON(http.StatusServiceUnavailable, gin.H{"availableTests": len(*wardrobe.tests), "tests": *wardrobe.tests})
		return
	}

	test := Test{tool, file.Filename}

	wardrobe = wardrobe.AddTest(&test)

	context.JSON(http.StatusOK, gin.H{"availableTests": len(*wardrobe.tests), "tests": *wardrobe.tests, "addedTest": nameParam})
}

func (wardrobe *Wardrobe) runTest(context *gin.Context) {
	test := wardrobe.GetTest(context.PostForm("name"))
	fmt.Println(test.Name)
	context.JSON(http.StatusOK, gin.H{"test": test})
}

func (wardrobe *Wardrobe) getTests(context *gin.Context) {
	if len(*wardrobe.tests) == 0 {
		context.JSON(http.StatusNoContent, gin.H{"testNumber": len(*wardrobe.tests), "tests": *wardrobe.tests})
	} else {
		context.JSON(http.StatusOK, gin.H{"testNumber": len(*wardrobe.tests), "tests": *wardrobe.tests})
	}
}

func (wardrobe *Wardrobe) getTools(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"tools": wardrobe.tools.Tools})
}
