package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Contacts_Update_a_Contact
//func (c *API) UpdateContact(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data UpdateContactResponse, err error) {
func (c *API) UpdateContact(request interface{}, contactId string) (data UpdateContactResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         ContactsModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s/%s", c.ZohoTLD, ContactsModule, contactId),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateContactResponse{},
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
		return UpdateContactResponse{}, fmt.Errorf("Failed to create contact: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateContactResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to update contact: %s", v.Message)
		}
		return *v, nil
	}
	return UpdateContactResponse{}, fmt.Errorf("Data retrieved was not 'UpdateContactResponse'")
}

type UpdateContactRequest struct {
	ContactName      string                  `json:"contact_name,omitempty"`
	CompanyName      string                  `json:"company_name,omitempty"`
	CustomerSubType  string                  `json:"customer_sub_type,omitempty"`
	PaymentTerms     int64                   `json:"payment_terms,omitempty"`
	CurrencyID       string                  `json:"currency_id,omitempty"`
	Website          string                  `json:"website,omitempty"`
	CustomFields     []CustomFieldRequest    `json:"custom_fields,omitempty"`
	BillingAddress   ContactAddress          `json:"billing_address,omitempty"`
	ShippingAddress  ContactAddress          `json:"shipping_address,omitempty"`
	ContactPersons   []ContactPerson         `json:"contact_persons,omitempty"`
	DefaultTemplates ContactDefaultTemplates `json:"default_templates,omitempty"`
	LanguageCode     string                  `json:"language_code,omitempty"`
	Notes            string                  `json:"notes,omitempty"`
	PlaceOfContact   string                  `json:"place_of_contact,omitempty"`
	GSTNo            string                  `json:"gst_no,omitempty"`
	GSTTreatment     string                  `json:"gst_treatment,omitempty"`
	TaxExemptionID   string                  `json:"tax_exemption_id,omitempty"`
	TaxAuthorityID   string                  `json:"tax_authority_id,omitempty"`
	TaxID            string                  `json:"tax_id,omitempty"`
	IsTaxable        string                  `json:"is_taxable,omitempty"`
	Facebook         string                  `json:"facebook,omitempty"`
	Twitter          string                  `json:"twitter,omitempty"`
}

type UpdateContactResponse struct {
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
		TaxExemptionID                   string  `json:"tax_exemption_id"`
		TaxAuthorityID                   string  `json:"tax_authority_id"`
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
		PaymentReminderEnabled           bool    `json:"payment_reminder_enabled"`
		CustomFields                     []struct {
			Value string `json:"value"`
			Index int64  `json:"index"`
			Label string `json:"label"`
		} `json:"custom_fields"`
		BillingAddress   ContactAddress          `json:"billing_address"`
		ShippingAddress  ContactAddress          `json:"shipping_address"`
		Facebook         string                  `json:"facebook"`
		Twitter          string                  `json:"twitter"`
		ContactPersons   []ContactPerson         `json:"contact_persons"`
		DefaultTemplates ContactDefaultTemplates `json:"default_templates"`
		Notes            string                  `json:"notes"`
		CreatedTime      string                  `json:"created_time"`
		LastModifiedTime string                  `json:"last_modified_time"`
	}
}
