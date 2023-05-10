package confluence

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	apiEndpoint = "%s/wiki/rest/api/audit?startDate=%d&endDate=%d"
)

var (
	username = os.Getenv("CONFLUENCE_USERNAME")
	apiToken = os.Getenv("CONFLUENCE_API_TOKEN")
	baseUrl  = os.Getenv("BASE_URL")
)

func validateEnvVars() {
	if username == "" {
		logrus.Fatalf("CONFLUENCE_USERNAME env var is not set")
	}
	if apiToken == "" {
		logrus.Fatalf("CONFLUENCE_API_TOKEN env var is not set")
	}
	if baseUrl == "" {
		logrus.Fatalf("BASE_URL env var is not set")
	}
}
