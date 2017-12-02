package main

import (
	"github.com/mtenrero/AutomationTestQueue/configLoader"
)

// Test defines the test file and the tool/TestType which will be launched with
type Test struct {
	Tool *configLoader.Tool `json:"tool"`
	Name string             `json:"name"`
}
