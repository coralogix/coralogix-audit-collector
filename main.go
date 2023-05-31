package main

import (
	"flag"
	"github.com/coralogix/c4c-ir-integrations/src/apis/confluence"
	"github.com/coralogix/c4c-ir-integrations/src/apis/googleworkspace"
	"github.com/coralogix/c4c-ir-integrations/src/apis/hackerone"
	"github.com/coralogix/c4c-ir-integrations/src/apis/intercom"
	"github.com/coralogix/c4c-ir-integrations/src/apis/jamfprotect"
	"github.com/coralogix/c4c-ir-integrations/src/apis/jfrog"
	"github.com/coralogix/c4c-ir-integrations/src/apis/jira"
	"github.com/coralogix/c4c-ir-integrations/src/apis/lastpass"
	"github.com/coralogix/c4c-ir-integrations/src/apis/monday"
	"github.com/coralogix/c4c-ir-integrations/src/apis/zoom"
	"github.com/coralogix/c4c-ir-integrations/src/coralogix"
	debug2 "github.com/coralogix/c4c-ir-integrations/src/debug"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var (
	dryRun                         = os.Getenv("DRY_RUN") == "true"
	integrationSearchDiffInMinutes = os.Getenv("INTEGRATION_SEARCH_DIFF_IN_MINUTES")
	integrationName                = os.Getenv("INTEGRATION_NAME")
	debug                          = os.Getenv("DEBUG")
	integrations                   = map[string]func() integration.API{
		"monday":          monday.NewFromEnv,
		"lastpass":        lastpass.NewFromEnv,
		"intercom":        intercom.NewFromEnv,
		"zoom":            zoom.NewFromEnv,
		"hackerone":       hackerone.NewFromEnv,
		"jamfprotect":     jamfprotect.NewFromEnv,
		"confluence":      confluence.NewFromEnv,
		"jira":            jira.NewFromEnv,
		"googleworkspace": googleworkspace.NewFromEnv,
		"jfrog":           jfrog.NewFromEnv,
	}
)

func init() {
	if debug == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

type Collector interface {
	Collect(v map[string]interface{})
	Send()
}

func main() {
	flag.Parse()
	runIntegration(integrationName, integrationSearchDiffInMinutes, dryRun)
}

func runIntegration(integrationName, integrationSearchDiffInMinutes string, dryRun bool) {
	integrationSearchDiffInMinutesInt, _ := strconv.Atoi(integrationSearchDiffInMinutes)
	timePeriodManager := integration.NewTimePeriodManagerFromNow(integrationSearchDiffInMinutesInt)
	api := integrations[integrationName]()
	var collector Collector
	if dryRun {
		collector = debug2.NewDryRunCollector()
	} else {
		collector = coralogix.NewCollector(integrationName)
	}

	useCase := NewPublishReportsUseCase(api, collector)
	err := useCase.execute(timePeriodManager)
	if err != nil {
		panic(err)
	}
}

func NewPublishReportsUseCase(api integration.API, collector Collector) *publishReportsUseCase {
	return &publishReportsUseCase{
		api:       api,
		collector: collector,
	}
}

type publishReportsUseCase struct {
	api       integration.API
	collector Collector
}

func (u *publishReportsUseCase) execute(timePeriodManager *integration.TimePeriodManager) error {
	reporting, err := u.api.GetReports(timePeriodManager)
	if err != nil {
		return err
	}

	logrus.Debugf("Found %d reports", len(reporting))
	for _, v := range reporting {
		u.collector.Collect(v)
	}
	u.collector.Send()
	return nil
}
