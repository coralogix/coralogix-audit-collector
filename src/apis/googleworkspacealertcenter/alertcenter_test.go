package googleworkspacealertcenter

import (
	"bytes"
	"context"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	alertcenter "google.golang.org/api/alertcenter/v1beta1"
	"google.golang.org/api/option"
	"io"
	"net/http"
	"testing"
)

type mockRoundTripper struct {
	response *http.Response
}

func (rt *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt.response, nil
}

type serviceMemory struct {
	alerts *alertcenter.AlertsService
}

func (g *serviceMemory) List() *alertcenter.AlertsListCall {
	return g.alerts.List()
}

func TestGoogleWorkspace_GetReports(t *testing.T) {
	timePeriodManager := integration.NewTimePeriodManagerFromNow(1)
	mockClient := &http.Client{
		Transport: &mockRoundTripper{
			response: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString(`{"alerts":[{"alertId":"b5824860-b7ca-4e9f-91ee-04cd528a00f7","createTime":"2023-07-06T08:57:23.334247Z","customerId":"049wggmi","data":{"@type":"type.googleapis.com/google.apps.alertcenter.type.MailPhishing","domainId":{"customerPrimaryDomain":"test.com"},"maliciousEntity":{"displayName":"test","entity":{"displayName":"test","emailAddress":"test.c5.73f@reply.test"}},"messages":[{"attachmentsSha256Hash":["30719d347bf4e87749232d66aa4dd2465417d81e5f4fc842efec515ced7ee844"],"date":"2023-07-06T07:49:04.140079Z","md5HashMessageBody":"3367166344828396ee292f1d380040eb","md5HashSubject":"bb67a5fac285b65d136a857453492f85","messageId":"20230706074902.a0dbd0d115c9632d@test","recipient":"test@test.com","sentTime":"2023-07-06T07:49:02Z"}]},"endTime":"2023-07-06T08:47:46.298204Z","etag":"1Ao0dAFQMvo=","metadata":{"alertId":"b5824860-b7ca-4e9f-91ee-04cd528a00f7","customerId":"049wggmi","etag":"1Ao0dAFQMvo=","severity":"MEDIUM","status":"NOT_STARTED","updateTime":"2023-07-06T08:57:23.334247Z"},"source":"Gmail phishing","startTime":"2023-07-06T07:49:04.140079Z","type":"Gmail potential employee spoofing","updateTime":"2023-07-06T08:57:23.334247Z"},{}]}`)),
			},
		},
	}

	s, _ := alertcenter.NewService(context.Background(), option.WithHTTPClient(mockClient))
	svc := &serviceMemory{
		alerts: s.Alerts,
	}
	gw := New(svc, 10)
	auditLogs, err := gw.GetReports(timePeriodManager)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(auditLogs) != 2 {
		t.Errorf("Expected 2 audit log, got %d", len(auditLogs))
	}

	auditLog := auditLogs[0]
	alertType := auditLog["type"].(string)
	if alertType != "Gmail potential employee spoofing" {
		t.Errorf("Expected Gmail potential employee spoofing, got %s", alertType)
	}
}
