package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Contact_Persons_Create_a_contact_person
//func (c *API) CreateContactPerson(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data CreateContactPersonResponse, err error) {
func (c *API) CreateContactPerson(request interface{}) (data CreateContactPersonResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         ContactsModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s/%s", c.ZohoTLD, ContactsModule, ContactsPersonSubModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateContactPersonResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
		RequestBody: &request,
		BodyFormat:  zoho.JSON_STRING,
		Headers: map[string]string{
			InvoiceAPIEndpointHeader: c.OrganizationID,
		},
	}

	/*for k, v := range params {
		endpoint.URLParameters[k] = v
	}*/

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateContactPersonResponse{}, fmt.Errorf("Failed to create contact person: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateContactPersonResponse); ok {
		// Do not test this endpoint code return against 0 as it may succeed with warnings
		/* Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to create contact person: %s", v.Message)
		}*/
		return *v, nil
	}
	return CreateContactPersonResponse{}, fmt.Errorf("Data retrieved was not 'CreateContactPersonResponse'")
}

type CreateContactPersonRequest struct {
	ContactID    string `json:"contact_id"`
	Salutation   string `json:"salutation,omitempty"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	Skype        string `json:"skype,omitempty"`
	Designation  string `json:"designation,omitempty"`
	Department   string `json:"department,omitempty"`
	EnablePortal bool   `json:"enable_portal"`
}

type CreateContactPersonResponse struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	ContactPerson struct {
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
	} `json:"contact_person"`
}
