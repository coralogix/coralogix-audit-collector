package googleworkspacealertcenter

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	supportedLogTypes = "saml,drive,calendar,login,admin,groups,user_accounts,gcp,mobile"
)

var (
	targetPrincipal      = os.Getenv("GOOGLE_TARGET_PRINCIPAL")
	rawJsonKey           = os.Getenv("GOOGLE_JSON_KEY")
	impersonateUserEmail = os.Getenv("IMPERSONATE_USER_EMAIL")
	maxPageSize          = os.Getenv("MAX_PAGE_SIZE")
)

func validateEnvVars() {
	if impersonateUserEmail == "" {
		logrus.Fatal("IMPERSONATE_USER_EMAIL is not set")
	}
}
