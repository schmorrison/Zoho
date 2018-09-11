package zoho

import (
	"fmt"
	"strings"
	"time"
)

type Time time.Time

var zohoTimeLayout = "2006-01-02T15:04:05-07:00"

func (t *Time) MarshalJSON() ([]byte, error) {
	if *t == Time(time.Time{}) {
		return []byte("null"), nil
	}
	stamp := fmt.Sprintf("\"%s\"", time.Time(*t).Format(zohoTimeLayout))
	return []byte(stamp), nil
}

func (t *Time) UnmarshalJSON(b []byte) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		blank := Time(time.Time{})
		t = &blank
		return
	}
	pTime, err := time.Parse(zohoTimeLayout, s)
	if err == nil {
		ref := Time(pTime)
		t = &ref
	}
	return
}

type Date time.Time

var zohoDateLayout = "2006-01-02"

func (d *Date) MarshalJSON() ([]byte, error) {
	if *d == Date(time.Time{}) {
		return []byte("null"), nil
	}
	stamp := fmt.Sprintf("\"%s\"", time.Time(*d).Format(zohoDateLayout))
	return []byte(stamp), nil
}

func (d *Date) UnmarshalJSON(b []byte) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		blank := Date(time.Time{})
		d = &blank
		return
	}
	pTime, err := time.Parse(zohoDateLayout, s)
	if err == nil {
		ref := Date(pTime)
		d = &ref
	}
	return
}
