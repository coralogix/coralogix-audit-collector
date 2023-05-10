# Confluence integration

This is a simple script to integrate with Confluence's reporting API and report to a Coralogix account.

## Requirements

- [Confluence API key](https://id.atlassian.com/manage-profile/security/api-tokens)

### Development/Usage

## Environment variables

| Variable | Description                                 | Example         |
|----------|---------------------------------------------|-----------------|
 | `CONFLUENCE_API_TOKEN` | API token generated from the user's account | `TOKEN`         |
 | `CONFLUENCE_USERNAME` | The Confluence username                           | `user@user.com` |

## Running

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e CONFLUENCE_USERNAME="$CONFLUENCE_USERNAME" \
    -e CONFLUENCE_API_TOKEN="$CONFLUENCE_API_TOKEN" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- [Basic Auth](https://developer.atlassian.com/cloud/jira/platform/basic-auth-for-rest-apis/)
- [Authentication](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/#authentication)
- [JIRA API](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/)
- [Audit Records](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-audit-records/#api-rest-api-3-auditing-record-get)
