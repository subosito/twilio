package twilio

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestStructToValues(t *testing.T) {
	type Foo struct {
		Bar string
		Baz string
	}

	f := &Foo{Bar: "hello", Baz: "world"}
	v := structToValues(f)

	want := url.Values{}
	want.Set("Bar", "hello")
	want.Set("Baz", "world")

	if !reflect.DeepEqual(v, want) {
		t.Errorf("structToValues returned %+v, want %+v", v, want)
	}
}

func TestCheckResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader(`{"status": 400, "code": 21201, "message": "invalid parameter"}`)),
	}

	err := CheckResponse(res).(*Exception)

	if err == nil {
		t.Error("CheckResponse expected error response")
	}

	want := &Exception{
		Status:  400,
		Code:    21201,
		Message: "invalid parameter",
	}

	if !reflect.DeepEqual(err, want) {
		t.Errorf("Exception = %#v, want %#v", err, want)
	}
}
