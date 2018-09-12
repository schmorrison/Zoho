package crm

import "github.com/schmorrison/Zoho"

type SingleLine string
type MultiLine string
type Email string
type Phone string
type PickList string
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
