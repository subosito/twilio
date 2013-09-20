package twilio

import (
	"encoding/json"
	"time"
)

type Price float32

func (p *Price) UnmarshalJSON(b []byte) error {
	return json.Unmarshal([]byte(unquote(b)), (*float32)(p))
}

type Timestamp time.Time

func (m *Timestamp) UnmarshalJSON(b []byte) error {
	s := unquote(b)

	if s == "null" {
		*m = Timestamp(time.Time{})
		return nil
	}

	t, err := time.Parse(time.RFC1123Z, s)
	if err != nil {
		return err
	}

	*m = Timestamp(t)
	return nil
}

func unquote(b []byte) string {
	switch b[0] {
	case '"':
		return string(b[1 : len(b)-1])
	default:
		return string(b)
	}
}
