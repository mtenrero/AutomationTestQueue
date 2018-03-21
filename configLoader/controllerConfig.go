package configLoader

// ControllerConfig contains all the configuration available for Controller Mode
type ControllerConfig struct {
	Port   string  `yaml:"port"`
	Images []Image `yaml:"images"`
}
