package configLoader

// Test defines the test file and the tool/TestType which will be launched with
type Test struct {
	Tool *Tool             `json:"tool"`
	Name string            `json:"name"`
	Envs map[string]string `json:"envs"`
}
