package serviceDiscovery

import (
	"bytes"
	"net"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// controllerPath defines the HTTP Path which will be called to register the container
var controllerPath = "/v1/container"

// Register a new container in the Registry
func Register() error {
	fcAddr, err := GetFlightControllerEnv()
	if err != nil {
		logger.WithFields(log.Fields{
			"event": "getFlightControllerEnv",
			"error": err,
		}).Error("Failed to get Controller endpoint environment variable")
		return err
	}

	ipAddr, err := getVIP()
	if err != nil {
		logger.WithFields(log.Fields{
			"event": "getVIP",
			"error": err,
		}).Error("Failed to get container/host address")
		return err
	}

	register(fcAddr, ipAddr)
	return nil
}

func register(fcAddr *url.URL, containerIP *net.IP) (*RegistryEntry, error) {
	form := url.Values{
		"containerIP": {containerIP.String()},
		"group":       {"ATQ"},
	}

	body := bytes.NewBufferString(form.Encode())

	_, err := http.Post(fcAddr.String()+controllerPath, "application/x-www-form-urlencoded", body)
	if err != nil {
		logger.WithFields(log.Fields{
			"event": "POSTregister",
			"error": err,
		}).Error("Failed to Register the container in the Controller")
		return nil, err
	}

	return nil, newErrMissingEnvs("Unable to contact with FlightController")
}
