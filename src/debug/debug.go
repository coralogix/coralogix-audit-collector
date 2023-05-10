package debug

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type DryRunCollector struct {
	log []map[string]interface{}
}

func NewDryRunCollector() *DryRunCollector {
	return &DryRunCollector{
		log: []map[string]interface{}{},
	}
}

func (c *DryRunCollector) Collect(v map[string]interface{}) {
	c.log = append(c.log, v)
}

func (c *DryRunCollector) Send() {
	for _, v := range c.log {
		j, _ := json.Marshal(v)
		logrus.Debug(string(j))
	}
}
