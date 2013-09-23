package twilio

import (
	"strconv"
)

type Price struct {
	float32
}

func (p *Price) UnmarshalJSON(b []byte) (err error) {
	str := string(b)

	if str == "null" {
		*p = Price{float32(0)}
		return nil
	}

	ustr, _ := strconv.Unquote(str)
	f, err := strconv.ParseFloat(ustr, 32)
	if err == nil {
		*p = Price{float32(f)}
	}

	return err
}
