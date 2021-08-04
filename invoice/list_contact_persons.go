package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Contact_Persons_List_contact_persons
//func (c *API) ListContactPersons(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data ListContactPersonsResponse, err error) {
func (c *API) ListContactPersons() (data ListContactPersonsResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         ContactsModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s/%s", c.ZohoTLD, ContactsModule, ContactsPersonSubModule),
		Method:       zoho.HTTPGet,
		ResponseData: &ListContactPersonsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
		BodyFormat: zoho.JSON_STRING,
		Headers: map[string]string{
			InvoiceAPIEndpointHeader: c.OrganizationID,
		},
	}

	/*for k, v := range params {
		endpoint.URLParameters[k] = v
	}*/

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ListContactPersonsResponse{}, fmt.Errorf("Failed to retrieve expense reports: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ListContactPersonsResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to list contact persons: %s", v.Message)
		}
		return *v, nil
	}
	return ListContactPersonsResponse{}, fmt.Errorf("Data retrieved was not 'ListContactPersonsResponse'")
}

type ListContactPersonsResponse struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	ContactPersons []struct {
		ContactID        string `json:"contact_id"`
		ContactPersonID  string `json:"contact_person_id"`
		Salutation       string `json:"salutation"`
		FirstName        string `json:"first_name"`
		LastName         string `json:"last_name"`
		Email            string `json:"email"`
		Phone            string `json:"phone,omitempty"`
		Mobile           string `json:"mobile,omitempty"`
		IsPrimaryContact bool   `json:"is_primary_contact"`
		Skype            string `json:"skype,omitempty"`
		Designation      string `json:"designation,omitempty"`
		Department       string `json:"department,omitempty"`
		IsAddedInPortal  bool   `json:"is_added_in_portal"`
	} `json:"contact_persons"`
}
