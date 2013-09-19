package twilio

import (
	"fmt"
	"net/http"
	"testing"
)

var message string = `
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

var messageList string = `
{
	"start": 0,
	"total": 261,
	"num_pages": 6,
	"page": 0,
	"page_size": 50,
	"end": 49,
	"uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json",
	"first_page_uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json?Page=0&PageSize=50",
	"last_page_uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json?Page=5&PageSize=50",
	"next_page_uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages.json?Page=1&PageSize=50",
	"previous_page_uri": null,
	"messages": [
		{
			"account_sid": "ACc51860f991f74032b73fdc58841d39fa",
			"api_version": "2010-04-01",
			"body": "Hey Jenny why aren't you returning my calls?",
			"num_segments": "1",
			"num_media": "0",
			"date_created": "Mon, 16 Aug 2010 03:45:01 +0000",
			"date_sent": "Mon, 16 Aug 2010 03:45:03 +0000",
			"date_updated": "Mon, 16 Aug 2010 03:45:03 +0000",
			"direction": "outbound-api",
			"from": "+14158141829",
			"price": "-0.02000",
			"sid": "SM800f449d0399ed014aae2bcc0cc2f2ec",
			"status": "sent",
			"to": "+15558675309",
			"uri": "/2010-04-01/Accounts/ACc51860f991f74032b73fdc58841d39fa/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec.json"
		}
	]
}`

func TestMessageService_Send(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("%s.%s", client.Message.endpoint(), apiFormat)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		fmt.Fprint(w, message)
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

func TestMessageService_Send_incompleteParams(t *testing.T) {
	setup()
	defer teardown()

	response := `{
					"status": 400,
					"message": "A 'From' phone number is required.",
					"code": 21603,
					"more_info": "http:\/\/www.twilio.com\/docs\/errors\/21603"
				}`

	endpoint := fmt.Sprintf("%s.%s", client.Message.endpoint(), apiFormat)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, response)
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

		fmt.Fprint(w, message)
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

		fmt.Fprint(w, messageList)
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
