package tests

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

func init() {
	debug := os.Getenv("DEBUG") == "true"
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

type MockClient struct {
	FakeDo func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.FakeDo(req)
}

func NewMockWithFixtureByUrlPath(json map[string][]string) *MockClient {
	return &MockClient{
		FakeDo: func(req *http.Request) (*http.Response, error) {
			logrus.Debugf("req.URL.Path: %s %d", req.URL.Path, len(json[req.URL.Path]))
			if len(json[req.URL.Path]) == 0 {
				logrus.Fatalf("Invalid path: %s", req.URL.Path)
			}
			b := []byte(json[req.URL.Path][0])
			json[req.URL.Path] = json[req.URL.Path][1:]

			r := io.NopCloser(bytes.NewReader(b))
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
}
