# Integrations for Incident Response

This repository contains a collection of integrations for incident response.

## Requirements

- Docker
- Coralogix API key

## Environment variables

| Variable | Description | Example                                         |
|----------|-------------|-------------------------------------------------|
| `CORALOGIX_LOG_URL` | Coralogix API URL | `https://ingress.eu2.coralogix.com/api/v1/logs` |
| `CORALOGIX_PRIVATE_KEY` | Coralogix private key | `UUID`                                          |
| `CORALOGIX_APP_NAME` | Coralogix application name | `APP_NAME`                                      |


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
- Pritnul
- Jfrog
- Imperva
- SalesForce
- K8s audit
- Teleport

### Built-in integrations in Coralogix

- Cloudtrail

### Webhooks

### Known limitations

- [Slack TBD](#TBD) - Requires an Enterprise Account
- [Notion TBD](#TBD) - Requires an Enterprise Account
- Lastpass doesn't support pagination

## Known issues

- The [Go Coralogix SDK](https://github.com/coralogix/go-coralogix-sdk) is not updated and [contains a bug](https://github.com/coralogix/go-coralogix-sdk/blob/v1.0.3/manager.go#L111-L117). A [patch has been added to this repository](./src/coralogix/coralogix_test.go) to fix the issue. This patch takes care of log sizes greater than the defined limit but still send them in bulks.
