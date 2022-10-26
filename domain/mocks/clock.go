package mocks

import "time"

type ClockMock struct {
	Time *time.Time
}

func (c ClockMock) Now() time.Time {
	if c.Time == nil {
		return time.Now()
	}
	return *c.Time
}
