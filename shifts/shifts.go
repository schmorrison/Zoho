package shifts

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

type Module = string

const (
	EmployeesModule    Module = "employees" // CRUD
	AvailabilityModule Module = "schedules" // CRUD - Shifts, Availability
	availabilityModule Module = "availability"
	shiftsModule       Module = "shifts"

	TimesheetsModule Module = "timesheets" // CRUD
	SettingsModule   Module = "settings"   // CRUD - Settings (Schedules, Positions, Job Sites)
	schedulesModule  Module = "schedules"
	positionsModule  Module = "positions"
	jobSitesModule   Module = "jobsites"

	TimeoffModule Module = "timeoff" // CRUD
)

// API is used for interacting with the Zoho Shifts API
// the exposed methods are primarily access to Shifts modules which provide access to Shifts Methods
type API struct {
	*zoho.Zoho
	id string
}

// New returns a *shifts.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			source := rand.NewSource(time.Now().UnixNano())
			rnd := rand.New(source)
			id = append(id, keyspace[rnd.Intn(len(keyspace))])
		}
		return string(id)
	}()

	return &API{
		Zoho: z,
		id:   id,
	}
}

const (
	timeLayout = ""
	dateLayout = "2006-01-02"
)

type Time time.Time

func (t Time) MarshalJSON() (b []byte, err error) {
	if t.IsZero() {
		return []byte{}, nil
	}

	return []byte(t.String()), nil
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	tm, err := time.Parse(timeLayout, s)
	if err != nil {
		return fmt.Errorf("failed to parse shifts.Time from JSON: %s", err)
	}
	*t = Time(tm)
	return nil
}

func (t *Time) String() string {
	tm := time.Time(*t)
	return fmt.Sprintf("%q", tm.Format(timeLayout))
}

func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}

type Date time.Time

func (d Date) MarshalJSON() (b []byte, err error) {
	if d.IsZero() {
		return []byte{}, nil
	}

	return []byte(d.String()), nil
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	dm, err := time.Parse(dateLayout, s)
	if err != nil {
		return fmt.Errorf("failed to parse shifts.Time from JSON: %s", err)
	}
	*d = Date(dm)
	return nil
}

func (d *Date) String() string {
	tm := time.Time(*d)
	return fmt.Sprintf("%q", tm.Format(dateLayout))
}

func (d Date) IsZero() bool {
	return time.Time(d).IsZero()
}
