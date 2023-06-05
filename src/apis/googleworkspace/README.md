# GoogleWorkspace's integration

This is a simple script to integrate with admin.google.com's reporting API and report to a Coralogix account. It replaces the [built-in fluentd Google Workspace integration](https://coralogix.com/docs/google-workspace-integration/).

## Requirements

- Enable Admin SDK API - https://console.developers.google.com/apis/api/admin.googleapis.com/overview?project=YOUR_PROJECT_ID
- Create a Service Account with Domain-Wide Delegation - https://developers.google.com/workspace/guides/create-credentials#service-account

### Development/Usage

## Environment variables

| Variable | Description           | Example | Required |
|----------|-----------------------|---------| -------- |
| IMPERSONATE_USER_EMAIL | The user to impersonate | `admin@yourdomain.com` | Yes |
| GOOGLE_TARGET_PRINCIPAL | The service account to impersonate | `...@....iam.gserviceaccount.com` | Yes |
| GOOGLE_APPLICATION_CREDENTIALS | The Service Account JSON key if not using the default `GOOGLE_APPLICATION_CREDENTIALS`  | `{... }` | NO |
| GOOGLE_JSON_KEY | The Service Account JSON key if not using the default `GOOGLE_APPLICATION_CREDENTIALS`  | `{... }` | NO |
| LOG_TYPES | Comma separated list of log types to fetch | supported: `saml,drive,calendar,login,admin,groups,user_accounts,gcp,mobile` (default)   | No |
| IGNORED_AUDIT_PARAMETERS | Comma separated list of audit parameters to ignore | e.g `IGNORED_AUDIT_PARAMETERS=doc_title` so that the name of the documents won't show in your logs. | No |

## Running (prod)

```shell
export IMPERSONATE_USER_EMAIL="user@yourdomain"
export INTEGRATION_SEARCH_DIFF_IN_MINUTES="5"
export GOOGLE_TARGET_PRINCIPAL="...@....iam.gserviceaccount.com"
docker run -it --rm \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e INTEGRATION_SEARCH_DIFF_IN_MINUTES="$INTEGRATION_SEARCH_DIFF_IN_MINUTES"
    -e INTEGRATION_NAME="googleworkspace" \
    -e IMPERSONATE_USER_EMAIL="$IMPERSONATE_USER_EMAIL" \
    -e BASE_URL="$BASE_URL" \
    -e DRY_RUN="true" \
    coralogixrepo/coralogix-audit-collector
```

## Running (dev)

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e GOOGLE_JSON_KEY="$GOOGLE_JSON_KEY" \
    -e IMPERSONATE_USER_EMAIL="$IMPERSONATE_USER_EMAIL" \
    -e LOG_TYPES="$LOG_TYPES" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- https://developers.google.com/admin-sdk/reports/v1/appendix/activity/admin
- https://developers.google.com/workspace/guides/create-credentials#service-account
- https://github.com/googleapis/google-api-go-client
