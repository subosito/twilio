package twilio

import (
	"net/http"
	"testing"
)

func TestTwilioInvalidAccountForSendingSMS(t *testing.T) {
	w := NewTwilio(accountSid, "abc")
	_, err := w.SimpleSendSMS(from, to, body)
	e := err.(*Exception)

	if e.Status != http.StatusUnauthorized {
		t.Errorf("e.Status: %d != %d", e.Status, http.StatusUnauthorized)
	}

	if code := 20003; e.Code != code {
		t.Errorf("e.Code: %d != %d", e.Code, code)
	}
}

func TestTwilioUnableSendSMS(t *testing.T) {
	w := NewTwilio(accountSid, authToken)
	_, err := w.SimpleSendSMS(from, "abc", body)
	e := err.(*Exception)

	if e.Status != http.StatusBadRequest {
		t.Errorf("e.Status: %d != %d", e.Status, http.StatusBadRequest)
	}

	if code := 21211; e.Code != code {
		t.Errorf("e.Code: %d != %d", e.Code, code)
	}
}

func TestTwilioSendSMS(t *testing.T) {
	w := NewTwilio(accountSid, authToken)
	s, _ := w.SendSMS(from, to, body, SMSParams{})

	if s.AccountSid != accountSid {
		t.Errorf("s.AccountSid: %s != %s", s.AccountSid, accountSid)
	}

	if s.ApiVersion != apiVersion {
		t.Errorf("s.ApiVersion: %s != %s", s.ApiVersion, apiVersion)
	}
}

func TestTwilioGetSMS(t *testing.T) {
	w := NewTwilio(accountSid, authToken)
	w.Transport = &RecordingTransport{
		Transport: http.DefaultTransport,
		Status:    200,
		Body:      response["SMS"],
	}

	r, _ := w.GetSMS("SM800f449d0399ed014aae2bcc0cc2f2ec")

	if r.AccountSid != accountSid {
		t.Errorf("s.AccountSid: %s != %s", r.AccountSid, accountSid)
	}

	if r.ApiVersion != apiVersion {
		t.Errorf("s.ApiVersion: %s != %s", r.ApiVersion, apiVersion)
	}
}

func TestTwilioListSMS(t *testing.T) {
	w := NewTwilio(accountSid, authToken)
	w.Transport = &RecordingTransport{
		Transport: http.DefaultTransport,
		Status:    200,
		Body:      response["SMSList"],
	}

	r, _ := w.ListSMS(map[string]string{})

	if start := 0; r.Start != start {
		t.Errorf("r.Start: %s != %s", r.Start, start)
	}
}
