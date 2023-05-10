# GoogleWorkspace's integration

This is a simple script to integrate with admin.google.com's reporting API and report to a Coralogix account. It replaces the [built-in fluentd Google Workspace integration](https://coralogix.com/docs/google-workspace-integration/).

## Requirements

- Enable Admin SDK API - https://console.developers.google.com/apis/api/admin.googleapis.com/overview?project=YOUR_PROJECT_ID
- Create a Service Account with Domain-Wide Delegation - https://developers.google.com/workspace/guides/create-credentials#service-account

### Development/Usage

## Environment variables

| Variable | Description           | Example | Required |
|----------|-----------------------|---------| -------- |
| GOOGLE_JSON_KEY | The Service ACcount JSON key | `{... }` | Yes |
| IMPERSONATE_USER_EMAIL | The user to impersonate | `admin@yourdomain.com` | Yes |
| LOG_TYPES | Comma separated list of log types to fetch | supported: `saml,drive,calendar,login,admin,groups,user_accounts,gcp,mobile` (default)   | No |
| IGNORED_AUDIT_PARAMETERS | Comma separated list of audit parameters to ignore | e.g `IGNORED_AUDIT_PARAMETERS=doc_title` so that the name of the documents won't show in your logs. | No |

## Running

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
