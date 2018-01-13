package configLoader

import "testing"
import "github.com/stretchr/testify/assert"

const ToolAlias = "ALIAS"
const ToolName = "NAME"
const ToolPath = "PATH"

var ToolEnvs = []string{"0", "1"}

func TestFindTool(t *testing.T) {
	tools := LoadConfigYaml("../tools_test.yml")

	tool := tools.Find("TEST_TOOL")

	if tool == nil {
		t.Error("The tool could not be found")
	}
}

func TestGetPath(t *testing.T) {
	tools := LoadConfigYaml("../tools_test.yml")

	tool := tools.Find("TEST_TOOL")

	if tool.GetPath() != tool.path {
		t.Error("GetPath function doesn't work properly")
	}
}

func TestToolNotFound(t *testing.T) {
	tools := LoadConfigYaml("../tools_test.yml")

	tool := tools.Find("NOT_DEFINED_TOOL")

	if tool != nil {
		t.Error("The tool shouldn't be found")
	}
}

func TestMakeTool(t *testing.T) {
	var tool Tool

	tool = *tool.Make(ToolAlias, ToolName, ToolPath, ToolEnvs)

	assert.Equal(t, ToolAlias, tool.Alias, "ERROR Making New Tool")
	assert.Equal(t, ToolName, tool.Name, "ERROR Making New Tool")
	assert.Equal(t, ToolPath, tool.path, "ERROR Making New Tool")
	assert.Equal(t, ToolEnvs, tool.Envs, "ERROR Making New Tool")
}
