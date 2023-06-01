package slack

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	accessToken = os.Getenv("SLACK_ACCESS_TOKEN")
	baseUrl     = os.Getenv("BASE_URL")
)

func validateEnvVars() {
	if accessToken == "" {
		logrus.Fatal("SLACK_ACCESS_TOKEN is not set")
	}

	if baseUrl == "" {
		logrus.Fatal("BASE_URL is not set")
	}
}
