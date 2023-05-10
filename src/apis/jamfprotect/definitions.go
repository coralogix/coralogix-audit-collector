package jamfprotect

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	accessTokenUri = "/token"
	apiUri         = "/graphql"
	query          = `
query listAuditLogsByDate($next: String, $pageSize: Int, $order: AuditLogsOrderInput, $condition: AuditLogsDateConditionInput) {
      listAuditLogsByDate(
        input: {next: $next, pageSize: $pageSize, order: $order, condition: $condition}
  ) {
        items {
          ...AuditLogFields
      __typename
    }
    pageInfo {
          next
      __typename
    }
    __typename
  }
}

fragment AuditLogFields on AuditLog {
      resourceId
  date
  args
  error
  ips
  op
  user
  __typename
}`
)

var (
	clientId = os.Getenv("JAMF_PROTECT_CLIENT_ID")
	apiToken = os.Getenv("JAMF_PROTECT_API_TOKEN")
	baseUrl  = os.Getenv("BASE_URL")
)

func validateEnvVars() {
	if clientId == "" {
		logrus.Fatalf("JAMF_PROTECT_CLIENT_ID env var is not set")
	}
	if apiToken == "" {
		logrus.Fatalf("JAMF_PROTECT_API_TOKEN env var is not set")
	}
	if baseUrl == "" {
		logrus.Fatalf("BASE_URL env var is not set")
	}
}
