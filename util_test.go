package twilio

import (
	"net/url"
	"reflect"
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
