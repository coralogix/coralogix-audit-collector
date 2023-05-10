package hackerone

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	auditLogAPIEndpoint = "%s/v1/programs/%s/audit_log"
)

var (
	clientId     = os.Getenv("HACKERONE_CLIENT_ID")
	clientSecret = os.Getenv("HACKERONE_CLIENT_SECRET")
	programId    = os.Getenv("HACKERONE_PROGRAM_ID")
	baseUrl      = os.Getenv("BASE_URL")
)

func validateEnvVars() {
	if clientId == "" {
		logrus.Fatal("HACKERONE_CLIENT_ID is not set")
	}

	if clientSecret == "" {
		logrus.Fatal("HACKERONE_CLIENT_SECRET is not set")
	}

	if programId == "" {
		logrus.Fatal("HACKERONE_PROGRAM_ID is not set")
	}

	if baseUrl == "" {
		logrus.Fatal("BASE_URL is not set")
	}
}
