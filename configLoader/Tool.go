package configLoader

import (
	"log"
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
