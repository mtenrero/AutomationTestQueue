package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mtenrero/AutomationTestQueue/configLoader"
	"github.com/stretchr/testify/assert"
)

const testVersion = "/v1"
const host = "localhost:8080"

const pingResponse = "{\"message\":\"pong\"}"
const testToolsResponse = "{\"tools\":[{\"Alias\":\"TEST_TOOL\",\"Name\":\"Unrunnable test tool\",\"Envs\":[\"env1\",\"env2\"]}]}"

func TestPing(t *testing.T) {
	tools := configLoader.LoadConfigYaml("tools.yml")
	gWardrobe := NewWardrobe(tools)

	engine := networkHandler(gWardrobe)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, pingResponse, recorder.Body.String())
}

func TestTool(t *testing.T) {
	tools := configLoader.LoadConfigYaml("tools_test.yml")
	gWardrobe := NewWardrobe(tools)

	engine := networkHandler(gWardrobe)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", testVersion+"/tools", nil)
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, testToolsResponse, recorder.Body.String())
}

func TestTestsNotLoaded(t *testing.T) {
	tools := configLoader.LoadConfigYaml("tools_test.yml")
	gWardrobe := NewWardrobe(tools)

	engine := networkHandler(gWardrobe)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", testVersion+"/test", nil)
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 204, recorder.Code)
}
