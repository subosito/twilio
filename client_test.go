package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(accountSid, authToken, nil)

	if c.BaseURL.String() != apiBaseURL {
		t.Errorf("NewClient BaseURL = %q, want %q", c.BaseURL.String(), apiBaseURL)
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

func TestDo_redirectLoop(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Error("Expected error to be returned.")
	}

	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected a URL error; got %#v.", err)
	}
}

func TestEndPoint(t *testing.T) {
	setup()
	defer teardown()

	u := client.EndPoint("Hello", "123")
	want, _ := url.Parse("/2010-04-01/Accounts/AC5ef87/Hello/123.json")

	if !reflect.DeepEqual(u, want) {
		t.Errorf("EndPoint returned %+v, want %+v", u, want)
	}
}
