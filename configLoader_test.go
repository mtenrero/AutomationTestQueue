package main

import (
	"testing"

	"github.com/mtenrero/AutomationTestQueue/configLoader"
)

func TestLoadConfig(t *testing.T) {
	var tools *configLoader.Tools

	tools = configLoader.LoadConfigYaml("tools_test.yml")

	if tools.Find("TEST_TOOL").Alias != "TEST_TOOL" {
		t.Fail()
	}

	if tools.Find("TEST_TOOL").Name != "Unrunnable test tool" {
		t.Fail()
	}
}
