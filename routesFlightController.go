package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mtenrero/AutomationTestQueue/serviceDiscovery"
)

type ATQContext struct {
	registry *serviceDiscovery.RegistryCollection
}

func (registryColl *ATQContext) postRegisterNewContainer(context *gin.Context) {
	containerIP := context.PostForm("containerIP")
	hostname := context.PostForm("hostname")
	group := context.PostForm("group")

	uuid := time.Now().Unix()

	regEntry := serviceDiscovery.MakeRegistryEntry(uuid, containerIP, hostname, group, 0)

	newRegistryColl, err := registryColl.registry.Add(regEntry)
	if err != nil {
		log.Print(err.Error())
		context.JSON(http.StatusConflict, gin.H{"message": "This container is already registered!", "registryEntry": registryColl.registry.IsRegistered(hostname)})
	} else {
		registryColl.registry = newRegistryColl
		context.JSON(http.StatusOK, gin.H{"message": "Registration was succesfull", "registryEntry": *regEntry})
	}
}
