package configLoader

// Variable defines the variable and its requirement need
type Variable struct {
	Name     string `yaml:"name"`
	Required bool   `yaml:"required"`
}
