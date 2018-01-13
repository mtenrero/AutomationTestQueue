package serviceDiscovery

// Registry is one entry in the Service-Discovery Registry
import (
	"fmt"
	"net"
)

// RegistryEntry is an entry of the Registry collection which contains the state of a registered container
type RegistryEntry struct {
	UUID      int64   `json:"uuid"`
	VirtualIP *net.IP `json:"virtualIP"`
	Status    Status  `json:"status"`
	Hostname  string  `json:"hostname"`
	Group     string  `json:"group"`
}

// ErrRegistry is the error caused in the Registry operations
type ErrRegistry struct {
	message string
}

func (e *ErrRegistry) Error() string {
	return e.message
}

// newErrRegistry initializes new Registry Error
func newErrRegistry(message string) *ErrRegistry {
	return &ErrRegistry{
		message: message,
	}
}

// RegistryCollection is a Collection of RegistryEntry(ies)
type RegistryCollection []RegistryEntry

// Add a RegistryEntry into the RegistryCollection
func (reg *RegistryCollection) Add(regist *RegistryEntry) (*RegistryCollection, error) {

	if reg.IsRegistered(regist.Hostname) != nil {
		return nil, newErrRegistry("This container is already registered!")
	}

	newRegistry := append(*reg, *regist)
	return &newRegistry, nil
}

// MakeRegistryEntry initializes a new RegistryEntry
func MakeRegistryEntry(uuid int64, vip, hostname, group string, status Status) *RegistryEntry {
	vipNet := net.ParseIP(vip)
	return &RegistryEntry{
		UUID:      uuid,
		VirtualIP: &vipNet,
		Status:    status,
		Hostname:  hostname,
		Group:     group,
	}
}

// NewRegistryCollection initializes an empty registry collection
func NewRegistryCollection() *RegistryCollection {
	registry := make(RegistryCollection, 0)
	return &registry
}

// IsRegistered checks if a hostname is already registered in the Registry
func (regColl *RegistryCollection) IsRegistered(hostname string) *RegistryEntry {
	for _, registryEntry := range *regColl {
		fmt.Println(registryEntry.Hostname)
		if registryEntry.Hostname == hostname {
			return &registryEntry
		}
	}

	return nil
}
