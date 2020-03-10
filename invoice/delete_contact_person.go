package invoice

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Contact_Persons_Delete_a_contact_person
//func (c *ZohoInvoiceAPI) DeleteContactPerson(request interface{}, organizationId string, params map[string]zoho.Parameter) (data DeleteContactPersonResponse, err error) {
func (c *ZohoInvoiceAPI) DeleteContactPerson(contactPersonID string) (data DeleteContactPersonResponse, err error) {

	// Renew token if necessary
	if c.Zoho.Token.CheckExpiry() {
		err := c.Zoho.RefreshTokenRequest()
		if err != nil {
			return DeleteContactPersonResponse{}, err
		}
	}

	endpoint := zoho.Endpoint{
		Name: ContactsModule,
		URL: fmt.Sprintf(
			InvoiceAPIEndPoint+"%s/%s/%s",
			ContactsModule,
			ContactsPersonSubModule,
			contactPersonID,
		),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteContactPersonResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}

	/*for k, v := range params {
		endpoint.URLParameters[k] = v
	}*/

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteContactPersonResponse{}, fmt.Errorf("Failed to delete contact person: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteContactPersonResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to delete contact person: %s", v.Message)
		}
		return *v, nil
	}
	return DeleteContactPersonResponse{}, fmt.Errorf("Data retrieved was not 'DeleteContactPersonResponse'")
}

type DeleteContactPersonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
