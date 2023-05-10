# Intercom integration

This is a simple script to integrate with intercom.com's reporting API and report to a Coralogix account.

## Requirements

- [Intercom API key](https://app.intercom.com/a/apps/xf85kx3u/developer-hub/app-packages/92741/basic-info)

### Development/Usage

## Environment variables

| Variable | Description                              | Example                                         |
|----------|------------------------------------------|-------------------------------------------------|
 | `INTERCOM_ACCESS_TOKEN` | API token generated from an intercom app | `TOKEN`                                         |

## Running

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e INTERCOM_ACCESS_TOKEN="$INTERCOM_ACCESS_TOKEN" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- [Admin Activity Log API](https://developers.intercom.com/intercom-api-reference/reference/list-all-activity-logs)
- [Authentication](https://developers.intercom.com/building-apps/docs/authentication-types)
