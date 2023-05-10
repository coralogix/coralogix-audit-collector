package googleworkspace

import (
	reports "google.golang.org/api/admin/reports/v1"
	"testing"
)

func TestCollectEventParameters(t *testing.T) {
	params := []*reports.ActivityEventsParameters{
		{
			Name:  "name",
			Value: "value",
		},
		{
			Name:     "name2",
			IntValue: 1,
		},
		{
			Name:      "name3",
			BoolValue: true,
		},
		{
			Name:          "name4",
			MultiIntValue: []int64{1, 2, 3},
		},
		{
			Name:       "name5",
			MultiValue: []string{"a", "b", "c"},
		},
	}
	event := &reports.ActivityEvents{
		Name:       "name",
		Type:       "type",
		Parameters: params,
	}

	g := GoogleWorkspace{}
	collected := g.collectEventParameters(event)
	collectedEvents := collected["Parameters"].(map[string]interface{})
	if len(collectedEvents) != 5 {
		t.Errorf("Expected 5 collected parameters, got %d, %+v", len(collected), collected)
	}

	if collectedEvents["name"] != "value" {
		t.Errorf("Expected collectedEvents[name] to be 'value', got %s", collectedEvents["name"])
	}

	if collectedEvents["name2"] != "1" {
		t.Errorf("Expected collectedEvents[name2] to be 1, got %s", collectedEvents["name2"])
	}

	if collectedEvents["name3"] != "true" {
		t.Errorf("Expected collectedEvents[name3] to be true, got %s", collectedEvents["name3"])
	}

	if len(collectedEvents["name4"].([]string)) != 3 {
		t.Errorf("Expected collectedEvents[name4] to be [1 2 3], got %+v", collectedEvents["name4"])
	}

	if collectedEvents["name4"].([]string)[0] != "1" {
		t.Errorf("Expected collectedEvents[name4][0] to be '1', got %s", collectedEvents["name4"].([]string)[0])
	}

	if len(collectedEvents["name5"].([]string)) != 3 {
		t.Errorf("Expected collectedEvents[name5] to be [a b c], got %+v", collectedEvents["name5"])
	}

	if collectedEvents["name5"].([]string)[0] != "a" {
		t.Errorf("Expected collectedEvents[name5][0] to be 'a', got %s", collectedEvents["name5"].([]string)[0])
	}
}
