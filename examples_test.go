package twilio_test

import (
	"fmt"
	"github.com/subosito/twilio"
)

// This example shows the usage of twilio package. You can get your AccountSid and AuthToken on Account Dashboard page.
func Example() error {
	// Prepare credentials
	AccountSid := "ac650108548e09aC2eed18ddb850c20b9"
	AuthToken := "2ecaf74387cbb28456aad6fb57b5ad682"

	// Initialize twilio client
	t := twilio.NewTwilio(AccountSid, AuthToken)

	// Send SMS
	params := map[string]string{"StatusCallback": "http://example.com/"}
	s, err := t.SendSMS("+15005550006", "+62821234567", "Hello Go!", params)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Printf("%+v\n", s)
	return nil
}
