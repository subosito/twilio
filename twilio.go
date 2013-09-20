package twilio

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version   = "0.1.0"
	userAgent = "subosito/twilio//" + version

	apiBaseURL = "https://api.twilio.com"
	apiVersion = "2010-04-01"
	apiFormat  = "json"
)

type Client struct {
	// HTTP client used to communicate with API
	client *http.Client

	// User agent used when communicating with Twilio API
	UserAgent string

	// Base URL for API requests
	BaseURL *url.URL

	AccountSid string
	AuthToken  string

	// Services used for communicating with different parts of the Twilio API
	Message *MessageService
}

func baseURL() *url.URL {
	baseURL, _ := url.Parse(fmt.Sprintf("%s/%s", apiBaseURL, apiVersion))
	return baseURL
}

func NewClient(accountSid, authToken string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	c := &Client{
		client:     client,
		UserAgent:  userAgent,
		AccountSid: accountSid,
		AuthToken:  authToken,
		BaseURL:    baseURL(),
	}

	c.Message = &MessageService{client: c}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL
	u.Path = u.Path + rel.Path

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	req.SetBasicAuth(c.AccountSid, c.AuthToken)

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")

	return req, nil
}

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
	Start           int    `json:"start"`
	Total           int    `json:"total"`
	NumPages        int    `json:"num_pages"`
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
	End             int    `json:"end"`
	Uri             string `json:"uri"`
	FirstPageUri    string `json:"first_page_uri"`
	LastPageUri     string `json:"last_page_uri"`
	NextPageUri     string `json:"next_page_uri"`
	PreviousPageUri string `json:"previous_page_uri"`
}

type ListParams struct {
	Page     int
	PageSize int
}
