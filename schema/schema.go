package schema

import (
	"strings"
	"time"
)

// The server doesn't use strict ISO timestamps, which trips up
// default JSON deserialization. Aliasing time.Time allows
// implementing custom deserialization.
type Timestamp time.Time

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
