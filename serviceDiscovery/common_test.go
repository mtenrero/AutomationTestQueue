package serviceDiscovery

import (
	"os"
	"strconv"
	"testing"
)

const HOSTNAME_VAL = "HOSTNAME_VAL"
const PORT_VAL = 1112

func TestGetEnv(t *testing.T) {
	os.Setenv(HOSTNAME, HOSTNAME_VAL)
	os.Setenv(PORT, strconv.Itoa(PORT_VAL))

	addr, err := GetFlightControllerEnv()
	if err != nil {
		t.Error(err)
	}

	if addr == nil {
		t.Fail()
	}

	if addr.Hostname != HOSTNAME_VAL || addr.Port != PORT_VAL {
		t.Fail()
	}
}

func TestGetEnvPortString(t *testing.T) {
	os.Setenv(HOSTNAME, HOSTNAME_VAL)
	os.Setenv(PORT, "STRINGPORT")

	_, err := GetFlightControllerEnv()
	if err == nil {
		t.Error(err)
	}
}

func TestGetEnvNoEnvs(t *testing.T) {
	os.Unsetenv(HOSTNAME)
	os.Unsetenv(PORT)

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
