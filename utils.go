package main

import "github.com/mtenrero/AutomationTestQueue/configLoader"

// LoadConfig return the initialized Config struct loaded with the tools
// specified in the YAML file specified in the given path
func LoadConfig(path string) *configLoader.Tools {
	var tools *configLoader.Tools

	tools = configLoader.LoadConfigYaml(path)

	return tools
}
