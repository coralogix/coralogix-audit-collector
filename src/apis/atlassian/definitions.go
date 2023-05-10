package atlassian

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	apiEndpoint = "%s/admin/v1/orgs/%s/events?startDate=%d&endDate=%d"
)

var (
	orgId    = os.Getenv("ATLASSIAN_ORG_ID")
	clientId = os.Getenv("ATLASSIAN_CLIENT_ID")
	apiToken = os.Getenv("ATLASSIAN_API_TOKEN")
	baseUrl  = os.Getenv("BASE_URL")
)

func validateEnvVars() {
	if orgId == "" {
		logrus.Fatalf("ATLASSIAN_ORG_ID env var is not set")
	}

	if clientId == "" {
		logrus.Fatalf("ATLASSIAN_CLIENT_ID env var is not set")
	}

	if apiToken == "" {
		logrus.Fatalf("ATLASSIAN_API_TOKEN env var is not set")
	}

	if baseUrl == "" {
		logrus.Fatalf("BASE_URL env var is not set")
	}
}
