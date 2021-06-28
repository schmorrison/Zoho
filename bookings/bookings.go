// Wrapper for Zoho Bookings API from https://www.zohoapis.com/bookings/v1/json/

package bookings

import (
	zoho "github.com/schmorrison/Zoho"
	"math/rand"
)

const (
	BookingsAPIEndpoint string = "https://www.zohoapis.in/bookings/v1/json/"
	GetAppointmentModule string = "getappointment"
	FetchWorkspacesModule string = "workspaces"
	FetchServicesModule string = "services"
	FetchStaffModule string = "staffs"
	FetchResourceModule string = "resources"
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
