package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mtenrero/AutomationTestQueue/serviceDiscovery"
)

type ATQContext struct {
	registry *serviceDiscovery.RegistryCollection
}

/**
 * @api {post} /container?containerIP=:containerIP&hostname=:hostname&group=:group Register new Container in the Controller
 * @apiName ATQ REST API
 * @apiGroup FlightController
 *
 * @apiUse RegistryEntry
 *
 * @apiParam (containerIP) {String} id Virtual IP of the calling container.
 * @apiParam (hostname) {String} String ID or hostname of the calling container
 * @apiParam (group) {String} group name of the cluster which will be added to
 *
 * @apiSuccess {String} message Response user-friendly message.
 * @apiSuccess {registryEntry} registryEntry Registered Container data.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *	{
 *	"message": "Registration was succesfull",
 *	"registryEntry": {
 * 		"uuid": 1516451579,
 *		"virtualIP": "10.0.0.22",
 * 		"status": 0,
 * 		"hostname": "hostname",
 *		"group": "group"
 * 		}
 *	}
 *
 * @apiError (409 Conflict) ContainerAlreadyRegistered The container is already registered
 * @apiErrorExample {json} Error-Response:
 *     HTTP/1.1 409 Conflict
 *     {
	"message": "This container is already registered!",
	"registryEntry": {
		"uuid": 1516451579,
		"virtualIP": "",
		"status": 0,
		"hostname": "",
		"group": ""
	}
	}
*/
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

// getContainer returns the container status if registered
/**
 * @api {get} /container/:hostname Get Container information
 * @apiGroup FlightController
 *
 * @apiUse RegistryEntry
 *
 * @apiParam (hostname) {String} String ID or hostname of the calling container
 *
 * @apiSuccess {String} message Response user-friendly message.
 * @apiSuccess {registryEntry} registryEntry Registered Container data.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *	{
 *	"message": "The given container hostname is registered",
 *	"registryEntry": {
 * 		"uuid": 1516451579,
 *		"virtualIP": "10.0.0.22",
 * 		"status": 0,
 * 		"hostname": "hostname",
 *		"group": "group"
 * 		}
 *	}
 *
 * @apiError (404 Not Found) ContainerNotRegistered The container is not registered
 * @apiErrorExample {json} Error-Response: ContainerNotRegistered
 *     HTTP/1.1 404 Not Found
 *     {
	"message": "The given container hostname is not registered yet",
	}
 *
 * @apiError (417 Expectation Failed) RequestWithoutParameters The Request hasn't required parameters
 * @apiErrorExample {json} Error-Response: RequestWithoutParameters
 *     HTTP/1.1 417 Expectation Failed
 *     {
	"message": "The required parameters are not provided! (hostname)",
	}
*/
func (registryColl *ATQContext) getContainer(context *gin.Context) {
	hostname := context.Param("hostname")

	fmt.Println("hostname: " + hostname)

	registryEntry := registryColl.registry.IsRegistered(hostname)
	if registryEntry == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "The given container hostname is not registered yet"})
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "The given container hostname is registered", "registryEntry": registryEntry})
	}

}

// GetContainersByGroup returns a list of containers members of a given group
/**
 * @api {get} /containers/:group Get Containers members of a group
 * @apiGroup FlightController
 *
 * @apiUse RegistryEntry
 *
 * @apiParam (group) {String} group name of the cluster
 *
 * @apiSuccess {String} message Response user-friendly message.
 * @apiSuccess {registryEntry} registryEntry Registered Container data.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *	{
 *	"message": "These containers are members of the group :group ",
 *	"registryCollection": [
	...
 	]
 *
 *	}
 *
 * @apiError (417 Expectation Failed) RequestWithoutParameters The Request hasn't required parameters
 * @apiErrorExample {json} Error-Response: RequestWithoutParameters
 *     HTTP/1.1 417 Expectation Failed
 *     {
	"message": "The required parameters are not provided! (group)",
	}
*/
func (registryColl *ATQContext) GetContainersByGroup(context *gin.Context) {
	group := context.Param("group")

	if len(group) == 0 {
		context.JSON(http.StatusExpectationFailed, gin.H{"message": "The required parameters are not provided! (group)"})
	} else {
		registryCollGrouped := registryColl.registry.RegistriesMembersOf(group)
		context.JSON(http.StatusOK, gin.H{"message": "These containers are members of the group " + group, "registryCollection": registryCollGrouped})
	}
}
