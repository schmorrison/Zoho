// Wrapper for Zoho Bookings API from https://www.zohoapis.com/bookings/v1/json/

package bookings

import (
	zoho "github.com/schmorrison/Zoho"
	"math/rand"
)

type BookingsModule = string

const (
	GetAppointmentModule BookingsModule = "getappointment"
	GetAvailabilityModule BookingsModule = "availableslots"
	FetchWorkspacesModule BookingsModule = "workspaces"
	FetchServicesModule BookingsModule = "services"
	FetchStaffModule BookingsModule = "staffs"
	FetchResourceModule BookingsModule = "resources"
	BookAppointmentModule BookingsModule = "appointment"
	RescheduleAppointmentModule BookingsModule = "rescheduleappointment"
	UpdateAppointmentModule BookingsModule = "updateappointment"
)

// API is used for interacting with the Zoho expense API
// the exposed methods are primarily access to expense modules which provide access to expense Methods
type API struct {
	*zoho.Zoho
	id string
}

// New returns a *expense.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			id = append(id, keyspace[rand.Intn(len(keyspace))])
		}
		return string(id)
	}()
	API := &API{
		Zoho: z,
		id:   id,
	}
	return API
}
