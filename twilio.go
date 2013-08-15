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
}

// Exception holds information about error response returned by Twilio API
type Exception struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"more_info"`
}

// Exception implements Error interface
func (e *Exception) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func NewTwilio(accountSid, authToken string) *Twilio {
	baseUrl := fmt.Sprintf("%s/%s", apiHost, apiVersion)
	return &Twilio{accountSid, authToken, baseUrl}
}

func (t *Twilio) post(u string, v url.Values) (b []byte, status int, err error) {
	res, err := t.request("POST", u, v)

	if err != nil {
		return
	}
	defer res.Body.Close()

	status = res.StatusCode
	b, err = ioutil.ReadAll(res.Body)
	return
}

func (t *Twilio) get(u string, v url.Values) (b []byte, status int, err error) {
	res, err := t.request("GET", u, v)

	if err != nil {
		return
	}
	defer res.Body.Close()

	status = res.StatusCode
	b, err = ioutil.ReadAll(res.Body)
	return
}

func (t *Twilio) request(method string, u string, v url.Values) (*http.Response, error) {
	req, err := http.NewRequest(method, u, strings.NewReader(v.Encode()))

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(t.AccountSid, t.AuthToken)

	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	client := &http.Client{}
	return client.Do(req)
}
