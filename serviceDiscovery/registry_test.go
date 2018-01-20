package serviceDiscovery

import (
	"testing"
)

func TestCheckHostnameRegistered(t *testing.T) {
	registry := NewRegistryCollection()

	regist, _ := registry.Add(MakeRegistryEntry(222, "VIP", "HOSTNAME", "GROUP", 0))

	registered := regist.IsRegistered("HOSTNAME")

	if registered == nil {
		t.Fail()
	}
}

func TestCheckHostnameAlreadyRegistered(t *testing.T) {
	registry := NewRegistryCollection()

	regist, _ := registry.Add(MakeRegistryEntry(222, "VIP", "HOSTNAME", "GROUP", 0))

	_, err := regist.Add(MakeRegistryEntry(32, "VIP", "HOSTNAME", "GROUP2", 0))

	if err == nil {
		t.Fail()
	}

}

func TestRegistriesMemberOf(t *testing.T) {
	registry := NewRegistryCollection()
	regist, _ := registry.Add(MakeRegistryEntry(222, "VIP", "HOSTNAME", "GROUP", 0))
	regist2, _ := regist.Add(MakeRegistryEntry(3222, "VIP2", "HOSTNAME2", "GROUP2", 0))
	regist3, _ := regist2.Add(MakeRegistryEntry(23322, "VIP3", "HOSTNAME3", "GROUP", 0))

	registryMembers := regist3.RegistriesMembersOf("GROUP")

	if len(*registryMembers) != 2 {
		t.Fail()
	}
}
