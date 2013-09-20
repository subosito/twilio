package twilio

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

const (
	accountSid = "AC5ef8732a3c49700934481addd5ce1659"
	authToken  = "2ecaf0108548e09a74387cbb28456aa2"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(accountSid, authToken, nil)
	client.BaseURL, _ = url.Parse(server.URL)
}

func teardown() {
	server.Close()
}

func encodeAuth() string {
	s := accountSid + ":" + authToken
	return ("Basic " + base64.StdEncoding.EncodeToString([]byte(s)))
}

func parseTimestamp(s string) Timestamp {
	tm, _ := time.Parse(time.RFC1123Z, s)
	return Timestamp{Time: tm}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if want != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(accountSid, authToken, nil)
	baseURL := "https://api.twilio.com"

	if c.BaseURL.String() != baseURL {
		t.Errorf("NewClient BaseURL = %q, want %q", c.BaseURL.String(), baseURL)
	}

	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %q, want %q", c.UserAgent, userAgent)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(accountSid, authToken, nil)

	inURL := "/foo"
	outURL := c.BaseURL.String() + "/foo"

	req, _ := c.NewRequest("GET", inURL, nil)

	// test that URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%q) URL = %q, want %q", inURL, req.URL, outURL)
	}

	// test that default user-agent is attached to the Request
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %q, want %q", userAgent, c.UserAgent)
	}

	// Test that basic authentication is attached to the Request
	authHash := encodeAuth()
	authHeader := req.Header.Get("Authorization")
	if authHeader != authHash {
		t.Errorf("NewRequest() Authorization = %q, want %q", authHeader, authHash)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(accountSid, authToken, nil)

	_, err := c.NewRequest("GET", ":", nil)

	if err == nil {
		t.Error("Expected error to be returned")
	}

	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		Bar string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		fmt.Fprint(w, `{"Bar":"bar"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(foo)
	client.Do(req, body)

	want := &foo{"bar"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 errror.")
	}
}
