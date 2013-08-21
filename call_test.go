package twilio

import (
	"testing"
)

func TestTwilioMakeCall(t *testing.T) {
	w := NewTwilio(accountSid, authToken)
	r, _ := w.MakeCall(from, to, CallParams{Url: callbackUrl})

	if r.AccountSid != accountSid {
		t.Errorf("s.AccountSid: %s != %s", r.AccountSid, accountSid)
	}

	if r.ApiVersion != apiVersion {
		t.Errorf("s.ApiVersion: %s != %s", r.ApiVersion, apiVersion)
	}
}
