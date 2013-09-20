package twilio

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type MessageService struct {
	client *Client
}

type Message struct {
	AccountSid  string    `json:"account_sid"`
	ApiVersion  string    `json:"api_version"`
	Body        string    `json:"body"`
	NumSegments int       `json:"num_segments,string"`
	NumMedia    int       `json:"num_media,string"`
	DateCreated Timestamp `json:"date_created,omitempty"`
	DateSent    Timestamp `json:"date_sent,omitempty"`
	DateUpdated Timestamp `json:"date_updated,omitempty"`
	Direction   string    `json:"direction"`
	From        string    `json:"from"`
	Price       Price     `json:"price,omitempty"`
	Sid         string    `json:"sid"`
	Status      string    `json:"status"`
	To          string    `json:"to"`
	Uri         string    `json:"uri"`
}

type MessageParams struct {
	// The text of the message you want to send, limited to 1600 characters.
	Body string

	// The URL of the media you wish to send out with the message. Currently support: gif, png, and jpeg.
	MediaUrl string

	StatusCallback string
	ApplicationSid string
}

func (p MessageParams) validates() error {
	if (p.Body == "") && (p.MediaUrl == "") {
		return errors.New(`One of the "Body" or "MediaUrl" is required.`)
	}

	return nil
}

func (p MessageParams) values() url.Values {
	v := make(url.Values)

	fields := reflect.TypeOf(&p).Elem()
	values := reflect.ValueOf(p)
	num := fields.NumField()

	for i := 0; i < num; i++ {
		key := fields.FieldByIndex([]int{i}).Name
		val := values.Field(i).String()

		if val != "" {
			v.Set(key, val)
		}
	}

	return v
}

func (s *MessageService) endpoint() string {
	return fmt.Sprintf("/Accounts/%s/Messages", s.client.AccountSid)
}

// Shortcut for sending SMS with no optional parameters support.
func (s *MessageService) SendSMS(from, to, body string) (*Message, *Response, error) {
	return s.Send(from, to, MessageParams{Body: body})
}

// Send Message with options. It's support required and optional parameters.
//
// One of these parameter is required:
//
//	Body     : The text of the message you want to send
//	MediaUrl : The URL of the media you wish to send out with the message. Currently support: gif, png, and jpeg.
//
// Optional parameters:
//
//	StatusCallback : A URL that Twilio will POST to when your message is processed.
//	ApplicationSid : Twilio will POST `MessageSid` as well as other statuses to the URL in the `MessageStatusCallback` property of this application
func (s *MessageService) Send(from, to string, params MessageParams) (*Message, *Response, error) {
	u := fmt.Sprintf("%s.%s", s.endpoint(), apiFormat)

	err := params.validates()
	if err != nil {
		return nil, nil, err
	}

	v := params.values()
	v.Set("From", from)
	v.Set("To", to)

	req, err := s.client.NewRequest("POST", u, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, nil, err
	}

	m := new(Message)
	resp, err := s.client.Do(req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, err
}

func (s *MessageService) Get(sid string) (*Message, *Response, error) {
	u := fmt.Sprintf("%s/%s.%s", s.endpoint(), sid, apiFormat)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	m := new(Message)
	resp, err := s.client.Do(req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, err
}

type MessageList struct {
	Pagination
	Messages []Message
}

type MessageListParams struct {
	To       string
	From     string
	DateSent string
}

func (p MessageListParams) values() url.Values {
	v := make(url.Values)

	fields := reflect.TypeOf(&p).Elem()
	values := reflect.ValueOf(p)
	num := fields.NumField()

	for i := 0; i < num; i++ {
		key := fields.FieldByIndex([]int{i}).Name
		val := values.Field(i).String()

		v.Set(key, val)
	}

	return v
}

func (s *MessageService) List(params MessageListParams) (*MessageList, *Response, error) {
	u := fmt.Sprintf("%s.%s", s.endpoint(), apiFormat)

	v := params.values()

	req, err := s.client.NewRequest("GET", u, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, nil, err
	}

	ml := new(MessageList)
	resp, err := s.client.Do(req, ml)
	if err != nil {
		return nil, resp, err
	}

	return ml, resp, err
}
