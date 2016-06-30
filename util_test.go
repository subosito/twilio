package twilio

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

type StructTest struct {
	Int         int
	Uint        uint
	Float32     float32
	Float64     float64
	Bool        bool
	String      string
	SliceString []string
	SliceInt    []int
}

func TestStructToMapString(t *testing.T) {
	st := StructTest{
		Int:         -134020434,
		Uint:        4039455774,
		Float32:     0.06563702,
		Float64:     0.6868230728671094,
		Bool:        true,
		String:      "hello",
		SliceString: []string{"foo", "bar"},
		SliceInt:    []int{1, 2, 3},
	}

	w := map[string][]string{
		"Int":         []string{"-134020434"},
		"Uint":        []string{"4039455774"},
		"Float32":     []string{"0.0656"},
		"Float64":     []string{"0.6868"},
		"Bool":        []string{"true"},
		"String":      []string{"hello"},
		"SliceString": []string{"foo", "bar"},
		"SliceInt":    []string{"1", "2", "3"},
	}

	assert.Equal(t, w, StructToMapString(&st))
}
