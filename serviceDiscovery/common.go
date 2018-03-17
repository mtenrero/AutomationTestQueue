package serviceDiscovery

import (
	"net"
	"net/url"
	"os"

	"github.com/mtenrero/AutomationTestQueue/dockerResolver"
	log "github.com/sirupsen/logrus"
)

// EndpointEnv URL given by user by Environment Variable
const EndpointEnv = "ATQCONTROLLER_ENDPOINT"

// GroupEnv holds the Env Variable for the group which will be registered
const GroupEnv = "ATQCONTROLLER_GROUP"

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
func GetFlightControllerEnv() (*url.URL, error) {
	endpointStr, endpointValid := os.LookupEnv(EndpointEnv)

	if !endpointValid {
		logger.WithField(EndpointEnv, endpointStr).Warn("The Controller Hostname Env has not been declared")
		return nil, newErrMissingEnvs("The environment variables should be already declared")
	}

	url, err := url.Parse(endpointStr)
	if err != nil {
		logger.WithField("err", err).Error("The given endpoint isn't valid!")
	}

	return url, nil
}

// GetControllerGroup returns the group specified by the user via Environment Variables
func GetControllerGroupEnv() (string, error) {
	groupStr, groupValid := os.LookupEnv(GroupEnv)

	if !groupValid {
		logger.WithField(GroupEnv, groupStr).Warn("The Controller Hostname Env has not been declared")
		return "", newErrMissingEnvs("The environment variables should be already declared")
	}

	return groupStr, nil
}

func getVIP() (*net.IP, error) {
	hostname, err := dockerResolver.GetHostname()
	if err != nil {
		logger.WithField("hostname", "dsa").Warn("ERROR")
		return nil, err
	}

	ipAddr, err := dockerResolver.GetVIP4(hostname)
	if err != nil {
		return nil, err
	}

	return ipAddr, nil
}
