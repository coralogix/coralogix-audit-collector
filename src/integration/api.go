package integration

type API interface {
	GetReports(timePeriodManager *TimePeriodManager) (APIResult, error)
}

type APIResult []map[string]interface{}
