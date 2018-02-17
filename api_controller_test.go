package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/mtenrero/AutomationTestQueue/serviceDiscovery"
	"github.com/stretchr/testify/assert"
)

const apiVersion = "/v1"

func TestControllerPing(t *testing.T) {
	registry := serviceDiscovery.NewRegistryCollection()

	atqContext := ATQContext{registry: registry}

	engine := controllerNetworkHandler(&atqContext, logrus.New())

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestRegistration(t *testing.T) {
	registry := serviceDiscovery.NewRegistryCollection()

	atqContext := ATQContext{registry: registry}

	engine := controllerNetworkHandler(&atqContext, logrus.New())

	form := url.Values{}

	form.Add("containerIP", "10.0.0.1")
	form.Add("hostname", "HOSTNAME")
	form.Add("group", "GROUP")

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", apiVersion+"/container", nil)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestCSV(t *testing.T) {
	registry := serviceDiscovery.NewRegistryCollection()

	atqContext := ATQContext{registry: registry}

	engine := controllerNetworkHandler(&atqContext, logrus.New())

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", apiVersion+"/containers/GROUP/csv", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}
