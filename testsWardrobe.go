package main

import (
	"github.com/mtenrero/AutomationTestQueue/configLoader"
)

// Wardrobe is a collection of available tests & testLaunchers available.
// Tests abvailable can be modified with the appropiate REST calls.
// Check API documentation if you need more information
type Wardrobe struct {
	tests *[]Test
	tools *configLoader.Tools
}

// NewWardrobe return a new initialized Wardrobe
func NewWardrobe(tools *configLoader.Tools) *Wardrobe {
	return &Wardrobe{&[]Test{}, tools}
}

// AddTest adds a Test into the Wardrobe
func (wardrobe *Wardrobe) AddTest(test *Test) *Wardrobe {
	eWardrobe := append(*wardrobe.tests, *test)
	wardrobe.tests = &eWardrobe
	return wardrobe
}
