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
		k := tp.Field(i).Name
		it := f.Interface()

		switch reflect.TypeOf(it).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(it)

			for ix := 0; ix < s.Len(); ix++ {
				v.Add(k, s.Index(ix).String())
			}

		case reflect.String:
			v.Set(k, f.String())
		}
	}

	return
}
