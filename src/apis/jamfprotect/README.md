# JAMF protect integration

This is a simple script to integrate with JAMF protect's reporting API and report to a Coralogix account.

## Requirements

- API token - https://your-jamfcloud.com/api-clients/?sort=created.desc&column=&type=&value=&filterType=

### Development/Usage

## Environment variables

| Variable              | Description                          | Example     |
|-----------------------|--------------------------------------|-------------|
 | `JAMF_PROTECT_CLIENT_ID` | API token generated from API clients | `TOKEN`     |
 | `JAMF_PROTECT_API_TOKEN` | The client id                        | `CLIENT_ID` |

## Running

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e JAMF_PROTECT_CLIENT_ID="$JAMF_PROTECT_CLIENT_ID" \
    -e JAMF_PROTECT_API_TOKEN="$JAMF_PROTECT_API_TOKEN" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- [GraphQL Queries](https://learn.jamf.com/bundle/jamf-protect-documentation/page/Queries_and_Mutations.html)
- [Authorization API](https://learn.jamf.com/bundle/jamf-protect-documentation/page/Jamf_Protect_API.html)
- [Examples](https://github.com/jamf/jamfprotect/blob/main/jamf_protect_api/scripts/python/list_audit_logs.py)
