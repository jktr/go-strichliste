package schema

import (
	"strings"
	"time"
)

type (
	ID        uint
	Barcode   string
	Currency  int // in cents
	Timestamp time.Time
)

const TimestampLayout = "2006-01-02 15:04:05"

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}
	pt, err := time.Parse(TimestampLayout, s)
	if err != nil {
		return nil
	}
	*t = Timestamp(pt)
	return nil
}
