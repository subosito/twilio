package twilio

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiHost    = "https://api.twilio.com"
	apiVersion = "2010-04-01"
	apiFormat  = "json"
)

type Twilio struct {
	AccountSid string
	AuthToken  string
	BaseUrl    string

	// Transport is the HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil
	Transport http.RoundTripper
}

func NewTwilio(accountSid, authToken string) *Twilio {
	baseUrl := fmt.Sprintf("%s/%s", apiHost, apiVersion)
	return &Twilio{accountSid, authToken, baseUrl, nil}
}

func (t *Twilio) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return http.DefaultTransport
}

func (t *Twilio) request(method string, u string, v url.Values) (b []byte, status int, err error) {
	// remove empty value
	for key, val := range v {
		if strings.Join(val, "") == "" {
			v.Del(key)
		}
	}

	req, err := http.NewRequest(method, u, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, 0, err
	}

	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	req.SetBasicAuth(t.AccountSid, t.AuthToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")

	client := &http.Client{Transport: t.transport()}

	res, err := client.Do(req)
	if err != nil {
		return nil, res.StatusCode, err
	}
	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)
	if err == nil {
		return b, res.StatusCode, nil
	}

	return nil, res.StatusCode, err
}
