# Atlassian Organization integration

This is a simple script to integrate with Atlassian Organization's reporting API and report to a Coralogix account.

## Requirements

- [Atlassian API key](https://admin.atlassian.com/)

### Development/Usage

## Environment variables

| Variable              | Description                                 | Example     |
|-----------------------|---------------------------------------------|-------------|
 | `ATLASSIAN_API_TOKEN` | API token generated from the user's account | `TOKEN`     |
 | `ATLASSIAN_CLIENT_ID` | The Atlassian client id                     | `CLIENT_ID` |

## Running

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e ATLASSIAN_CLIENT_ID="$ATLASSIAN_CLIENT_ID" \
    -e ATLASSIAN_API_TOKEN="$ATLASSIAN_API_TOKEN" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- [audit API](https://developer.atlassian.com/cloud/admin/organization/rest/api-group-events/)
