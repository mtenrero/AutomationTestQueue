package dockerMiddleware

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.WithFields(logrus.Fields{
	"action": "dockerAPI",
	"mode":   "controller",
})
