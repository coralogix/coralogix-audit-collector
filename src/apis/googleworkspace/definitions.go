package googleworkspace

import (
	"github.com/coralogix/c4c-ir-integrations/src/utils"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	supportedLogTypes = "saml,drive,calendar,login,admin,groups,user_accounts,gcp,mobile"
)

var (
	targetPrincipal              = os.Getenv("GOOGLE_TARGET_PRINCIPAL")
	googleApplicationCredentials = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	rawJsonKey                   = os.Getenv("GOOGLE_JSON_KEY")
	impersonateUserEmail         = os.Getenv("IMPERSONATE_USER_EMAIL")
	logTypesStr                  = os.Getenv("LOG_TYPES")
	ignoredAuditParametersStr    = os.Getenv("IGNORED_AUDIT_PARAMETERS")
	logTypes                     []string
	ignoredAuditParameters       []string
)

func validateEnvVars() {
	if impersonateUserEmail == "" {
		logrus.Fatal("IMPERSONATE_USER_EMAIL is not set")
	}

	if logTypesStr == "" {
		logTypesStr = supportedLogTypes
	}

	logTypes = utils.SplitAndRemoveSpaces(logTypesStr)
	if len(logTypes) == 0 {
		logrus.Fatal("LOG_TYPES is empty")
	}

	if checkSupportedLogTypes(logTypes) == false {
		logrus.Fatal("LOG_TYPES contains unsupported log type")
	}

	if ignoredAuditParametersStr != "" {
		ignoredAuditParameters = utils.SplitAndRemoveSpaces(ignoredAuditParametersStr)
	}
}

func checkSupportedLogTypes(logTypes []string) bool {
	for _, logType := range logTypes {
		if strings.Contains(supportedLogTypes, logType) == false {
			return false
		}
	}
	return true
}
