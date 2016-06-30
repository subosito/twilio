package twilio

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	exception := new(Exception)
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &exception)
	}

	return exception
}

func structToUrlValues(i interface{}) url.Values {
	v := url.Values{}
	m := structToMapString(i)
	for k, s := range m {
		switch {
		case len(s) == 1:
			v.Set(k, s[0])
		case len(s) > 1:
			for i := range s {
				v.Add(k, s[i])
			}
		}
	}

	return v
}

// structToMapString converts struct as map string
func structToMapString(i interface{}) map[string][]string {
	ms := map[string][]string{}
	iv := reflect.ValueOf(i).Elem()
	tp := iv.Type()

	for i := 0; i < iv.NumField(); i++ {
		k := tp.Field(i).Name
		f := iv.Field(i)
		ms[k] = valueToString(f)
	}

	return ms
}

// valueToString converts supported type of f as slice string
func valueToString(f reflect.Value) []string {
	var v []string

	switch reflect.TypeOf(f.Interface()).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = []string{strconv.FormatInt(f.Int(), 10)}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = []string{strconv.FormatUint(f.Uint(), 10)}
	case reflect.Float32:
		v = []string{strconv.FormatFloat(f.Float(), 'f', 4, 32)}
	case reflect.Float64:
		v = []string{strconv.FormatFloat(f.Float(), 'f', 4, 64)}
	case reflect.Bool:
		v = []string{strconv.FormatBool(f.Bool())}
	case reflect.Slice:
		for i := 0; i < f.Len(); i++ {
			if s := valueToString(f.Index(i)); len(s) == 1 {
				v = append(v, s[0])
			}
		}
	case reflect.String:
		v = []string{f.String()}
	}

	return v
}
