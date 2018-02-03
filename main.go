package main

import (
	"flag"
	"os"

	"github.com/mtenrero/AutomationTestQueue/serviceDiscovery"
	log "github.com/sirupsen/logrus"
)

var gWardrobe Wardrobe

func init() {
	log.SetOutput(os.Stdout)
}

func main() {

	var mode = flag.String("mode", "controller", "ATQ start mode: controller/registrator/brainmaster")
	flag.Parse()

	switch *mode {
	case "controller":
		log.WithField("mode", *mode).Info("Starting Controller Mode…")
		startController()
	case "registrator":
		log.WithField("mode", *mode).Info("Starting Registrator Mode…")
		startRegistrator()
	case "brainmaster":
		log.WithField("mode", *mode).Info("Starting BrainMaster Mode…")
		startBrainMaster()
	default:
		log.WithField("mode", *mode).Fatal("No valid ATQ execution mode detected!")
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
	serviceDiscovery.Register()
}

func startBrainMaster() {
	tools := LoadConfig("tools.yml")
	gWardrobe := NewWardrobe(tools)

	engine := networkHandler(gWardrobe)
	engine.Run() // listen and serve on 0.0.0.0:8080 by default
}
