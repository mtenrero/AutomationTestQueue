package serviceDiscovery

import (
	"os"
	"strconv"

	"github.com/mtenrero/AutomationTestQueue/network"
)

const HOSTNAME = "FLIGHTCONTROLLER_HOSTNAME"
const PORT = "FLIGHTCONTROLLER_PORT"

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
		return nil, newErrMissingEnvs("The environment variables should be already declared")
	}

	fcPort, error := strconv.Atoi(fcPortStr)
	if error != nil {
		return nil, newErrMissingEnvs("The specified PORT is not a number!")
	}

	return &network.Address{Hostname: fcHostname, Port: fcPort}, nil
}
