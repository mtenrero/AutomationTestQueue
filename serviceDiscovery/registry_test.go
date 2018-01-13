package serviceDiscovery

import "testing"

func TestCheckHostnameRegistered(t *testing.T) {
	registry := NewRegistryCollection()

	regist, _ := registry.Add(MakeRegistryEntry(222, "VIP", "HOSTNAME", "GROUP", 0))

	registered := regist.IsRegistered("HOSTNAME")

	if registered == nil {
		t.Fail()
	}
}
