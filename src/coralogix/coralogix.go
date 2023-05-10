package coralogix

import (
	coralogixsdk "github.com/coralogix/go-coralogix-sdk"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	coralogixPrivateKey = os.Getenv("CORALOGIX_PRIVATE_KEY")
	coralogixAppname    = os.Getenv("CORALOGIX_APP_NAME")
	dryRun              = os.Getenv("DRY_RUN") == "true"
)

func NewCollector(coralogixSubsystemname string) *Collector {
	coralogixsdk.SetDebug(true)
	logger := coralogixsdk.NewCoralogixLogger(
		coralogixPrivateKey,
		coralogixAppname,
		coralogixSubsystemname,
	)
	logger.Destroy()

	return &Collector{
		logger: logger,
	}
}

type Collector struct {
	logger *coralogixsdk.CoralogixLogger
}

func (c *Collector) Collect(v map[string]interface{}) {
	c.logger.Info(v)
}

func (c *Collector) Send() {
	if dryRun {
		return
	}

	manager := c.logger.LoggerManager
	buffer := &manager.LogsBuffer

	for {
		logToSend, err := c.sendBulk(buffer)
		if err != nil && err != io.EOF {
			logrus.Warnf("Error while sending bulk: %s", err)
			break
		} else if err == io.EOF {
			break
		}

		LogsBulk := coralogixsdk.NewBulk(manager.Credentials)
		for _, Record := range logToSend {
			LogsBulk.AddRecord(Record)
		}

		logrus.Debugf("Sending batch with %d logs", len(logToSend))
		coralogixsdk.SendRequest(LogsBulk)
	}

	c.logger.Destroy()
}

func (c *Collector) sendBulk(buffer *coralogixsdk.LogBuffer) ([]coralogixsdk.Log, error) {
	bufferLen := buffer.Len()
	if bufferLen < 1 {
		return nil, io.EOF
	}

	var logsToSend []coralogixsdk.Log
	if buffer.Size() > coralogixsdk.MaxLogChunkSize {
		for buffer.Size() > coralogixsdk.MaxLogChunkSize {
			bufferLen = bufferLen / 2
			logsToSend = append(logsToSend, buffer.Slice(bufferLen)...)
		}
	} else {
		logsToSend = buffer.Slice(bufferLen)
	}

	return logsToSend, nil
}
