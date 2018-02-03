package serviceDiscovery

import (
	"net"
	"os"
	"strconv"

	"github.com/mtenrero/AutomationTestQueue/dockerResolver"
	"github.com/mtenrero/AutomationTestQueue/network"
	log "github.com/sirupsen/logrus"
)

const HOSTNAME = "FLIGHTCONTROLLER_HOSTNAME"
const PORT = "FLIGHTCONTROLLER_PORT"

var logger = log.WithFields(log.Fields{
	"action": "register",
	"mode":   "registrator",
})

// ErrMissingEnvs struct
type ErrMissingEnvs struct {
	message string
}

// newErrNoIPV4Found initializes new Error when parsing IPV4 address
func newErrMissingEnvs(message string) *ErrMissingEnvs {
	return &ErrMissingEnvs{
		message: message,
	}
}

func (e *ErrMissingEnvs) Error() string {
	return e.message
}

// GetFlightControllerEnv returns the Address of the FlightController if defined
func GetFlightControllerEnv() (*network.Address, error) {
	fcHostname, fcHostnameValid := os.LookupEnv(HOSTNAME)
	fcPortStr, fcPortValid := os.LookupEnv(PORT)

	if !fcHostnameValid || !fcPortValid {
		logger.WithField(HOSTNAME, fcHostname).Warn("The Controller Hostname Env has not been declared")
		return nil, newErrMissingEnvs("The environment variables should be already declared")
	}

	fcPort, error := strconv.Atoi(fcPortStr)
	if error != nil {
		logger.WithField(PORT, fcPort).Warn("The Controller Port Env has not been declared")
		return nil, newErrMissingEnvs("The specified PORT is not a number!")
	}

	return &network.Address{Hostname: fcHostname, Port: fcPort}, nil
}

func getVIP() (*net.IP, error) {
	hostname, err := dockerResolver.GetHostname()
	if err != nil {
		logger.WithField("hostname", "dsa").Warn("ERROR")
		return nil, err
	}

	ipAddr, err := dockerResolver.GetVIP4(hostname)
	if err == nil {
		return nil, err
	}

	return ipAddr, nil
}
