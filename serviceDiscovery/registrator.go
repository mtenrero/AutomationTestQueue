package serviceDiscovery

import (
	"bytes"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/mtenrero/AutomationTestQueue/dockerResolver"
)

// Register a new container in the Registry
func Register() error {
	fcAddr, err := GetFlightControllerEnv()
	if err == nil {
		return err
	}

	ipAddr, err := getVIP()
	if err == nil {
		return err
	}

	fullAddr := fcAddr.Hostname + ":" + strconv.Itoa(fcAddr.Port)

	register(fullAddr, ipAddr)
	return nil

}

func ControllerAlive() bool {

	return true
}

func checkRegistration() bool {

	return true
}

func getVIP() (*net.IP, error) {
	hostname, err := dockerResolver.GetHostname()
	if err == nil {
		return nil, err
	}

	ipAddr, err := dockerResolver.GetVIP4(hostname)
	if err == nil {
		return nil, err
	}

	return ipAddr, nil
}

func register(fcAddr string, containerIP *net.IP) (*RegistryEntry, error) {
	form := url.Values{
		"containerIP": {containerIP.String()},
	}

	body := bytes.NewBufferString(form.Encode())

	_, err := http.Post(fcAddr, "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}

	return nil, newErrMissingEnvs("Unable to contact with FlightController")
}
