# Audit Logs Collector

This chart create a cronjob that collects audit logs from different sources and sends them to Coralogix.

## Requirements

- A Coralogix account

## Installing the Chart

Using the [.env.example](.env.example) file, create a secret with your integrations configuration:

```bash
export NAMESPACE="coralogix-audit-collector"

kubectl -n $NAMESPACE create secret generic \
    coralogix-audit-collector-secret \
    --from-env-file=./.env.example \
    --save-config \
    --dry-run=client \
    -o yaml |\
     kubectl -n $NAMESPACE apply -f -
```

Then install the chart:

```bash
helm upgrade --install coralogix-audit-collector \
    coralogix-audit-collector \
    --namespace $NAMESPACE \
    --create-namespace \
    --values ./values.yaml \
    --repo https://cgx.jfrog.io/artifactory/coralogix-charts-virtual
```

### Values.yaml

Each integration contains the following values:

| Parameter | Description | Default | Required                                                |
|-----------|-------------|---------|---------------------------------------------------------|
| `enabled` | Whether to enable the integration | `false` | Yes |
| `baseUrl` | Base URL for the integration | `""` | Yes |
| `schedule` | Cron schedule | "" | No - if not defined `.Values.cron.schedule` will be used |

## Integrations

Each integration has its own requirements. The supported integrations are listed in the main [README](../README.md).
