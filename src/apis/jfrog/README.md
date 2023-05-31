
# Jfrog Integration

This is a simple script to integrate with Jfrog's reporting API and report to a Coralogix account. This integration uses the [Jfrog Audit Trail Log](https://jfrog.com/help/r/jfrog-platform-administration-documentation/audit-trail-log) to gather audit log information.

## Requirements

- [Jfrog Authentication Token](https://www.jfrog.com/confluence/display/JFROG/Access+Tokens#AccessTokens-UsingAccessTokens)
- [Enable Log Collection](https://www.jfrog.com/confluence/display/JFROG/JFrog+Platform+REST+API#JFrogPlatformRESTAPI-EnableLogCollection)

## Usage 

### Environment variables

| Variable          | Description                              | Example                 |
|-------------------|------------------------------------------|-------------------------|
 | `JFROG_USERNAME`  | The app username                         | `somename`              |
| `JFROG_API_TOKEN` | The generated API token                  | `a-65-chars-string`     |
| `BASE_URL`        | The base URL for the JFrog instance     | `https://name.jfrog.io` |

### Running (prod)

```
docker run -it --rm \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e INTEGRATION_NAME="jfrog" \
    -e JFROG_USERNAME="$JFROG_USERNAME" \
    -e JFROG_API_TOKEN="$JFROG_API_TOKEN" \
    -e BASE_URL="$BASE_URL" \
    -e DRY_RUN="true" \
    coralogixrepo/audit-logs-collector
```

### Running (dev)

```
docker run -it --rm \                                                                                                                          02:02:27
    -w "/app/src" \
    -e INTEGRATION_SEARCH_DIFF_IN_MINUTES="1440" \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e INTEGRATION_NAME="jfrog" \
    -e JFROG_USERNAME="$JFROG_USERNAME" \
    -e JFROG_API_TOKEN="$JFROG_API_TOKEN" \
    -e BASE_URL="$BASE_URL" \
    -p 6000:6000 \
    -v $(pwd):/app/src \
    cosmtrek/air -c air.toml
```

## You should know

> The logs will be collected in a dedicated system repository called jfrog-logs, which is created when the API is called. The repository will be displayed in the Platform UI (if the user has admin permissions) together with the other repositories (such as build-info, etc.). The repository will also have a dedicated icon to indicate it.

> Logs will be uploaded once in 24 hours or when 25 MB are reached (the earliest).

https://jfrog.com/help/r/jfrog-platform-administration-documentation/cloud-log-collection

## References
- 
- [Cloud Log Collection Output](https://www.jfrog.com/confluence/display/JFROG/Cloud+Log+Collection#CloudLogCollection-OptingOut)
