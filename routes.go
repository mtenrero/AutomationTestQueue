package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (wardrobe *Wardrobe) uploadTest(context *gin.Context) {
	nameParam := context.PostForm("name")
	testTypeParam := context.PostForm("testType")

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

	testType := TestType{testTypeParam, "Apache Jmeter", "jmeterEntrypoint.sh"}

	wardrobe = wardrobe.AddTest(&Test{&testType, file.Filename})

	context.JSON(http.StatusOK, gin.H{"availableTests": len(*wardrobe.tests), "tests": *wardrobe.tests, "addedTest": nameParam})
}

func runTest(context *gin.Context) {

}

func (wardrobe *Wardrobe) getTests(context *gin.Context) {
	if len(*wardrobe.tests) == 0 {
		context.JSON(http.StatusNoContent, gin.H{"testNumber": len(*wardrobe.tests), "tests": *wardrobe.tests})
	} else {
		context.JSON(http.StatusOK, gin.H{"testNumber": len(*wardrobe.tests), "tests": *wardrobe.tests})
	}
}
