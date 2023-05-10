# Zoom integration

This is a simple script to integrate with zoom.com's reporting API and report to a Coralogix account.

## Requirements

- [ZOOM App key](https://marketplace.zoom.us/develop/apps/ef_R3-ilQNGBUPwP91IVpg/activation)

### Development/Usage

## Environment variables

| Variable | Description               | Example                                         |
|----------|---------------------------|-------------------------------------------------|
 | `ZOOM_ACCOUNT_ID` | Generated from a zoom app | `TOKEN`                                         |
 | `ZOOM_CLIENT_ID` | Generated from a zoom app | `TOKEN`                                         |
 | `ZOOM_CLIENT_SECRET` | Generated from a zoom app | `TOKEN`                                         |

## Running

```
docker run -it --rm \
    -w "/app/src" \
    -e "air_wd=/app/src" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e ZOOM_ACCOUNT_ID="$ZOOM_ACCOUNT_ID" \
    -e ZOOM_CLIENT_ID="$ZOOM_CLIENT_ID" \
    -e ZOOM_CLIENT_SECRET="$ZOOM_CLIENT_SECRET" \
    -v $(pwd):/app/src \
    -p 6000:6000 \
    cosmtrek/air -c air.toml
```

## References

- [Report Operation Log API](https://marketplace.zoom.us/docs/api-reference/zoom-api/methods/#operation/reportOperationLogs)
- [Authentication](https://marketplace.zoom.us/docs/guides/build/server-to-server-oauth-app/)
