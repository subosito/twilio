package twilio

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiBaseURL = "https://api.twilio.com"
	apiVersion = "2010-04-01"
	apiFormat  = "json"
	version   = "0.1.0"
	userAgent = "subosito/twilio//" + version
)

// A client manages communication with Twilio API
type Client struct {
	// HTTP client used to communicate with API
	client *http.Client

	// User agent used when communicating with Twilio API
	UserAgent string

	// Services used for communicating with different parts of the Twilio API
	Messages *MessageService

	// The Twilio API base URL
	BaseURL *url.URL

	// Credentials which is used for authentication during API request
	AccountSid string
	AuthToken  string
}

// NewClient returns a new Twilio API client. This will load default http.Client if httpClient is nil.
func NewClient(accountSid, authToken string, httpClient *http.Client) *Client {
	if httpClient == nil {
		tr := &http.Transport{
			ResponseHeaderTimeout: time.Duration(3050) * time.Millisecond,
		}

		httpClient = &http.Client{Transport: tr}
	}

	baseURL, _ := url.Parse(apiBaseURL)

	c := &Client{
		client:     httpClient,
		UserAgent:  userAgent,
		AccountSid: accountSid,
		AuthToken:  authToken,
		BaseURL:    baseURL,
	}

	c.Messages = &MessageService{client: c}

	return c
}

// Constructing API endpoint. This will returns an *url.URL. Here's the example:
//
//	c := NewClient("1234567", "token", nil)
//	c.EndPoint("Messages", "abcdef") // "/2010-04-01/Accounts/1234567/Messages/abcdef.json"
//
func (c *Client) EndPoint(parts ...string) (*url.URL, error) {
	up := []string{apiVersion, "Accounts", c.AccountSid}
	up = append(up, parts...)
	u, err := url.Parse(strings.Join(up, "/"))
	if err != nil {
		return nil, err
	}

	u.Path = fmt.Sprintf("/%s.%s", u.Path, apiFormat)

	return u, nil
}

func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	ul, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(ul)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	req.SetBasicAuth(c.AccountSid, c.AuthToken)

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")

	return req, nil
}

// Wraps http.Response. So we can add more functionalities later.
type Response struct {
	*http.Response
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &Response{resp}

	err = checkResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return response, err
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

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	exception := new(Exception)
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &exception)
	}

	return exception
}

type Pagination struct {
	Page            int    `json:"page"`
	NumPages        int    `json:"num_pages"`
	PageSize        int    `json:"page_size"`
	Total           int    `json:"total"`
	Start           int    `json:"start"`
	End             int    `json:"end"`
	Uri             string `json:"uri"`
	FirstPageUri    string `json:"first_page_uri"`
	PreviousPageUri string `json:"previous_page_uri"`
	NextPageUri     string `json:"next_page_uri"`
	LastPageUri     string `json:"last_page_uri"`
}

type ListParams struct {
	PageSize int
}
