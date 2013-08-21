package twilio

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Parameter interface {
	urlValues() url.Values
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

func unquote(b []byte) string {
	switch b[0] {
	case '"':
		return string(b[1 : len(b)-1])
	default:
		return string(b)
	}
}

type Price float32

func (p *Price) UnmarshalJSON(b []byte) error {
	return json.Unmarshal([]byte(unquote(b)), (*float32)(p))
}

type Timestamp time.Time

func (m *Timestamp) UnmarshalJSON(b []byte) error {
	s := unquote(b)

	if s == "null" {
		*m = Timestamp(time.Time{})
		return nil
	}

	t, err := time.Parse(time.RFC1123Z, s)
	if err != nil {
		return err
	}

	*m = Timestamp(t)
	return nil
}
