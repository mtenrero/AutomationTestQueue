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

	group, err := GetControllerGroupEnv()
	if err != nil {
		logger.WithFields(log.Fields{
			"event": "getControllerGroupEnv",
			"error": err,
		}).Error("Failed to get container group environment variable")
		return err
	}

	register(fcAddr, ipAddr, group)
	return nil
}

func register(fcAddr *url.URL, containerIP *net.IP, group string) (*RegistryEntry, error) {

	registerLogger := logger.WithFields(log.Fields{
		"event": "POSTregister"})

	form := url.Values{
		"containerIP": {containerIP.String()},
		"group":       {group},
	}

	body := bytes.NewBufferString(form.Encode())

	resp, err := http.Post(fcAddr.String()+controllerPath, "application/x-www-form-urlencoded", body)
	if err != nil {
		registerLogger.WithFields(log.Fields{
			"error": err,
		}).Error("Failed to Register the container in the Controller")
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
		registerLogger.WithFields(log.Fields{
			"statusCode": resp.StatusCode}).Info("The container was succesfully registered")
		break
	case 409:
		registerLogger.WithFields(log.Fields{
			"statusCode": resp.StatusCode}).Warn("The container it's already registered")
		break
	default:
		registerLogger.WithFields(log.Fields{
			"statusCode": resp.StatusCode}).Error("Unexptected HTTP Response registering the container")
	}

	return nil, newErrMissingEnvs("Unable to contact with FlightController")
}
