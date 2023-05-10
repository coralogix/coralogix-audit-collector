package zoom

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	apiEndpoint = "%s/v2/report/operationlogs?from=%s&to=%s"
)

var (
	accountId      = os.Getenv("ZOOM_ACCOUNT_ID")
	clientId       = os.Getenv("ZOOM_CLIENT_ID")
	clientSecret   = os.Getenv("ZOOM_CLIENT_SECRET")
	baseUrl        = os.Getenv("BASE_URL")
	accessTokenUrl = os.Getenv("ACCESS_TOKEN_URL")
)

func validateEnvVars() {
	if accountId == "" {
		logrus.Fatal("ZOOM_ACCOUNT_ID is not set")
	}

	if clientId == "" {
		logrus.Fatal("ZOOM_CLIENT_ID is not set")
	}

	if clientSecret == "" {
		logrus.Fatal("ZOOM_CLIENT_SECRET is not set")
	}

	if baseUrl == "" {
		logrus.Fatal("BASE_URL is not set")
	}

	if accessTokenUrl == "" {
		logrus.Fatal("ACCESS_TOKEN_URL is not set")
	}
}
