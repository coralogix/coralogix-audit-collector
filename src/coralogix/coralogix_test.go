package coralogix

import (
	coralogixsdk "github.com/coralogix/go-coralogix-sdk"
	"github.com/sirupsen/logrus"
	"io"
	"testing"
)

func TestCollector_Send(t *testing.T) {
	var err error
	c := NewCollector("test")
	testData := []struct {
		length      int
		expectedEOF int
	}{
		{length: 100, expectedEOF: 2},
		{length: 1000, expectedEOF: 2},
		{length: 10000, expectedEOF: 2},
		{length: 100000, expectedEOF: 3},
	}

	for _, d := range testData {
		buffer := &coralogixsdk.LogBuffer{}
		for i := 0; i < d.length; i++ {
			buffer.Append(coralogixsdk.Log{
				Timestamp: 1,
				Text:      "test",
			})
		}

		i := 1
		for {
			_, err = c.sendBulk(buffer)
			logrus.Debugf("Buffer size is %d (%d)", buffer.Size(), buffer.Len())
			if d.expectedEOF == i && err != io.EOF {
				t.Errorf("expected EOF: %d, got %v", d.expectedEOF, err)
			} else if d.expectedEOF != i && err == io.EOF {
				t.Errorf("expected no EOF, got %v", err)
			} else if err != nil && err != io.EOF {
				t.Errorf("unexpected error: %v", err)
			}

			if d.expectedEOF == i {
				break
			}
			i += 1
		}
	}
}
