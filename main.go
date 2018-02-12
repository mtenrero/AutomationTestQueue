package main

import (
	"flag"
	"os"

	"github.com/mtenrero/AutomationTestQueue/serviceDiscovery"
	"github.com/sirupsen/logrus"
)

var gWardrobe Wardrobe

func init() {
	logrus.SetOutput(os.Stdout)
}

func main() {
	logger := logrus.New()

	var mode = flag.String("mode", "controller", "ATQ start mode: controller/registrator/brainmaster")
	flag.Parse()

	switch *mode {
	case "controller":
		logger.WithField("mode", *mode).Info("Starting Controller Mode…")
		startController(logger)
	case "registrator":
		logger.WithField("mode", *mode).Info("Starting Registrator Mode…")
		startRegistrator()
	case "brainmaster":
		logger.WithField("mode", *mode).Info("Starting BrainMaster Mode…")
		startBrainMaster()
	default:
		logger.WithField("mode", *mode).Fatal("No valid ATQ execution mode detected!")
		os.Exit(-1)
	}
}

func startController(logger *logrus.Logger) {

	registry := serviceDiscovery.NewRegistryCollection()

	atqContext := ATQContext{registry: registry}

	engine := controllerNetworkHandler(&atqContext, logger)
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
