package twilio

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type RecordingTransport struct {
	Transport http.RoundTripper
	Status    int
	Body      string
}

func (rx *RecordingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	re := &http.Response{
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Status:        http.StatusText(rx.Status),
		StatusCode:    rx.Status,
		Body:          ioutil.NopCloser(bytes.NewBufferString(rx.Body)),
		ContentLength: int64(len(rx.Body)),
		Request:       r,
	}

	return re, nil
}

var accountSid string = "AC65c2ee6fddb57b5ad6818ddb850c20b9"
var authToken string = "2ecaf0108548e09a74387cbb28456aa2"
var from string = "+15005550006"
var to string = "+62821234567"
var body string = "Hello Go!"
var callbackUrl = "http://subosito.com/"

var response = map[string]string{
	"SMS": `{
				"account_sid": "AC65c2ee6fddb57b5ad6818ddb850c20b9",
				"api_version": "2010-04-01",
				"body": "Hello raw response!",
				"date_created": "Mon, 16 Aug 2010 03:45:01 +0000",
				"date_sent": "Mon, 16 Aug 2010 03:45:03 +0000",
				"date_updated": "Mon, 16 Aug 2010 03:45:03 +0000",
				"direction": "outbound-api",
				"from": "+14158141829",
				"price": "-0.02000",
				"sid": "SM800f449d0399ed014aae2bcc0cc2f2ec",
				"status": "sent",
				"to": "+14159978453",
				"uri": "/2010-04-01/Accounts/AC65c2ee6fddb57b5ad6818ddb850c20b9/SMS/Messages/SM800f449d0399ed014aae2bcc0cc2f2ec.json"
			}`,
}
