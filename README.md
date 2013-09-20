# Twilio

[![Build Status](https://drone.io/github.com/subosito/twilio/status.png)](https://drone.io/github.com/subosito/twilio/latest)
[![Coverage Status](https://coveralls.io/repos/subosito/twilio/badge.png?branch=master)](https://coveralls.io/r/subosito/twilio?branch=master)

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
	// f := urlfetch.Client(a)
	// c := twilio.NewClient(AccountSid, AuthToken, f)

	// Send Message
	params := twilio.MessageParams{
		Body: "Hello Go!",
	}
	s, err := c.Messages.Send("+15005550006", "+62801234567", params)
	if err == nil {
		fmt.Printf("%+v\n", s)
	}

	// You can also using lower level function: Create
	s, err := c.Messages.Create(url.Values{
		"From": {"+15005550006"},
		"To" :  {"+62801234567"},
		"Body": {"Hello Go!"},
	})
	if err == nil {
		fmt.Printf("%+v\n", s)
	}
}
```

## Resources

Documentation: http://godoc.org/github.com/subosito/twilio

## Status

Under heavy development. Please be patient :)

