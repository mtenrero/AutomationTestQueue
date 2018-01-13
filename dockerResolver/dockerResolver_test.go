package dockerResolver

import (
	"testing"
)

func TestHostname(t *testing.T) {
	hostname, err := GetHostname()
	if err != nil {
		t.Error(err)
	}

	t.Log(hostname)

	if len(hostname) == 0 {
		t.Fail()
	}
}

func TestGetVIP(t *testing.T) {
	ipaddr, err := GetVIP4("localhost")
	if err != nil {
		t.Error(err)
	}

	if len(ipaddr.String()) == 0 {
		t.Fail()
	}
}

func TestGetVIPBadHostname(t *testing.T) {
	_, err := GetVIP4("BADHOST")
	if err == nil {
		t.Fail()
	}
}
