package twilio

import (
	"net/http"
	"testing"
)

const (
	accountSid = "AC65c2ee6fddb57b5ad6818ddb850c20b9"
	authToken  = "2ecaf0108548e09a74387cbb28456aa2"
)

var sample = map[string]string{
	"from": "+15005550006",
	"to":   "+62821234567",
	"body": "Hello Go!",
}

func TestTwilioInvalidAccountForSendingSMS(t *testing.T) {
	w := NewTwilio(accountSid, "abc")
	_, err := w.SimpleSendSMS(sample["from"], sample["to"], sample["body"])
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
	_, err := w.SimpleSendSMS(sample["from"], "abc", sample["body"])
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
	s, _ := w.SendSMS(sample["from"], sample["to"], sample["body"], SMSParams{})

	if s.AccountSid != accountSid {
		t.Errorf("s.AccountSid: %s != %s", s.AccountSid, accountSid)
	}

	if s.ApiVersion != apiVersion {
		t.Errorf("s.ApiVersion: %s != %s", s.ApiVersion, apiVersion)
	}
}
