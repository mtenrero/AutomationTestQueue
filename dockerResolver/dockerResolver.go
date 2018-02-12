package dockerResolver

import (
	"net"
	"os"
	"strings"

	"github.com/lextoumbourou/goodhosts"
	"github.com/sirupsen/logrus"
)

var logger = logrus.WithFields(logrus.Fields{
	"action": "dockerresolver",
	"mode":   "registrator",
})

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
		logger.WithField("hostname", name).Info("Hostname acquired")
		return name, nil
	}
	return "", newErrNoHostname("This host or container doesn't have a hostname configured")
}

// GetVIP4 returns the Virtual Ip (IPV4) if it's present in hosts file
func GetVIP4(hostname string) (*net.IP, error) {

	hosts, err := goodhosts.NewHosts()
	if err != nil {
		logger.WithField("goodhosts", err.Error()).Error("GoodHosts library failed!")
		return nil, err
	}

	for _, line := range hosts.Lines {
		for _, host := range line.Hosts {
			if strings.Contains(host, hostname) {
				logger.WithField("goodhosts", host).Info("Hostname found in hosts file")
				ipaddr := net.ParseIP(line.IP)

				if ipaddr.To4() != nil {
					return &ipaddr, nil
				}
			}
		}
	}

	return nil, newErrNoIPV4Found("The hostname IPV4 couldn't be found because it's not present in hosts file", nil)
}
