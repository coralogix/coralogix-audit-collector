package lastpass

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	timeFormat  = "2006-01-02 15:04:05"
	apiEndpoint = "%s/enterpriseapi.php"
)

var (
	cid      = os.Getenv("LASTPASS_CID")
	provhash = os.Getenv("LASTPASS_PROVHASH")
	apiuser  = os.Getenv("LASTPASS_APIUSER")
	baseUrl  = os.Getenv("BASE_URL")
)

func validateEnvVars() {
	if baseUrl == "" {
		logrus.Fatalf("Missing environment variable: BASE_URL")
	}

	if cid == "" {
		logrus.Fatalf("Missing environment variable: LASTPASS_CID")
	}

	if provhash == "" {
		logrus.Fatalf("Missing environment variable: LASTPASS_PROVHASH")
	}

	if apiuser == "" {
		logrus.Fatalf("Missing environment variable: LASTPASS_APIUSER")
	}
}
