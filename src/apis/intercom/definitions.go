package intercom

import "os"

var (
	accessToken = os.Getenv("INTERCOM_ACCESS_TOKEN")
	baseUrl     = os.Getenv("BASE_URL")
	apiEndpoint = "%s/admins/activity_logs?created_at_after=%s&created_at_before=%s"
)

func validateEnvVars() {
	if accessToken == "" {
		panic("INTERCOM_ACCESS_TOKEN env var is not set")
	}
	if baseUrl == "" {
		panic("BASE_URL env var is not set")
	}
}
