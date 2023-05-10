package integration

import (
	"testing"
	"time"
)

func TestTimePeriodManager(t *testing.T) {
	now := time.Now()
	tm := &TimePeriodManager{
		StartFrom:     now,
		ToEnd:         now,
		DiffInMinutes: 1,
	}

	if now.Minute()-1 != tm.From().Minute() {
		t.Errorf("Expected %d, got %d", now.Minute()-1, tm.From().Minute())
	}

	if now.Minute() != tm.To().Minute() {
		t.Errorf("Expected %d, got %d", now.Minute(), tm.To().Minute())
	}
}
