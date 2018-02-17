# Automation Test Queue (ATQ)

[![Build Status](https://api.travis-ci.org/mtenrero/AutomationTestQueue.svg)](https://travis-ci.org/mtenrero/AutomationTestQueue)
[![Coverage Status](https://coveralls.io/repos/github/mtenrero/AutomationTestQueue/badge.svg)](https://coveralls.io/github/mtenrero/AutomationTestQueue)
[![API Documentation](https://img.shields.io/badge/API-Documentation-orange.svg)](https://mtenrero.github.io/AutomationTestQueue/)

ATQ is an HTTP API designed to launch tests over remote machines.
It's designed to be integrated in containers or pods with other services.

ATQ exposes an API you can call and give orders to the services you had configured previously,
like upload test files to the container/machine and run them with another API call with the
predefined tool in a config file.

![Architecture](https://github.com/mtenrero/AutomationTestQueue/raw/master/readmeFiles/ATQ_arch.png)

This simplifies the file handling in already started containers giving an abstraction layer to the
developer avoiding to handle TTY/SCP/SSH/Socket connections to the container using the Docker interface.

## Launching Modes

ATQ can be launched in different modes:

### Controller

Controller is the mode which will be used to coordinate all the containers and its statuses. 

It exposes an API where the registrator component configured in the containers will call to in order to register all the instances.

```bash
./atq -mode=controller
```

### Registrator

Registrator is the component which should be added at the container image startup in order to register and communicate to the Controller.

```bash
./atq -mode=registrator
```

In order to work, this mode requires an environment variable **ATQCONTROLLER_ENDPOINT** specifying the endpoint and port where the Controller is running.

The environment variable can be global, local or specified at the program startup: 

```bash
env ATQCONTROLLER_ENDPOINT=http://localhost:8080 ./atq -mode=registrator
```

## Versioning

ATQ will not broke your code upgrading the Release version cause all major changes to the API will be developed in
a brand new version endpoint.

The initial version starts with `v1` tag: `https:\\host:port\v1\...`.
And the following versions will be `https:\\host:port\v2...`

## Configuration

ATQ need some previous configuration in order to work properly, like adding the desired tools in a config file. 

### **Tools**

The tools you want to make available to use thorugh the HTTP API must be specified in the file **tools.yml** and must be in the root path of the ATQ launcher. You can check up the example files available in this repository.

Available fields:

* **Alias**: Public alias exposed in the API. **Unique & Required**
* **Name**: Tool's detailed name. **Required**
* **Path**: **Full** path to the tool in the system. Relative paths may not work in certain systems. **Required**
    If your tool requires some args, you should put the name of the Env in uppercase characters before a **$**(USD) symbol.
    If the tool requires a **File** in the arguments, you can put the **$FILE** identifier.
* **Envs**: Environment variables required to launch the tool. **Optional**
    Environment variables can be converted to program arguments specifying the declared environment variable name in the path after a USD symbol

    NOTE: Check the tools.yml and tools_test.yml files for examples.

## API

### **GET** /v1/tools

This request doesn't need any parameters

* **200** Returns the available tools in the server in JSON format:

```json
{
    "tools": [
        {
            "Alias": "JMX_MOD",
            "Name": "Apache Jmeter Test",
            "Envs": [
                "file",
                "mode",
                "remotes"
            ]
        },
        {
            "Alias": "ECHO",
            "Name": "Echo Tool",
            "Envs": [
                "message"
            ]
        }
    ]
}
```

* **204** There isn't any tool available. You must configure the tool.yml before launch ATQ!.

### **GET** /v1/test

Returns the uploaded tests available to launch in the server. May not include any request parameter

* **204** There are not any test in the server yet.
* **200** Returns a JSON structure with the uploaded tests to the server.

### **POST** /v1/uploadTest (Multipart)

Send a file to the server.
**Requires at least two parameters:**

* **name**: Test Name
* **toolAlias**: Required tool wich will be launched with. Must be declared in the ***config.yml*** file.

**Optional:**

* **envs**: Environment variables required to launch the test.

### **POST** /v1/test

Run the test with the specified tool
**CURRENTLY IN DEVELOPMENT**

## Testing & Good practices

All commits should improve the test coverage and any commit that couldn't pass the tests will not be merged in the master branch.

Master branch is stable. All improvements or fixes should be done in the golang-dev branch or inside a specific branch.

Any feature should be well documented and tested. 

Godocs compatible code must be used.

## Disclaimer

**THIS IS AN INITIAL VERSION AND IT'S NOT COMPLETE !!**

_Guarantee it's not provided, use at your own risk._
