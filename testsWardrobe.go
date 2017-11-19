package main

// Wardrobe is a collection of available tests & testLaunchers available.
// Tests abvailable can be modified with the appropiate REST calls.
// Check API documentation if you need more information
type Wardrobe struct {
	tests     *[]Test
	testTypes *[]TestType
}

// NewWardrobe return a new initialized Wardrobe
func NewWardrobe() *Wardrobe {
	return &Wardrobe{&[]Test{}, &[]TestType{}}
}

// AddTest adds a Test into the Wardrobe
func (wardrobe *Wardrobe) AddTest(test *Test) *Wardrobe {
	eWardrobe := append(*wardrobe.tests, *test)
	wardrobe.tests = &eWardrobe
	return wardrobe
}
