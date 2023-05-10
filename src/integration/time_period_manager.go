package integration

import (
	"time"
)

type TimePeriodManager struct {
	StartFrom     time.Time
	ToEnd         time.Time
	DiffInMinutes time.Duration
}

func NewTimePeriodManagerFromNow(diffInMinutes int) *TimePeriodManager {
	now := time.Now()
	return &TimePeriodManager{
		StartFrom:     now,
		ToEnd:         now,
		DiffInMinutes: time.Duration(diffInMinutes),
	}
}

func (t *TimePeriodManager) From() time.Time {
	return t.StartFrom.Add(-t.DiffInMinutes * time.Minute)
}

func (t *TimePeriodManager) To() time.Time {
	return t.ToEnd
}

func NewTimePeriodManager(startDate time.Time, diffInMinutes int) *TimePeriodManager {
	return &TimePeriodManager{
		StartFrom:     startDate,
		ToEnd:         startDate,
		DiffInMinutes: time.Duration(diffInMinutes),
	}
}
