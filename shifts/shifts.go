package shifts

import (
	"math/rand"
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
