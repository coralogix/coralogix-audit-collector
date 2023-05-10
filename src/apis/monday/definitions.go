package monday

import (
	"os"
	"time"
)

const (
	resultsPerPage = 100
	timeFormat     = time.RFC3339
	userPage       = "%s/users/%s"
	apiEndpoint    = "%s/audit-api/get-logs"
)

var (
	apiToken = os.Getenv("MONDAY_API_TOKEN")
	baseUrl  = os.Getenv("BASE_URL")
)

func validateEnvVars() {
	if apiToken == "" {
		panic("MONDAY_API_TOKEN env var is not set")
	}
	if baseUrl == "" {
		panic("BASE_URL env var is not set")
	}
}
