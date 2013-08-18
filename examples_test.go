package twilio_test

import (
	"fmt"
	"github.com/subosito/twilio"
)

// This example shows the usage of twilio package. You can get your AccountSid and AuthToken on Account Dashboard page.
func Example() {
	// Common stuffs
	AccountSid := "ac650108548e09aC2eed18ddb850c20b9"
	AuthToken := "2ecaf74387cbb28456aad6fb57b5ad682"
	from := "+15005550006"
	to := "+62801234567"
	callbackUrl := "http://subosito.com/"

	// Initialize twilio client
	t := twilio.NewTwilio(AccountSid, AuthToken)

	// You can set custom Transport, eg: when you're using `appengine/urlfetch` on Google's appengine.
	// c := appengine.NewContext(r) // r is a *http.Request
	// t.Transport = urlfetch.Transport{Context: c}

	// Send SMS
	params := twilio.SMSParams{StatusCallback: callbackUrl}
	s, err := t.SendSMS(from, to, "Hello Go!", params)

	// or, make a voice call
	// params := twilio.CallParams{Url: callbackUrl}
	// s, err := w.MakeCall(from, to, params)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", s)
	return
}

func ExampleTwilio_MakeCall() {
	// Common stuffs
	AccountSid := "ac650108548e09aC2eed18ddb850c20b9"
	AuthToken := "2ecaf74387cbb28456aad6fb57b5ad682"
	from := "+15005550006"
	to := "+62801234567"
	callbackUrl := "http://subosito.com/"

	// Initialize twilio client
	t := twilio.NewTwilio(AccountSid, AuthToken)

	// Voice call
	params := twilio.CallParams{Url: callbackUrl, Timeout: 90}
	s, err := t.MakeCall(from, to, params)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", s)
	return
}
