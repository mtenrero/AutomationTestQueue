package serviceDiscovery

// Status represents the available statuses for a registered containers
type Status int

// Available statuses for a registered containers
const (
	WaitingTest Status = 1 + iota
	Testing
	Finished
	Dead
)
