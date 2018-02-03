package serviceDiscovery

import (
	"bytes"
	"net"
	"net/http"
	"net/url"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Register a new container in the Registry
func Register() error {
	fcAddr, err := GetFlightControllerEnv()
	if err == nil {
		logger.WithFields(log.Fields{
			"event": "getFlightControllerEnv",
			"key":   err,
		}).Error("Failed to get Controller endpoint environment variable")
		return err
	}

	ipAddr, err := getVIP()
	if err == nil {
		logger.WithFields(log.Fields{
			"event": "getVIP",
			"key":   err,
		}).Error("Failed to get container/host address")
		return err
	}

	fullAddr := fcAddr.Hostname + ":" + strconv.Itoa(fcAddr.Port)

	register(fullAddr, ipAddr)
	return nil
}

func register(fcAddr string, containerIP *net.IP) (*RegistryEntry, error) {
	form := url.Values{
		"containerIP": {containerIP.String()},
		"group":       {"ATQ"},
	}

	body := bytes.NewBufferString(form.Encode())

	_, err := http.Post(fcAddr, "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}

	return nil, newErrMissingEnvs("Unable to contact with FlightController")
}
