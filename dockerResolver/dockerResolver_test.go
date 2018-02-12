package dockerResolver

import (
	"fmt"
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
	ipaddr, err := GetVIP4("MacBook-Pro-de-Marcos.local")
	if err != nil {
		t.Error(err)
	}
	t.Log(ipaddr.String())
	fmt.Println(ipaddr.String())

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
