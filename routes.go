package main

import (
	"fmt"
	"net/http"

	"github.com/mtenrero/AutomationTestQueue/configLoader"

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
		return
	}

	test := configLoader.Test{tool, file.Filename, nil}

	wardrobe = wardrobe.AddTest(&test)

	context.JSON(http.StatusOK, gin.H{"availableTests": len(*wardrobe.tests), "tests": *wardrobe.tests, "addedTest": nameParam})
}

func parseEnvs(tool *configLoader.Tool, context *gin.Context) map[string]string {
	requirementsOK := tool.CheckEnvs(context)
	mapEnvs := make(map[string]string)

	if requirementsOK {
		for _, envName := range tool.Envs {
			mapEnvs[envName] = context.PostForm(envName)
		}
	}
	return mapEnvs
}

func (wardrobe *Wardrobe) runTest(context *gin.Context) {
	test, err := wardrobe.GetTest(context.PostForm("name"))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Test Not Found", "output": nil})
	} else {
		fmt.Println(test.Name)

		tool := test.Tool
		requirementsOK := tool.CheckEnvs(context)

		//out := toolRunner.ExecTest(test)

		if requirementsOK {
			context.JSON(http.StatusOK, gin.H{"output": "out"})
		}
	}

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
