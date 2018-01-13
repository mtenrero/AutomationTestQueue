package main

import (
	"flag"
	"log"
	"os"

	"github.com/mtenrero/AutomationTestQueue/serviceDiscovery"
)

var gWardrobe Wardrobe

func main() {

	var mode = flag.String("mode", "controller", "ATQ start mode: controller/registrator/brainmaster")
	flag.Parse()

	switch *mode {
	case "controller":
		startController()
	case "registrator":
		startRegistrator()
	case "brainmaster":
		startBrainMaster()
	default:
		log.Fatal("No valid ATQ execution mode detected! controller/registrator/brainmaster")
		os.Exit(-1)
	}
}

func startController() {
	registry := serviceDiscovery.NewRegistryCollection()

	atqContext := ATQContext{registry: registry}

	engine := controllerNetworkHandler(&atqContext)
	engine.Run()

}

func startRegistrator() {

}

func startBrainMaster() {
	tools := LoadConfig("tools.yml")
	gWardrobe := NewWardrobe(tools)

	engine := networkHandler(gWardrobe)
	engine.Run() // listen and serve on 0.0.0.0:8080 by default
}
