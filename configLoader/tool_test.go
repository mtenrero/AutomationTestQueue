package configLoader

import "testing"

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
