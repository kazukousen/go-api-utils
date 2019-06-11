package utils

import "time"

// Now returns a now time (or fakeTime value).
func Now() time.Time {
	if !fakeTime.IsZero() {
		return fakeTime
	}
	return time.Now()
}

var fakeTime time.Time

// SetFakeTime updates the value returned by Now().
// DO NOT USING PRODUCTION.
func SetFakeTime(t time.Time) {
	fakeTime = t
}
