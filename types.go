package geckoboard

import (
	"time"
)

type Date struct {
	time.Time
}

func (d Date) String() string {
	return d.Format("2006-01-02")
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
}
func NewDate(t time.Time) *Date {
	result := new(Date)
	result.Time = t
	return result
}


type DateTime struct {
	time.Time
}
func (d DateTime) String() string {
	return d.Format("2006-01-02T15:04:05-0700")
}
func (d *DateTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
}
