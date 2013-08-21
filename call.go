package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CallParams struct {
	// required, choose one of these
	Url            string
	ApplicationSid string

	Method               string
	FallbackUrl          string
	FallbackMethod       string
	StatusCallback       string
	StatusCallbackMethod string
	SendDigits           string
	IfMachine            string // Continue or Hangup
	Timeout              int
	Record               bool
}

type CallResponse struct {
	CommonResponse
	ParentCallSid  string    `json:"parent_call_sid"`
	AccountSid     string    `json:"account_sid"`
	To             string    `json:"to"`
	ToFormatted    string    `json:"to_formatted"`
	From           string    `json:"from"`
	FromFormatted  string    `json:"from_formatted"`
	PhoneNumberSid string    `json:"phone_number_sid"`
	StartTime      Timestamp `json:"start_time,omitempty"`
	EndTime        Timestamp `json:"end_time,omitempty"`
	Duration       string    `json:"duration,omitempty"`
	Price          Price     `json:"price,omitempty"`
	Direction      string    `json:"direction"`
	AnsweredBy     string    `json:"answered_by,omitempty"`
	ApiVersion     string    `json:"api_version"`
	ForwardedFrom  string    `json:"forwarded_from,omitempty"`
	CallerName     string    `json:"caller_name,omitempty"`
	SubresourceUris struct {
		Notifications string `json:"notifications"`
		Recordings    string `json:"recordings"`
	} `json:"subresource_uris"`
}

func (t *Twilio) callEndpoint() string {
	return fmt.Sprintf("%s/Accounts/%s/Calls", t.BaseUrl, t.AccountSid)
}

// Make a voice call. You need to set one of `Url` or `ApplicationSid` parameter on `CallParams`
func (t *Twilio) MakeCall(from, to string, p CallParams) (r *CallResponse, err error) {
	endpoint := fmt.Sprintf("%s.%s", t.callEndpoint(), apiFormat)
	params := url.Values{}
	params.Set("From", from)
	params.Set("To", to)

	if (p.Url == "") && (p.ApplicationSid == "") {
		return nil, errors.New("One of the Url or ApplicationSid is required.")
	}

	if p.Url != "" {
		params.Set("Url", p.Url)
		params.Set("Method", p.Method)
		params.Set("FallbackUrl", p.FallbackUrl)
		params.Set("FallbackMethod", p.FallbackMethod)
		params.Set("StatusCallback", p.StatusCallback)
		params.Set("StatusCallbackMethod", p.StatusCallbackMethod)

		p.ApplicationSid = "" // reset
	}

	if p.ApplicationSid != "" {
		params.Set("ApplicationSid", p.ApplicationSid)
	}

	// set default timeout
	if p.Timeout == 0 {
		p.Timeout = 60
	}

	params.Set("SendDigits", p.SendDigits)
	params.Set("IfMachine", p.IfMachine)
	params.Set("Timeout", strconv.Itoa(p.Timeout))
	params.Set("Record", fmt.Sprintf("%t", p.Record))

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

	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return
}
