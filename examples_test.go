package twilio_test

import (
	"fmt"
	"github.com/subosito/twilio"
)

// This example shows the usage of twilio package. You can get your AccountSid and AuthToken on Account Dashboard page.
func Example() {
	// Prepare credentials
	AccountSid := "ac650108548e09aC2eed18ddb850c20b9"
	AuthToken := "2ecaf74387cbb28456aad6fb57b5ad682"

	// Initialize twilio client
	t := twilio.NewTwilio(AccountSid, AuthToken)

	// You can set custom Transport, eg: when you're using `appengine/urlfetch` on Google's appengine.
	// c := appengine.NewContext(r) // r is a *http.Request
	// t.Transport = urlfetch.Transport{Context: c}

	// Send SMS
	params := map[string]string{"StatusCallback": "http://example.com/"}
	s, err := t.SendSMS("+15005550006", "+62821234567", "Hello Go!", params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", s)
	return
}
