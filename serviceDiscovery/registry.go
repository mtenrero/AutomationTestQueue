package serviceDiscovery

// Registry is one entry in the Service-Discovery Registry
import (
	"fmt"
	"net"
	"strings"
)

// RegistryEntry is an entry of the Registry collection which contains the state of a registered container
/**
 * @apiDefine RegistryEntry
 * @apiSuccess {int} 		uuid 			unique Identifier of the container in the Controller.
 * @apiSuccess {string} 	virtualIP 		Virtual IP of the container
 * @apiSuccess {status} 	status 			Current status of the container
 * @apiSuccess {string} 	hostname		Container's hostname
 * @apiSuccess {string} 	group 			Group which the container belongs to
 */
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
		if registryEntry.Hostname == hostname {
			return &registryEntry
		}
	}

	return nil
}

// RegistriesMembersOf returns a Collection of RegistryEntries with the containers included in a given group
func (regColl *RegistryCollection) RegistriesMembersOf(group string) *RegistryCollection {
	registry := NewRegistryCollection()

	for _, registryEntry := range *regColl {
		if registryEntry.Group == group {
			newRegistry, _ := registry.Add(&registryEntry)
			registry = newRegistry
		}
	}

	return registry
}

// IpsToCsv exports the contained IPs to CSV format
func (reg *RegistryCollection) IpsToCsv() string {
	var ips = make([]string, 0, 0)
	for _, registryEntry := range *reg {
		vip := *registryEntry.VirtualIP
		ips = append(ips, vip.String())
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ips)), ","), "[]")
}
