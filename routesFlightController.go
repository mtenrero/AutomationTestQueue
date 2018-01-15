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

// PostRegisterNewContainer registers a new container in the controller
func (registryColl *ATQContext) PostRegisterNewContainer(context *gin.Context) {
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

// GetContainer returns the container status if registered
func (registryColl *ATQContext) GetContainer(context *gin.Context) {
	hostname := context.PostForm("hostname")

	if len(hostname) == 0 {
		context.JSON(http.StatusExpectationFailed, gin.H{"message": "The required parameters are not provided! (hostname)"})
	} else {
		registryEntry := registryColl.registry.IsRegistered(hostname)
		if registryEntry == nil {
			context.JSON(http.StatusNotFound, gin.H{"message": "The given container hostname is not registered yet"})
		} else {
			context.JSON(http.StatusNotFound, gin.H{"message": "The given container hostname is registered", "registryEntry": registryEntry})
		}
	}
}

// GetContainersByGroup returns a list of containers members of a given group
func (registryColl *ATQContext) GetContainersByGroup(context *gin.Context) {
	group := context.PostForm("group")

	if len(group) == 0 {
		context.JSON(http.StatusExpectationFailed, gin.H{"message": "The required parameters are not provided! (group)"})
	} else {
		registryCollGrouped := registryColl.registry.RegistriesMembersOf(group)
		context.JSON(http.StatusNotFound, gin.H{"message": "These containers are members of the group " + group, "registryCollection": registryCollGrouped})
	}
}
