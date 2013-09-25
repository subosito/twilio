package twilio

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestException(t *testing.T) {
	data := `{
		"status": 400,
		"message": "No to number is specified",
		"code": 21201,
		"more_info": "http:\/\/www.twilio.com\/docs\/errors\/21201"
	}`

	ex := new(Exception)
	err := json.Unmarshal([]byte(data), &ex)
	if err != nil {
		t.Errorf("json.Unmarshal returned an error %+v", err)
	}

	want := &Exception{
		Status:   400,
		Message:  "No to number is specified",
		Code:     21201,
		MoreInfo: "http://www.twilio.com/docs/errors/21201",
	}

	if !reflect.DeepEqual(ex, want) {
		t.Errorf("Exception returned %+v, want %+v", ex, want)
	}
}

func TestException_Error(t *testing.T) {
	ex := &Exception{
		Status:   400,
		Message:  "No to number is specified",
		Code:     21201,
		MoreInfo: "http://www.twilio.com/docs/errors/21201",
	}

	want := "21201: No to number is specified"

	if ex.Error() != want {
		t.Errorf("Exception.Error returned %q, want %q", ex.Error(), want)
	}
}
