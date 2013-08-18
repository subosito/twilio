# Twilio

[![Build status](http://goci.me/project/image/github.com/subosito/twilio)](http://goci.me/project/github.com/subosito/twilio)

Simple Twilio API wrapper in Go.

## Usage

As usual you can `go get` the twilio package by issuing:

```bash
$ go get github.com/subosito/twilio
```

Then you can use it on your application:

```go
package main

import (
	"fmt"
	"github.com/subosito/twilio"
)

func main() {
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
```

## Resources

Documentation: http://godoc.org/github.com/subosito/twilio

