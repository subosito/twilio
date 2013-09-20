package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestMessageService_Send(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("%s.%s", client.Message.endpoint(), apiFormat)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		fmt.Fprint(w, `{"sid": "abcdef", "num_media": "1", "price": "0.74", "date_sent": null}`)
	})

	params := MessageParams{
		Body:     "Jenny please?! I love you <3",
		MediaUrl: "http://www.example.com/hearts.png",
	}

	m, _, err := client.Message.Send("+14158141829", "+15558675309", params)

	if err != nil {
		t.Errorf("Message.Send returned error: %q", err)
	}

	want := &Message{Sid: "abcdef", NumMedia: 1, Price: 0.74, DateSent: Timestamp(time.Time{})}

	if !reflect.DeepEqual(m, want) {
		t.Errorf("Message.SendSMS returned %+v, want %+v", m, want)
	}
}

func TestMessageService_SendSMS(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("%s.%s", client.Message.endpoint(), apiFormat)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		fmt.Fprint(w, `{"sid": "abcdef", "num_media": "0", "price": "0.74", "date_created": "Wed, 18 Aug 2010 20:01:40 +0000"}`)
	})

	m, _, err := client.Message.SendSMS("+1234567", "+7654321", "Hello!")

	if err != nil {
		t.Errorf("Message.SendSMS returned error: %v", err)
	}

	tm, _ := time.Parse(time.RFC1123Z, "Wed, 18 Aug 2010 20:01:40 +0000")
	want := &Message{Sid: "abcdef", NumMedia: 0, Price: 0.74, DateCreated: Timestamp(tm)}

	if !reflect.DeepEqual(m, want) {
		t.Errorf("Message.SendSMS returned %+v, want %+v", m, want)
	}
}

func TestMessageService_Send_incompleteParams(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("%s.%s", client.Message.endpoint(), apiFormat)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": 400, "message": "A 'From' phone number is required.", "code": 21603}`)
	})

	_, r, err := client.Message.Send("", "+15558675309", MessageParams{Body: "Hello"})

	if r.StatusCode != http.StatusBadRequest {
		t.Errorf("Send() status code = %d, want %d", r.StatusCode, http.StatusBadRequest)
	}

	ex, _ := err.(*Exception)
	want := &Exception{Status: 400, Code: 21603, Message: "A 'From' phone number is required."}

	if !reflect.DeepEqual(ex, want) {
		t.Errorf("Message.SendSMS returned %+v, want %+v", ex, want)
	}
}

func TestMessageService_Get(t *testing.T) {
	setup()
	defer teardown()

	sid := "MM90c6fc909d8504d45ecdb3a3d5b3556e"
	endpoint := fmt.Sprintf("%s/%s.%s", client.Message.endpoint(), sid, apiFormat)

	output := `
{
	"account_sid": "AC5ef8732a3c49700934481addd5ce1659",
	"api_version": "2010-04-01",
	"body": "Jenny please?! I love you <3",
	"num_segments": "1",
	"num_media": "1",
	"date_created": "Wed, 18 Aug 2010 20:01:40 +0000",
	"date_sent": null,
	"date_updated": "Wed, 18 Aug 2010 20:01:40 +0000",
	"direction": "outbound-api",
	"from": "+14158141829",
	"price": null,
	"sid": "MM90c6fc909d8504d45ecdb3a3d5b3556e",
	"status": "queued",
	"to": "+15558675309",
	"uri": "/2010-04-01/Accounts/AC5ef8732a3c49700934481addd5ce1659/Messages/MM90c6fc909d8504d45ecdb3a3d5b3556e.json"
}`

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		fmt.Fprint(w, output)
	})

	m, _, err := client.Message.Get(sid)

	if err != nil {
		t.Errorf("Message.SendSMS returned error: %v", err)
	}

	tm, _ := time.Parse(time.RFC1123Z, "Wed, 18 Aug 2010 20:01:40 +0000")
	want := &Message{
		AccountSid:  "AC5ef8732a3c49700934481addd5ce1659",
		ApiVersion:  "2010-04-01",
		Body:        "Jenny please?! I love you <3",
		NumSegments: 1,
		NumMedia:    1,
		DateCreated: Timestamp(tm),
		DateSent:    Timestamp(time.Time{}),
		DateUpdated: Timestamp(tm),
		Direction:   "outbound-api",
		From:        "+14158141829",
		Price:       Price(0),
		Sid:         "MM90c6fc909d8504d45ecdb3a3d5b3556e",
		Status:      "queued",
		To:          "+15558675309",
		Uri:         "/2010-04-01/Accounts/AC5ef8732a3c49700934481addd5ce1659/Messages/MM90c6fc909d8504d45ecdb3a3d5b3556e.json",
	}

	if !reflect.DeepEqual(m, want) {
		t.Errorf("Message.SendSMS returned %+v, want %+v", m, want)
	}
}

func TestMessageService_List(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("%s.%s", client.Message.endpoint(), apiFormat)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		output := `{
			"page": 0,
			"page_size": 50,
			"end": 49,
			"uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json",
			"previous_page_uri": null,
			"messages": [{ "sid": "MM90c6fc909d8504d45ecdb3a3d5b3556e" }]
		}`

		fmt.Fprint(w, output)
	})

	ml, r, err := client.Message.List(MessageListParams{})

	if err != nil {
		t.Error("Get() err expected to be nil")
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("Send() status code = %d, want %d", r.StatusCode, http.StatusOK)
	}

	want := &MessageList{
		Pagination: Pagination{
			Page:            0,
			PageSize:        50,
			End:             49,
			Uri:             "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json",
			PreviousPageUri: "",
		},
		Messages: []Message{Message{Sid: "MM90c6fc909d8504d45ecdb3a3d5b3556e"}},
	}

	if !reflect.DeepEqual(ml, want) {
		t.Errorf("Message.List returned %+v, want %+v", ml, want)
	}
}
