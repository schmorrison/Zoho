package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Contacts_Get_a_Contact
//func (c *API) GetContact(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data GetContactResponse, err error) {
func (c *API) GetContact(contactId string) (data GetContactResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         InvoicesModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s/%s", c.ZohoTLD, ContactsModule, contactId),
		Method:       zoho.HTTPGet,
		ResponseData: &GetContactResponse{},
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
	  }
	*/

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetContactResponse{}, fmt.Errorf("Failed to retrieve contact: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*GetContactResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to retrieve contact: %s", v.Message)
		}
		return *v, nil
	}
	return GetContactResponse{}, fmt.Errorf("Data retrieved was not 'GetContactResponse'")
}

type GetContactResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Contact struct {
		ContactID                        string  `json:"contact_id"`
		ContactName                      string  `json:"contact_name"`
		CompanyName                      string  `json:"company_name"`
		HasTransaction                   bool    `json:"has_transaction"`
		ContactType                      string  `json:"contact_type"`
		IsTaxable                        bool    `json:"is_taxable"`
		TaxID                            string  `json:"tax_id"`
		TaxName                          string  `json:"tax_name"`
		TaxPercentage                    float64 `json:"tax_percentage"`
		TaxAuthorityID                   string  `json:"tax_authority_id"`
		TaxExemptionID                   string  `json:"tax_exemption_id"`
		GSTNo                            string  `json:"gst_no"`
		GSTTreatment                     string  `json:"gst_treatment"`
		IsLinkedWithZohocrm              bool    `json:"is_linked_with_zohocrm"`
		Website                          string  `json:"website"`
		PrimaryContactID                 string  `json:"primary_contact_id"`
		PaymentTerms                     int64   `json:"payment_terms"`
		PaymentTermsLabel                string  `json:"payment_terms_label"`
		CurrencyID                       string  `json:"currency_id"`
		CurrencyCode                     string  `json:"currency_code"`
		CurrencySymbol                   string  `json:"currency_symbol"`
		LanguageCode                     string  `json:"language_code"`
		OutstandingReceivableAmount      float64 `json:"outstanding_receivable_amount"`
		OutstandingReceivableAmountBcy   float64 `json:"outstanding_receivable_amount_bcy"`
		UnusedCreditsReceivableAmount    float64 `json:"unused_credits_receivable_amount"`
		UnusedCreditsReceivableAmountBcy float64 `json:"unused_credits_receivable_amount_bcy"`
		Status                           string  `json:"status"`
		Facebook                         string  `json:"facebook"`
		Twitter                          string  `json:"twitter"`
		PaymentReminderEnabled           bool    `json:"payment_reminder_enabled"`
		CustomFields                     []struct {
			Value string `json:"value"`
			Index int64  `json:"index"`
			Label string `json:"label"`
		} `json:"custom_fields"`
		BillingAddress   ContactAddress          `json:"billing_address"`
		ShippingAddress  ContactAddress          `json:"shipping_address"`
		ContactPersons   []ContactPerson         `json:"contact_persons"`
		DefaultTemplates ContactDefaultTemplates `json:"default_templates"`
		Notes            string                  `json:"notes"`
		CreatedTime      string                  `json:"created_time"`
		LastModifiedTime string                  `json:"last_modified_time"`
	} `json:"contact"`
}
