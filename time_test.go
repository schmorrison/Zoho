package zoho

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Log("Testing time")
	t.Run("zoho time", func(t *testing.T) {
		timestamp := "2021-11-19T04:45:30-06:00"
		timeTest(t, timestamp, fmt.Sprintf("\"%s\"", timestamp))
	})

	t.Run("null time", func(t *testing.T) {
		timeTest(t, "null", "null")
	})

	t.Run("empty time", func(t *testing.T) {
		timeTest(t, "", "null")
	})

	t.Run("quoted time", func(t *testing.T) {
		timeTest(t, "\"null\"", "null")
	})
}

func timeTest(t *testing.T, teststamp string, expect string) {
	zt := Time{}
	if err := zt.UnmarshalJSON([]byte(teststamp)); err != nil {
		t.Errorf("zoho time: UnmarshalJSON failed '%s': %s", teststamp, err)
	}

	ts, err := zt.MarshalJSON()
	if err != nil {
		t.Errorf("zoho time: MarshalJSON failed '%s': %s", teststamp, err)
	}
	assert.Equalf(t, expect, string(ts), "Encode/Decode teststamp failed '%s': %s", teststamp, string(ts))
}

func TestDate(t *testing.T) {
	t.Log("Testing date")
	t.Run("zoho date", func(t *testing.T) {
		datestamp := "2021-11-19"
		dateTest(t, datestamp, fmt.Sprintf("\"%s\"", datestamp))
	})

	t.Run("null date", func(t *testing.T) {
		dateTest(t, "null", "null")
	})

	t.Run("empty date", func(t *testing.T) {
		dateTest(t, "", "null")
	})

	t.Run("quoted date", func(t *testing.T) {
		dateTest(t, "\"null\"", "null")
	})
}

func dateTest(t *testing.T, teststamp string, expect string) {
	zd := Date{}
	if err := zd.UnmarshalJSON([]byte(teststamp)); err != nil {
		t.Errorf("zoho date: UnmarshalJSON failed '%s': %s", teststamp, err)
	}

	ds, err := zd.MarshalJSON()
	if err != nil {
		t.Errorf("zoho date: MarshalJSON failed '%s': %s", teststamp, err)
	}
	assert.Equalf(t, expect, string(ds), "Encode/Decode teststamp failed '%s': %s", teststamp, string(ds))
}
