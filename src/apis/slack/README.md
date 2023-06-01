# Slack integration

## Requirements

- Enterprise Grid account
- [Follow the steps for the audit logs](https://api.slack.com/admins/audit-logs)
- [Create a Slack app](https://api.slack.com/apps)
- [Slack Audit Log API](https://api.slack.com/admins/audit-logs-call)


## Usage 

### Environment variables

| Variable          | Description                                            | Example                 |
|-------------------|--------------------------------------------------------|-------------------------|
 | `SLACK_ACCESS_TOKEN`  | The generated acces token after the apps' installation | `xoxp-....`             |
| `BASE_URL`        | The base URL for the Slack instance                    | `https://api.slack.com` |

### Running (prod)

```
docker run -it --rm \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e INTEGRATION_NAME="slack" \
    -e SLACK_ACCESS_TOKEN="SLACK_ACCESS_TOKEN" \
    -e BASE_URL="$BASE_URL" \
    -e DRY_RUN="true" \
    coralogixrepo/coralogix-audit-collector
```

### Running (dev)

```
docker run -it --rm \                                                                                                                          02:02:27
    -w "/app/src" \
    -e INTEGRATION_SEARCH_DIFF_IN_MINUTES="1440" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e INTEGRATION_NAME="slack" \
    -e SLACK_ACCESS_TOKEN="SLACK_ACCESS_TOKEN" \
    -e BASE_URL="$BASE_URL" \
    -p 6000:6000 \
    -v $(pwd):/app/src \
    cosmtrek/air -c air.toml
```

## You should know

- You msut have an Enterprise Grid account for this integration to work
- Generating the access token is a manual process (see [README.md](./app/README.md))
