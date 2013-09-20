# Twilio

[![Build Status](https://travis-ci.org/subosito/twilio.png)](https://travis-ci.org/subosito/twilio)

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

var (
	AccountSid = "AC5ef8732a3c49700934481addd5ce1659"
	AuthToken  = "2ecaf0108548e09a74387cbb28456aa2"
)

func main() {
	// Initialize twilio client
	c := twilio.NewClient(AccountSid, AuthToken, nil)

	// You can set custom Client, eg: you're using `appengine/urlfetch` on Google's appengine
	// a := appengine.NewContext(r) // r is a *http.Request
	// f := urlfetch.Client(c)
	// c := twilio.NewClient(AccountSid, AuthToken, f)

	// Send Message
	params := twilio.MessageParams{
		Body: "Hello Go!",
	}
	s, err := c.Message.Send("+15005550006", "+62801234567", params)

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

