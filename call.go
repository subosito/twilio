package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CallResponse struct {
	Sid             string    `json:"sid"`
	DateCreated     Timestamp `json:"date_created,omitempty"`
	DateUpdated     Timestamp `json:"date_updated,omitempty"`
	ParentCallSid   string    `json:"parent_call_sid"`
	AccountSid      string    `json:"account_sid"`
	To              string    `json:"to"`
	ToFormatted     string    `json:"to_formatted"`
	From            string    `json:"from"`
	FromFormatted   string    `json:"from_formatted"`
	PhoneNumberSid  string    `json:"phone_number_sid"`
	Status          string    `json:"status"`
	StartTime       Timestamp `json:"start_time,omitempty"`
	EndTime         Timestamp `json:"end_time,omitempty"`
	Duration        string    `json:"duration,omitempty"`
	Price           Price     `json:"price,omitempty"`
	Direction       string    `json:"direction"`
	AnsweredBy      string    `json:"answered_by,omitempty"`
	ApiVersion      string    `json:"api_version"`
	ForwardedFrom   string    `json:"forwarded_from,omitempty"`
	CallerName      string    `json:"caller_name,omitempty"`
	Uri             string    `json:"uri"`
	SubresourceUris struct {
		Notifications string `json:"notifications"`
		Recordings    string `json:"recordings"`
	} `json:"subresource_uris"`
}

func (t *Twilio) callEndpoint() string {
	return fmt.Sprintf("%s/Accounts/%s/Calls", t.BaseUrl, t.AccountSid)
}

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

func (p CallParams) urlValues() url.Values {
	uv := url.Values{}

	if p.Url != "" {
		uv.Set("Url", p.Url)
		uv.Set("Method", p.Method)
		uv.Set("FallbackUrl", p.FallbackUrl)
		uv.Set("FallbackMethod", p.FallbackMethod)
		uv.Set("StatusCallback", p.StatusCallback)
		uv.Set("StatusCallbackMethod", p.StatusCallbackMethod)

		p.ApplicationSid = "" // reset
	}

	if p.ApplicationSid != "" {
		uv.Set("ApplicationSid", p.ApplicationSid)
	}

	// set default timeout
	if p.Timeout == 0 {
		p.Timeout = 60
	}

	uv.Set("SendDigits", p.SendDigits)
	uv.Set("IfMachine", p.IfMachine)
	uv.Set("Timeout", strconv.Itoa(p.Timeout))
	uv.Set("Record", fmt.Sprintf("%t", p.Record))
	return uv
}

// Make a voice call. You need to set one of `Url` or `ApplicationSid` parameter on `CallParams`
func (t *Twilio) MakeCall(from, to string, p CallParams) (r *CallResponse, err error) {
	endpoint := fmt.Sprintf("%s.%s", t.callEndpoint(), apiFormat)

	if (p.Url == "") && (p.ApplicationSid == "") {
		err := errors.New("One of the Url or ApplicationSid is required.")
		return nil, err
	}

	params := p.urlValues()
	params.Set("From", from)
	params.Set("To", to)

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

type CallSipParams struct {
	From            string
	SipAuthUsername string
	SipAuthPassword string
}

func (p CallSipParams) urlValues() url.Values {
	uv := url.Values{}
	uv.Set("From", p.From)
	uv.Set("SipAuthUsername", p.SipAuthUsername)
	uv.Set("SipAuthPassword", p.SipAuthPassword)
	return uv
}

func (t *Twilio) MakeSipCall(to string, p CallSipParams) (r *CallResponse, err error) {
	endpoint := fmt.Sprintf("%s.%s", t.callEndpoint(), apiFormat)

	params := p.urlValues()
	params.Set("To", to)

	b, status, err := t.request("POST", endpoint, params)
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

	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return
}
