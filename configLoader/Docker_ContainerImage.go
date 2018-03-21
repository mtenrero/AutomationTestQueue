package configLoader

// Image defines the container name and its available variables
type Image struct {
	Name      string     `yaml:"name"`
	Variables []Variable `yaml:"variables"`
}

// Images defines a collection of Docker Containers
type Images struct {
	Images []Image `yaml:"images"`
}
