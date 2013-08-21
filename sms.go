package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SMSResponse struct {
	CommonResponse
	AccountSid string    `json:"account_sid"`
	ApiVersion string    `json:"api_version"`
	Body       string    `json:"body"`
	DateSent   Timestamp `json:"date_sent,omitempty"`
	Direction  string    `json:"direction"`
	From       string    `json:"from"`
	Price      Price     `json:"price,omitempty"`
	To         string    `json:"to"`
}

type SMSListResponse struct {
	Pagination
	SMSMessages []SMSResponse
}

type SMSParams struct {
	StatusCallback string
	ApplicationSid string
}

func (t *Twilio) smsEndpoint() string {
	return fmt.Sprintf("%s/Accounts/%s/SMS/Messages", t.BaseUrl, t.AccountSid)
}

// Simple version of Send SMS with no optional parameters support.
func (t *Twilio) SimpleSendSMS(from, to, body string) (*SMSResponse, error) {
	return t.SendSMS(from, to, body, SMSParams{})
}

// Send SMS with more verbose options. It's support optional parameters.
//	StatusCallback : A URL that Twilio will POST to when your message is processed.
//	ApplicationSid : Twilio will POST `SMSSid` as well as other statuses to the URL in the `SMSStatusCallback` property of this application
func (t *Twilio) SendSMS(from, to, body string, p SMSParams) (s *SMSResponse, err error) {
	endpoint := fmt.Sprintf("%s.%s", t.smsEndpoint(), apiFormat)
	params := url.Values{}
	params.Set("From", from)
	params.Set("To", to)
	params.Set("Body", body)

	if p.StatusCallback != "" {
		params.Set("StatusCallback", p.StatusCallback)
	}

	if p.ApplicationSid != "" {
		params.Set("ApplicationSid", p.ApplicationSid)
	}

	b, status, err := t.request("POST", endpoint, params)
	if err != nil {
		return
	}

	if status != http.StatusCreated {
		e := new(Exception)
		err = json.Unmarshal(b, &e)
		if err != nil {
			return
		}

		return nil, e
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}

	return
}

func (t *Twilio) GetSMS(sid string) (s *SMSResponse, err error) {
	endpoint := fmt.Sprintf("%s/%s.%s", t.smsEndpoint(), sid, apiFormat)

	b, status, err := t.request("GET", endpoint, url.Values{})
	if err != nil {
		return
	}

	if status != http.StatusOK {
		e := new(Exception)
		err = json.Unmarshal(b, &e)
		if err != nil {
			return
		}

		return nil, e
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}

	return
}

// Returns a list of SMS messages associates with your account. It's support list filters via `map[string]string`:
//	"To" : Only show SMS messages to this phone number
//	"From" : Only show SMS messages from this phone number
//	"DateSent" : Only show SMS messages sent on this date (in GMT format), given as `YYYY-MM-DD`.
func (t *Twilio) ListSMS(filters map[string]string) (sl *SMSListResponse, err error) {
	endpoint := fmt.Sprintf("%s.%s", t.smsEndpoint(), apiFormat)
	params := url.Values{}

	for key, value := range filters {
		params.Set(key, value)
	}

	b, status, err := t.request("GET", endpoint, params)
	if err != nil {
		return
	}

	if status != http.StatusOK {
		e := new(Exception)
		err = json.Unmarshal(b, &e)
		if err != nil {
			return
		}

		return nil, e
	}

	err = json.Unmarshal(b, &sl)
	if err != nil {
		return nil, err
	}

	return
}
