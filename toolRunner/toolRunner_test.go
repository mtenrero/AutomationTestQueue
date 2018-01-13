package toolRunner

import "testing"
import "github.com/mtenrero/AutomationTestQueue/configLoader"
import "fmt"

func TestBasicRun(t *testing.T) {
	var tool configLoader.Tool
	tool = *tool.Make("pwd", "pwd", "pwd", []string{})

	var test configLoader.Test
	test = configLoader.Test{Tool: &tool, Name: "UnitTest", Envs: nil}

	out := ExecTest(&test)

	fmt.Println(string(out))

	if len(string(out)) < 1 {
		t.Error("OS Exec failed")
	}
}
