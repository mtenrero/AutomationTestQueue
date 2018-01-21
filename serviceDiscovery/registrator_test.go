package serviceDiscovery

import (
	"net"
	"testing"
)

func TestRegisterAuxiliary(t *testing.T) {
	fcAddress := "hhg"
	ipAddr := net.ParseIP("10.0.0.1")

	_, err := register(fcAddress, &ipAddr)

	if err == nil {
		t.Fail()
	}
}
