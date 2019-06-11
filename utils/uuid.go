package utils

import "github.com/google/uuid"

// NewUUID returns random UUID (or fakeUUID value).
func NewUUID() string {
	if fakeUUID != "" {
		return fakeUUID
	}
	return uuid.New().String()
}

var fakeUUID string

// SetFakeUUID updates the value returned by NewUUID().
// DO NOT USING PRODUCTION.
func SetFakeUUID(fake string) {
	fakeUUID = fake
}
