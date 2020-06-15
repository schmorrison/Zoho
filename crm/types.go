package crm

import (
	"encoding/json"
	"go-zoho/zoho"
)

type Error struct {
	Code    string `json:"code,omitempty"`
	Details struct {
	} `json:"details,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

type PageInfo struct {
	PerPage     int  `json:"per_page,omitempty"`
	Count       int  `json:"count,omitempty"`
	Page        int  `json:"page,omitempty"`
	MoreRecords bool `json:"more_records,omitempty"`
}

type MultiSelect []string
type Date = zoho.Date
type Time = zoho.Time
type Number int
type Currency float64
type Decimal float64
type Percent float64
type Long int64
type Checkbox bool
type URL string
type Lookup struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
type Owner struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
type Layout struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
type AutoNumber string

// SingleLine is the field type in Zoho that defines a single line input field
type SingleLine string

func (s *SingleLine) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t := SingleLine("")
		s = &t
		return nil
	}

	var t string
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	tp := SingleLine(t)
	s = &tp
	return nil
}

func (s SingleLine) MarshalJSON() ([]byte, error) {
	if string(s) == "" {
		return []byte{}, nil
	}
	return []byte(s), nil
}

// MultiLine is the field type in Zoho that defines a multiline input field, like text area in HTML
type MultiLine string

func (s *MultiLine) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t := MultiLine("")
		s = &t
		return nil
	}

	var t string
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	tp := MultiLine(t)
	s = &tp
	return nil
}

func (s MultiLine) MarshalJSON() ([]byte, error) {
	if string(s) == "" {
		return []byte{}, nil
	}
	return []byte(s), nil
}

// Email is the field type in Zoho that defines an email address field
type Email string

func (s *Email) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t := Email("")
		s = &t
		return nil
	}

	var t string
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	tp := Email(t)
	s = &tp
	return nil
}

func (s Email) MarshalJSON() ([]byte, error) {
	if string(s) == "" {
		return []byte{}, nil
	}
	return []byte(s), nil
}

// Phone is the field type in Zoho that defines a phone number field
type Phone string

func (s *Phone) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t := Phone("")
		s = &t
		return nil
	}

	var t string
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	tp := Phone(t)
	s = &tp
	return nil
}

func (s Phone) MarshalJSON() ([]byte, error) {
	if string(s) == "" {
		return []byte{}, nil
	}
	return []byte(s), nil
}

// PickList is the field type in Zoho that defines a dropdown that has been selected
type PickList string

func (s *PickList) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t := PickList("")
		s = &t
		return nil
	}

	var t string
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	tp := PickList(t)
	s = &tp
	return nil
}

func (s PickList) MarshalJSON() ([]byte, error) {
	if string(s) == "" {
		return []byte{}, nil
	}
	return []byte(s), nil
}
