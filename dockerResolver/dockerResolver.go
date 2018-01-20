package dockerResolver

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/lextoumbourou/goodhosts"
)

type ErrNoHostname struct {
	message string
}

type ErrNoIPV4Found struct {
	ipaddr  net.IP
	message string
}

func (e *ErrNoHostname) Error() string {
	return e.message
}

func (e *ErrNoIPV4Found) Error() string {
	return e.message
}

// newErrNoHostname initializes new Error when no hostname found
func newErrNoHostname(message string) *ErrNoHostname {
	return &ErrNoHostname{
		message: message,
	}
}

// newErrNoIPV4Found initializes new Error when parsing IPV4 address
func newErrNoIPV4Found(message string, ip net.IP) *ErrNoIPV4Found {
	return &ErrNoIPV4Found{
		message: message,
		ipaddr:  ip,
	}
}

// GetHostname returns the container or host system hostname
func GetHostname() (string, error) {
	name, err := os.Hostname()
	if err == nil {
		return name, nil
	}
	return "", newErrNoHostname("This host or container doesn't have a hostname configured")
}

// GetVIP4 returns the Virtual Ip (IPV4) if it's present in hosts file
func GetVIP4(hostname string) (*net.IP, error) {

	hosts, err := goodhosts.NewHosts()
	if err != nil {
		return nil, err
	}

	for _, line := range hosts.Lines {
		for _, host := range line.Hosts {
			if strings.Contains(host, hostname) {
				ipaddr := net.ParseIP(line.IP)

				if ipaddr.To4() != nil {
					fmt.Println(line.IP)
					return &ipaddr, nil
				}
			}
		}
	}

	return nil, newErrNoIPV4Found("The hostname IPV4 couldn't be found because it's not present in hosts file", nil)
}
