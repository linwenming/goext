package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger_Logger(t *testing.T) {
	InitLogger("./logs/app.log")
	Error("logger:test log")
}

func TestLogger_Logrus(t *testing.T) {
	InitDefault("./logs/app.log")
	logrus.Info("logrus:test log")
}