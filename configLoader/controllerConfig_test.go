package configLoader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadControllerConfigYaml(t *testing.T) {
	config, err := LoadControllerConfigYaml("../controller-config.yaml")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "1050", config.Port, "Port reading from YAML failed")
	assert.Equal(t, "mtenrero/jmeter", config.Images[0].Name, "Read Containers from YAML Failed!")
}

func TestLoadControllerConfigYamlFileErr(t *testing.T) {
	_, err := LoadControllerConfigYaml("../controller-ERROR.yaml")

	if err == nil {
		t.Log(err)
	}
}
