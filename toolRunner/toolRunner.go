package toolRunner

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/mtenrero/AutomationTestQueue/configLoader"
)

func ExecTest(test *configLoader.Test) []byte {
	var tool *configLoader.Tool
	tool = test.Tool

	fmt.Println()

	out, err := exec.Command(tool.GetPath()).Output()
	if err != nil {
		log.Fatal(err)
		panic("some error found")
	}

	return out
}
