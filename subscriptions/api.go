package subscriptions

import (
	"math/rand"

	zoho "github.com/schmorrison/Zoho"
)

const ZohoSubscriptionsEndpointHeader = "X-com-zoho-subscriptions-organizationid"

// API is used for interacting with the Zoho Subscriptions API
type API struct {
	*zoho.Zoho
	id             string
	OrganizationID string
}

// New returns a *subscriptions.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho, organizationID string) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			id = append(id, keyspace[rand.Intn(len(keyspace))])
		}
		return string(id)
	}()

	return &API{
		id:             id,
		Zoho:           z,
		OrganizationID: organizationID,
	}
}
