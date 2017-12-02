package configLoader

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tool struct represents the tools available on the target system.
// All paths must be valid
type Tool struct {
	Alias string   `yaml:"alias"`
	Name  string   `yaml:"name"`
	path  string   `yaml:"path"`
	Envs  []string `yaml:"envs"`
}

// Tools represent the container of all Tool
type Tools struct {
	Tools []Tool `yaml:"tools"`
}

// Find tries to find for a Tool in the loaded Tools by alias, if the Tool
// can't be found, then returns nil
func (tools *Tools) Find(alias string) *Tool {
	for _, tool := range tools.Tools {
		if tool.Alias == alias {
			log.Printf("TOOL FOUND: %s.\n", alias)
			return &tool
		}
	}
	log.Printf("TOOL NOT FOUND: %s.\n", alias)
	return nil
}

// GetPath exports the path variable to ve visible during development but not in yaml exports
func (tool *Tool) GetPath() string {
	return tool.path
}

// CheckEnvs ensures that all the required envs are available in the given Request Context
func (tool *Tool) CheckEnvs(context *gin.Context) bool {
	for _, envName := range tool.Envs {
		if context.PostForm(envName) == "" {
			context.JSON(http.StatusBadRequest, gin.H{"message": "The request not specifies all the variables required to run the test", "envs": tool.Envs})
			return false
		}
	}
	return true
}
