package twilio

import (
	"reflect"
	"testing"
)

func testPrice(t *testing.T) {
	p := &Price{}
	err := p.UnmarshalJSON([]byte(`0.74`))
	if err != nil {
		t.Error("Price.UnmarshalJSON returned an error %q", err)
	}

	want := &Price{float32(0.74)}
	if !reflect.DeepEqual(p, want) {
		t.Errorf("Price.UnmarshalJSON returned %+v, want %+v", p, want)
	}
}

func TestPrice_UnmarshalJSON_string(t *testing.T) {
	p := &Price{}
	err := p.UnmarshalJSON([]byte(`"0.74"`))
	if err != nil {
		t.Error("Price.UnmarshalJSON returned an error %q", err)
	}

	want := &Price{float32(0.74)}
	if !reflect.DeepEqual(p, want) {
		t.Errorf("Price.UnmarshalJSON returned %+v, want %+v", p, want)
	}
}

func TestPrice_UnmarshalJSON_null(t *testing.T) {
	p := &Price{}
	err := p.UnmarshalJSON([]byte(`null`))
	if err != nil {
		t.Error("Price.UnmarshalJSON returned an error %q", err)
	}

	want := &Price{float32(0)}
	if !reflect.DeepEqual(p, want) {
		t.Errorf("Price.UnmarshalJSON returned %+v, want %+v", p, want)
	}
}

func TestPrice_Float64(t *testing.T) {
	p := &Price{0.75}
	want := float64(0.75)

	if p.Float64() != want {
		t.Errorf("Float64 returned %+v, want %+v", p.Float64(), want)
	}
}
