package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMessageService_Send(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("%s.%s", client.Message.endpoint(), apiFormat)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		fmt.Fprint(w, output["message"])
	})

	params := MessageParams{
		Body:     "Jenny please?! I love you <3",
		MediaUrl: "http://www.example.com/hearts.png",
	}

	m, _, err := client.Message.Send("+14158141829", "+15558675309", params)

	if err != nil {
		t.Errorf("Message.Send returned error: %q", err)
	}

	if m.AccountSid != accountSid {
		t.Errorf("Message.AccountdSid = %q, want %q", m.AccountSid, accountSid)
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

		fmt.Fprint(w, `{"sid": "abcdef", "num_media": "0"}`)
	})

	m, _, err := client.Message.SendSMS("+1234567", "+7654321", "Hello!")

	if err != nil {
		t.Errorf("Message.SendSMS returned error: %v", err)
	}

	want := &Message{Sid: "abcdef", NumMedia: 0}
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
		fmt.Fprint(w, output["message_error"])
	})

	params := MessageParams{
		Body: "Jenny please?! I love you <3",
	}

	_, r, err := client.Message.Send("", "+15558675309", params)

	if r.StatusCode != http.StatusBadRequest {
		t.Errorf("Send() status code = %d, want %d", r.StatusCode, http.StatusBadRequest)
	}

	ex, _ := err.(*Exception)

	if ex.Status != http.StatusBadRequest {
		t.Errorf("Exception.Status = %d, want %d", ex.Status, http.StatusBadRequest)
	}
}

func TestMessageService_Get(t *testing.T) {
	setup()
	defer teardown()

	sid := "MM90c6fc909d8504d45ecdb3a3d5b3556e"
	endpoint := fmt.Sprintf("%s/%s.%s", client.Message.endpoint(), sid, apiFormat)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		fmt.Fprint(w, output["message"])
	})

	m, r, err := client.Message.Get(sid)

	if err != nil {
		t.Error("Get() err expected to be nil")
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("Send() status code = %d, want %d", r.StatusCode, http.StatusOK)
	}

	if m.Sid != sid {
		t.Errorf("Send() sid = %q, want %q", m.Sid, sid)
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

		fmt.Fprint(w, output["message_list"])
	})

	ml, r, err := client.Message.List(MessageListParams{})

	if err != nil {
		t.Error("Get() err expected to be nil")
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("Send() status code = %d, want %d", r.StatusCode, http.StatusOK)
	}

	if total := 261; ml.Total != total {
		t.Errorf("MessageList.Total = %d, want %d", ml.Total, total)
	}
}
