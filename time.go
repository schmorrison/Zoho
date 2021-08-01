package zoho

import (
	"fmt"
	"strings"
	"time"
)

// Time is a time.Time which can be marshalled/unmarshalled according to Zoho's specific time scheme
type Time time.Time

var zohoTimeLayout = "2006-01-02T15:04:05-07:00"

// MarshalJSON is the json marshalling function for Time internal type
func (t *Time) MarshalJSON() ([]byte, error) {
	if *t == Time(time.Time{}) {
		return []byte("null"), nil
	}
	stamp := fmt.Sprintf("\"%s\"", time.Time(*t).Format(zohoTimeLayout))
	return []byte(stamp), nil
}

// UnmarshalJSON is the json unmarshalling function for Time internal type
func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*t = Time(time.Time{})
		return nil
	}
	pTime, err := time.Parse(zohoTimeLayout, s)
	if err == nil {
		*t = Time(pTime)
	}
	return err
}

// Date iis a time.Time which can be marshalled/unmarshalled according to Zoho's specific date scheme
type Date time.Time

var zohoDateLayout = "2006-01-02"

// MarshalJSON is the json marshalling function for Date internal type
func (d *Date) MarshalJSON() ([]byte, error) {
	if *d == Date(time.Time{}) {
		return []byte("null"), nil
	}
	stamp := fmt.Sprintf("\"%s\"", time.Time(*d).Format(zohoDateLayout))
	return []byte(stamp), nil
}

// UnmarshalJSON is the json unmarshalling function for Date internal type
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*d = Date(time.Time{})
		return nil
	}
	pTime, err := time.Parse(zohoDateLayout, s)
	if err == nil {
		*d = Date(pTime)
	}
	return err
}
