# Monday integration

This is a simple script to integrate with monday.com's reporting API and report to a Coralogix account.

## Requirements

- Monday API key

### Development/Usage

## Environment variables

| Variable | Description                                                              | Example                                         |
|----------|--------------------------------------------------------------------------|-------------------------------------------------|
 | `MONDAY_API_TOKEN` | API token generated from https://COMPANY.monday.com/admin/security/login | `TOKEN`                                         |

## Running

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e MONDAY_API_TOKEN="$MONDAY_API_TOKEN" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- [Monday Audit Log API](https://support.monday.com/hc/en-us/articles/4406042650002-Audit-Log-API)
