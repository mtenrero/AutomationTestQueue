package main

// TestType defines the user assigned id, testType name (should be descriptive)
// and the test tool path whose test will be launched with:
// For example:
// Apache Jmeter is the type and test.jmx the test. So Jmeter should point to the
// valid full jmeter.jar paath in the target system
type TestType struct {
	ID   string `json:"id"`
	name string `json:"name"`
	path string `json:"path"`
}

// Test defines the test file and the tool/TestType which will be launched with
type Test struct {
	TestType *TestType `json:"testType"`
	Name     string    `json:"name"`
}
