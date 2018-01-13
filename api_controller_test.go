package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/mtenrero/AutomationTestQueue/serviceDiscovery"
	"github.com/stretchr/testify/assert"
)

const apiVersion = "/v1"

func TestControllerPing(t *testing.T) {
	registry := serviceDiscovery.NewRegistryCollection()

	atqContext := ATQContext{registry: registry}

	engine := controllerNetworkHandler(&atqContext)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestRegistration(t *testing.T) {
	registry := serviceDiscovery.NewRegistryCollection()

	atqContext := ATQContext{registry: registry}

	engine := controllerNetworkHandler(&atqContext)

	form := url.Values{}

	form.Add("containerIP", "10.0.0.1")
	form.Add("hostname", "HOSTNAME")
	form.Add("group", "GROUP")

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", apiVersion+"/register", nil)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}
