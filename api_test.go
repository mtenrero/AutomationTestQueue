package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/mtenrero/AutomationTestQueue/configLoader"
	"github.com/stretchr/testify/assert"
)

const testVersion = "/v1"
const host = "localhost:8080"

const pingResponse = "{\"message\":\"pong\"}"
const testToolsResponse = "{\"tools\":[{\"Alias\":\"TEST_TOOL\",\"Name\":\"Unrunnable test tool\",\"Envs\":[\"env1\",\"env2\"]}]}"

const toolAlias = "TEST_TOOL"
const testName = "test"

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

func TestEnvParsing(t *testing.T) {
	tools := configLoader.LoadConfigYaml("tools_test.yml")
	gWardrobe := NewWardrobe(tools)

	engine := networkHandler(gWardrobe)

	form := url.Values{}

	form.Add("name", testName)
	form.Add("toolAlias", toolAlias)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", testVersion+"/test", nil)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 404, recorder.Code)
}

func TestUploadTestInvalidToolAlias(t *testing.T) {
	tools := configLoader.LoadConfigYaml("tools_test.yml")
	gWardrobe := NewWardrobe(tools)

	engine := networkHandler(gWardrobe)

	form := url.Values{}

	form.Add("name", testName)
	form.Add("toolAlias", "INVALID")

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", testVersion+"/uploadTest", nil)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 416, recorder.Code)
}

func TestUploadTestInvalidFile(t *testing.T) {
	tools := configLoader.LoadConfigYaml("tools_test.yml")
	gWardrobe := NewWardrobe(tools)

	engine := networkHandler(gWardrobe)

	form := url.Values{}

	form.Add("name", testName)
	form.Add("toolAlias", toolAlias)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", testVersion+"/uploadTest", nil)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
}

func TestGetEnvsFromRequest(t *testing.T) {
	form := url.Values{}

	form.Add("name", testName)
	form.Add("toolAlias", toolAlias)

	req, _ := http.NewRequest("POST", testVersion+"/test", nil)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	values := req.URL.Query

	fmt.Println(values())
}
