package twilio

import (
	"net/url"
	"reflect"
)

func structToValues(i interface{}) (v url.Values) {
	v = url.Values{}

	ival := reflect.ValueOf(i).Elem()
	tp := ival.Type()

	for i := 0; i < ival.NumField(); i++ {
		f := ival.Field(i)

		var val string

		switch f.Interface().(type) {
		case string:
			val = f.String()
		}

		v.Set(tp.Field(i).Name, val)
	}

	return
}
