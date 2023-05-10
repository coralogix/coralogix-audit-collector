# Lastpass integration

This is a simple script to integrate with lastpass' reporting API and report to a Coralogix account.

## Requirements

- Lastpass API key

### Development/Usage

## Environment variables

| Variable | Description | Example                                         |
|----------|-------------|-------------------------------------------------|
| `LASTPASS_CID` | Lastpass customer ID | `ACCOUNT_NUMBER`                                |
| `LASTPASS_PROVHASH` | Lastpass provisioning hash | `HASH`                                          |
| `LASTPASS_APIUSER` | Lastpass API user | `USERNAME`                                      |

## Running

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e LASTPASS_CID="$LASTPASS_CID" \
    -e LASTPASS_PROVHASH="$LASTPASS_PROVHASH" \
    -e LASTPASS_APIUSER="default-test-user" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- [Reporting API](https://support.lastpass.com/help/event-reporting-via-lastpass-api)
