
## References

### Cloud Log Collection Output
- [Cloud Log Collection Output](https://www.jfrog.com/confluence/display/JFROG/Cloud+Log+Collection#CloudLogCollection-OptingOut)

#### Collection Repository

> The logs will be collected in a dedicated system repository called jfrog-logs, which is created when the API is called. The repository will be displayed in the Platform UI (if the user has admin permissions) together with the other repositories (such as build-info, etc.). The repository will also have a dedicated icon to indicate it.

#### Collection Frequency

> Logs will be uploaded once in 24 hours or when 25 MB are reached (the earliest).

### Enabling Log Collection
- [Enable Log Collection](https://www.jfrog.com/confluence/display/JFROG/JFrog+Platform+REST+API#JFrogPlatformRESTAPI-EnableLogCollection)

```bash
curl -XPOST  \
    -H "Content-Type: application/json" \
    -d'{"enabled": true}' \
    -H "Authorization: Bearer $JFROG_ADMIN_TOKEN" \
    -v https://$JFROG_PLATFORM_URL/access/api/v1/logshipping/config
```

### Possible Solution

Using the following command, we can get the latest (yesterday) log file from the `jfrog-logs` repository.

```shell
export JFROG_URL="JFROG_URL"
export JFROG_USERNAME="JFROG_USERNAME"
export JFROG_TOKEN="JFROG_TOKEN"
YESTERDAY=$(date -d "yesterday 23:50" '+%Y-%m-%d')
curl -u "$JFROG_USERNAME:$JFROG_TOKEN" \
    '$JFROG_URL/artifactory/api/search/artifact?name=*access-security*$YESTERDAY*&repos=jfrog-logs' |\
    jq '.results[0]["uri"]'
```

This can run once a day because the log is generated once a day.
