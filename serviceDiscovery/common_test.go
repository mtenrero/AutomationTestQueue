package serviceDiscovery

import (
	"os"
	"testing"
)

const EndpointVal = "http://localhost:8080"

func TestGetEnv(t *testing.T) {
	os.Setenv(EndpointEnv, EndpointVal)

	addr, err := GetFlightControllerEnv()
	if err != nil {
		t.Error(err)
	}

	if addr == nil {
		t.Fail()
	}
}

func TestGetEnvNoEnvs(t *testing.T) {
	os.Unsetenv(EndpointEnv)

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

func TestGetGroupEnv(t *testing.T) {
	os.Setenv(GroupEnv, EndpointVal)

	group, err := GetControllerGroupEnv()
	if err != nil {
		t.Error(err)
	}

	if group != EndpointVal {
		t.Fail()
	}
}

func TestGetGroupEnvNoEnv(t *testing.T) {
	os.Unsetenv(GroupEnv)

	_, err := GetControllerGroupEnv()
	if err == nil {
		t.Error(err)
	}
}
