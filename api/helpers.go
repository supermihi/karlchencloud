package api

import "time"

func NewTimestamp(time time.Time) *Timestamp {
	return &Timestamp{Nanos: time.UnixNano()}
}
