package subscriptions

import (
	zoho "github.com/schmorrison/Zoho"
)

// API is used for interacting with the Zoho Subscriptions API
type API struct {
	*zoho.Zoho
}

// New returns a *subscriptions.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho) *API {
	return &API{
		Zoho: z,
	}
}
