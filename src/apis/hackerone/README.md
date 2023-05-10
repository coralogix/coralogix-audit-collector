# Hackerone's integration

This is a simple script to integrate with hackerone.com's reporting API and report to a Coralogix account.

## Requirements

- https://hackerone.com/YOUR_COMPANY/api

### Development/Usage

## Environment variables

| Variable | Description           | Example |
|----------|-----------------------|---------|
 | `HACKERONE_PROGRAM_ID` | The H1 program number | `12345` |
 | `HACKERONE_CLIENT_ID` | Client ID             | `TOKEN` |
 | `HACKERONE_CLIENT_SECRET` | Client Secret         | `TOKEN` |

## Running

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e HACKERONE_PROGRAM_ID="$HACKERONE_PROGRAM_ID" \
    -e HACKERONE_CLIENT_ID="$HACKERONE_CLIENT_ID" \
    -e HACKERONE_CLIENT_SECRET="$HACKERONE_CLIENT_SECRET" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- [Program Audit Log](https://api.hackerone.com/customer-resources/#programs-get-audit-log)
