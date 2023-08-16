# Google Workspace Alert Center Integration

A simple script to integrate with admin.google.com's alert center API and report to a Coralogix account.

## Requirements

- Ensure to correct APIs are enabled 
  - IAM Service Account Credentials API - https://console.cloud.google.com/apis/library/iamcredentials.googleapis.com
  - Google Workspace Alert Center API - https://console.cloud.google.com/marketplace/product/google/alertcenter.googleapis.com
- Create a Service Account with Domain-Wide Delegation - https://developers.google.com/workspace/guides/create-credentials#service-account
- The Domain-Wide Delegation OAuth scopes for Alert Center is - https://www.googleapis.com/auth/apps.alerts

### Development/Usage

## Environment variables

| Variable | Description                                                                            | Example | Required |
|----------|----------------------------------------------------------------------------------------|---------| -------- |
| IMPERSONATE_USER_EMAIL | The user to impersonate                                                                | `admin@yourdomain.com` | Yes |
| GOOGLE_TARGET_PRINCIPAL | The service account to impersonate, used with `GOOGLE_APPLICATION_CREDENTIALS`         | `...@....iam.gserviceaccount.com` | Yes |
| GOOGLE_APPLICATION_CREDENTIALS | The Service Account JSON key if not using the default `GOOGLE_APPLICATION_CREDENTIALS` | `{... }` | NO |
| GOOGLE_JSON_KEY | The Service Account JSON key if not using the default `GOOGLE_APPLICATION_CREDENTIALS` | `{... }` | NO |

## Running (prod)

```shell
export IMPERSONATE_USER_EMAIL="user@yourdomain"
export INTEGRATION_SEARCH_DIFF_IN_MINUTES="5"
# IF USING GOOGLE_APPLICATION_CREDENTIALS then use GOOGLE_APPLICATION_CREDENTIALS with GOOGLE_TARGET_PRINCIPAL
export GOOGLE_TARGET_PRINCIPAL="...@....iam.gserviceaccount.com"
docker run -it --rm \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e INTEGRATION_SEARCH_DIFF_IN_MINUTES="$INTEGRATION_SEARCH_DIFF_IN_MINUTES"
    -e INTEGRATION_NAME="googleworkspacealertcenter" \
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
    -e INTEGRATION_NAME="googleworkspacealertcenter" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- https://developers.google.com/admin-sdk/alertcenter/guides
- https://developers.google.com/workspace/guides/create-credentials#service-account
