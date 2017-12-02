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
