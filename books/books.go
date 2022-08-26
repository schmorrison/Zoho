package books

import zoho "github.com/recap-technologies/Zoho"

// API is used for interacting with the Zoho CRM API
// the exposed methods are primarily access to CRM modules which provide access to CRM Methods
type API struct {
	*zoho.Zoho
	id string
}
