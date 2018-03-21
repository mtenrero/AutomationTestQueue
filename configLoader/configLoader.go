package configLoader

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// LoadConfigYaml reads the specified YAML config file and return it
// into the Config struct
func LoadConfigYaml(filePath string) *Tools {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	var tools Tools

	err = yaml.Unmarshal(yamlFile, &tools)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &tools
}

// LoadControllerConfigYaml reads the specified YAML config file for Controller Mode
func LoadControllerConfigYaml(filePath string) (*ControllerConfig, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var controllerConfig ControllerConfig

	err = yaml.Unmarshal(yamlFile, &controllerConfig)
	if err != nil {
		return nil, err
	}

	return &controllerConfig, nil
}
