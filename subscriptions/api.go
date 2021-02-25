package subscriptions

import (
	zoho "github.com/schmorrison/Zoho"
)

const ZohoSubscriptionsOriganizationId = "X-com-zoho-subscriptions-organizationid"

// API is used for interacting with the Zoho Subscriptions API
type API struct {
	*zoho.Zoho
	OrganizationId string
}

// New returns a *subscriptions.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho, organizationId string) *API {
	return &API{
		Zoho:           z,
		OrganizationId: organizationId,
	}
}
