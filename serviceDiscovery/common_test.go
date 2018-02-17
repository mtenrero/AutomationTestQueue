package serviceDiscovery

import (
	"os"
	"testing"
)

const ENDPOINT_VAL = "http://localhost:8080"

func TestGetEnv(t *testing.T) {
	os.Setenv(ENDPOINT, ENDPOINT_VAL)

	addr, err := GetFlightControllerEnv()
	if err != nil {
		t.Error(err)
	}

	if addr == nil {
		t.Fail()
	}
}

func TestGetEnvNoEnvs(t *testing.T) {
	os.Unsetenv(ENDPOINT)

	_, err := GetFlightControllerEnv()
	if err == nil {
		t.Error(err)
	}
}

func TestGetVip(t *testing.T) {
	_, err := getVIP()

	if err != nil {
		t.Fail()
	}

}
