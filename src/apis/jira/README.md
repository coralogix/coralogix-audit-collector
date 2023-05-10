# JIRA integration

This is a simple script to integrate with JIRA's reporting API and report to a Coralogix account.

## Requirements

- [JIRA API key](https://id.atlassian.com/manage-profile/security/api-tokens)

### Development/Usage

## Environment variables

| Variable | Description                                 | Example         |
|----------|---------------------------------------------|-----------------|
 | `JIRA_API_TOKEN` | API token generated from the user's account | `TOKEN`         |
 | `JIRA_USERNAME` | The JIRA username                           | `user@user.com` |

## Running

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e JIRA_USERNAME="$JIRA_USERNAME" \
    -e JIRA_API_TOKEN="$JIRA_API_TOKEN" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## Known issues

Jira API hides the user's information due to GDPR. This topic is discussed in [this thread](https://jira.atlassian.com/browse/JRACLOUD-77455) and handled by using the [user/migrate](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-bulk-migration-get) API.

## References

- [Basic Auth](https://developer.atlassian.com/cloud/jira/platform/basic-auth-for-rest-apis/)
- [Authentication](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/#authentication)
- [JIRA API](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/)
- [Audit Records](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-audit-records/#api-rest-api-3-auditing-record-get)
