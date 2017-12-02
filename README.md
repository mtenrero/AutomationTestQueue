# Automation Test Queue (ATQ) 
[![Build Status](https://semaphoreci.com/api/v1/mtenrero/automationtestqueue/branches/master/shields_badge.svg)](https://semaphoreci.com/mtenrero/automationtestqueue)

ATQ is an HTTP API designed to launch tests over remote machines. 
It's designed to be integrated in containers or pods with other services.

ATQ exposes an API you can call and give orders to the services you had configured previously,
like upload test files to the container/machine and run them with another API call with the
predefined tool in a config file. 

This simplifies the file handling in already started containers giving an abstraction layer to the
developer avoiding to handle TTY/SCP/SSH/Socket connections to the container using the Docker interface.

**** THIS IS AN INITIAL VERSION AND IT'S NOT COMPLETE !! ****

Guarantee it's not provided, use at your own risk.