package main

import (
	"errors"
	"log"

	"github.com/mtenrero/AutomationTestQueue/configLoader"
)

// Wardrobe is a collection of available tests & testLaunchers available.
// Tests abvailable can be modified with the appropiate REST calls.
// Check API documentation if you need more information
type Wardrobe struct {
	tests *[]configLoader.Test
	tools *configLoader.Tools
}

// NewWardrobe return a new initialized Wardrobe
func NewWardrobe(tools *configLoader.Tools) *Wardrobe {
	return &Wardrobe{&[]configLoader.Test{}, tools}
}

// AddTest adds a Test into the Wardrobe
func (wardrobe *Wardrobe) AddTest(test *configLoader.Test) *Wardrobe {
	eWardrobe := append(*wardrobe.tests, *test)
	wardrobe.tests = &eWardrobe
	return wardrobe
}

// GetTest look up for a test in the wardrobe given its name
func (wardrobe *Wardrobe) GetTest(name string) (*configLoader.Test, error) {
	for _, test := range *wardrobe.tests {
		if test.Name == name {
			log.Printf("TEST FOUND: %s.\n", name)
			return &test, nil
		}
	}
	log.Printf("TEST NOT FOUND: %s.\n", name)
	return nil, errors.New("Test " + name + " not found in the Wardrobe")
}
