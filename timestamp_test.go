package twilio

import (
	"testing"
	"time"
)

func TestTimestamp_IsZero(t *testing.T) {
	m := &Timestamp{}
	if !m.IsZero() {
		t.Error("Timestamp.IsZero() should be true")
	}

	p := &Timestamp{Time: time.Date(2013, 4, 87, 21, 10, 56, 0, time.UTC)}
	if p.IsZero() {
		t.Error("Timestamp.IsZero() should be false")
	}
}
