package jfrog

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	username = os.Getenv("JFROG_USERNAME")
	apiToken = os.Getenv("JFROG_API_TOKEN")
	baseUrl  = os.Getenv("BASE_URL")
)

func validateEnvVars() {
	if username == "" {
		logrus.Fatal("JFROG_USERNAME env var is not set")
	}
	if apiToken == "" {
		logrus.Fatal("JFROG_API_TOKEN env var is not set")
	}
	if baseUrl == "" {
		logrus.Fatal("BASE_URL env var is not set")
	}
}
