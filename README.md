# Coralogix Audit Collector

This project integrates Coralogix with various APIs to collect audit logs and send them to Coralogix.

## Requirements

- Docker
- Coralogix API key

## Environment variables

| Variable | Description | Example                                         |
|----------|-------------|-------------------------------------------------|
| `CORALOGIX_LOG_URL` | Coralogix API URL | `https://ingress.eu2.coralogix.com/api/v1/logs` |
| `CORALOGIX_PRIVATE_KEY` | Coralogix private key | `UUID`                                          |
| `CORALOGIX_APP_NAME` | Coralogix application name | `APP_NAME`                                      |

## Usage

### Chart

A helm chart is available [here](./chart/README.md), TL;DR:

### Docker

You can also run each integration by itself using `docker`. 

```bash
docker run -it --rm  \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e INTEGRATION_NAME="$INTEGRATION_NAME" \
    coralogixrepo/coralogix-audit-collector
```

#### JIRA Example

```bash
docker run -it --rm \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e JIRA_USERNAME="$JIRA_USERNAME" \
    -e JIRA_API_TOKEN="$JIRA_API_TOKEN" \
    -e INTEGRATION_NAME="jira" \
    -e BASE_URL="https://your-org.atlassian.net" \
    -e DRY_RUN="true" \
    coralogixrepo/coralogix-audit-collector
```

#### Google Workspace Example

`GOOGLE_APPLICATION_CREDENTIALS` is optional. Running this directly on a GCP instance will use the instance's service account.

```bash
GOOGLE_APPLICATION_CREDENTIALS="/app/src/apis/googleworkspace/credentials.json"
GOOGLE_TARGET_PRINCIPAL="...@....iam.gserviceaccount.com"
LOG_TYPE="saml,drive,calendar,login,admin,groups,user_accounts,gcp,mobile"

docker run -it --rm \
    -e CORALOGIX_LOG_URL="https://ingress.eu2.coralogix.com/api/v1/logs" \
    -e CORALOGIX_PRIVATE_KEY="$CORALOGIX_PRIVATE_KEY" \
    -e CORALOGIX_APP_NAME="$CORALOGIX_APP_NAME" \
    -e GOOGLE_APPLICATION_CREDENTIALS="$GOOGLE_APPLICATION_CREDENTIALS" \
    -e GOOGLE_TARGET_PRINCIPAL="$GOOGLE_TARGET_PRINCIPAL" \
    -e IMPERSONATE_USER_EMAIL="$IMPERSONATE_USER_EMAIL" \
    -e INTEGRATION_NAME="googleworkspace" \
    -e LOG_TYPES="$LOG_TYPES" \
    -e DRY_RUN="true" \
    coralogixrepo/coralogix-audit-collector
```

### Notes

* You can avoid sending logs to Coralogix by setting `-e DRY_RUN=true`.
* You can include the debug logs by setting `-e DEBUG=true`.
* You can override the default time range by setting `-e INTEGRATION_SEARCH_DIFF_IN_MINUTES=60m`.

Visit the integration's README for more information.

## Integrations

### Cron

- [Lastpass](src/apis/lastpass/README.md)
- [Monday](src/apis/monday/README.md)
- [Zoom](src/apis/zoom/README.md)
- [Intercom](src/apis/intercom/README.md)
- [Jira](src/apis/jira/README.md)
- [Confluence](src/apis/confluence/README.md)
- [JAMF Protect](src/apis/jamfprotect/README.md)
- [Atlassian](src/apis/atlassian/README.md)
- [HackerOne](src/apis/hackerone/README.md)
- [Google Workspace](src/apis/googleworkspace/README.md)
- [Jfrog](src/apis/jfrog/README.md)
- [Slack](src/apis/slack/README.md)

### TBD

- Cloudtrail
- Notion TBD
- Pritnul
- Imperva
- SalesForce
- K8s audit
- Teleport
- Replace `INTEGRATION_SEARCH_DIFF_IN_MINUTES` with `INTEGRATION_SEARCH_DIFF=1h` (or `1d`, or `30m`, etc.).

## Known issues

- The [Go Coralogix SDK](https://github.com/coralogix/go-coralogix-sdk) is not updated and [contains a bug](https://github.com/coralogix/go-coralogix-sdk/blob/v1.0.3/manager.go#L111-L117). A [patch has been added to this repository](./src/coralogix/coralogix_test.go) to fix the issue. This patch takes care of log sizes greater than the defined limit but still send them in bulks.
