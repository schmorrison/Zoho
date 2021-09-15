package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Contacts_Create_a_Contact
//func (c *API) CreateContact(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data ListContactsResponse, err error) {
func (c *API) CreateContact(request interface{}, enablePortal bool) (data CreateContactResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         ContactsModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s", c.ZohoTLD, ContactsModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateContactResponse{},
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
		return CreateContactResponse{}, fmt.Errorf("Failed to create contact: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateContactResponse); ok {
		// Check if the request succeeded
		if v.Contact.ContactID == "" {
			return *v, fmt.Errorf("Failed to create contact: %s", v.Message)
		}
		// Enable portal if requested
		if enablePortal {
			endpoint := zoho.Endpoint{
				Name:         ContactsModule,
				URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s/%s/portal/enable", c.ZohoTLD, ContactsModule, v.Contact.ContactID),
				Method:       zoho.HTTPPost,
				ResponseData: &EnableContactDashboardResponse{},
				URLParameters: map[string]zoho.Parameter{
					"filter_by": "",
				},
				RequestBody: &EnableContactDashboardRequest{
					ContactPersons: []EnableContactDashboardRequestPersonID{{
						ContactPersonID: v.Contact.ContactPersons[0].ContactPersonID,
					}}},
				BodyFormat: zoho.JSON_STRING,
				Headers: map[string]string{
					InvoiceAPIEndpointHeader: c.OrganizationID,
				},
			}
			err = c.Zoho.HTTPRequest(&endpoint)
			if err != nil {
				return CreateContactResponse{}, fmt.Errorf("Failed to enable main person portal: %s", err)
			}
		}
		return *v, nil
	}
	return CreateContactResponse{}, fmt.Errorf("Data retrieved was not 'CreateContactResponse'")
}

type EnableContactDashboardRequest struct {
	ContactPersons []EnableContactDashboardRequestPersonID `json:"contact_persons"`
}

type EnableContactDashboardRequestPersonID struct {
	ContactPersonID string `json:"contact_person_id"`
}

type EnableContactDashboardResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CreateContactRequest struct {
	ContactName      string                  `json:"contact_name"`
	CompanyName      string                  `json:"company_name"`
	CustomerSubType  string                  `json:"customer_sub_type,omitempty"`
	PaymentTerms     int64                   `json:"payment_terms,omitempty"`
	CurrencyID       string                  `json:"currency_id,omitempty"`
	Website          string                  `json:"website,omitempty"`
	CustomFields     []CustomFieldRequest    `json:"custom_fields"`
	BillingAddress   ContactAddress          `json:"billing_address"`
	ShippingAddress  ContactAddress          `json:"shipping_address"`
	ContactPersons   []ContactPerson         `json:"contact_persons"`
	DefaultTemplates ContactDefaultTemplates `json:"default_templates,omitempty"`
	LanguageCode     string                  `json:"language_code"`
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

type CreateContactResponse struct {
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
	} `json:"contact"`
}

type ContactPerson struct {
	ContactPersonID  string `json:"contact_person_id,omitempty"`
	Salutation       string `json:"salutation,omitempty"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Mobile           string `json:"mobile,omitempty"`
	IsPrimaryContact bool   `json:"is_primary_contact,omitempty"`
}

type ContactAddress struct {
	Attention string `json:"attention,omitempty"`
	Address   string `json:"address,omitempty"`
	Street2   string `json:"street2,omitempty"`
	StateCode string `json:"state_code,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
	Fax       string `json:"fax,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type ContactDefaultTemplates struct {
	InvoiceTemplateID           string `json:"invoice_template_id,omitempty"`
	InvoiceTemplateName         string `json:"invoice_template_name,omitempty"`
	EstimateTemplateID          string `json:"estimate_template_id,omitempty"`
	EstimateTemplateName        string `json:"estimate_template_name,omitempty"`
	CreditnoteTemplateID        string `json:"creditnote_template_id,omitempty"`
	CreditnoteTemplateName      string `json:"creditnote_template_name,omitempty"`
	InvoiceEmailTemplateID      string `json:"invoice_email_template_id,omitempty"`
	InvoiceEmailTemplateName    string `json:"invoice_email_template_name,omitempty"`
	EstimateEmailTemplateID     string `json:"estimate_email_template_id,omitempty"`
	EstimateEmailTemplateName   string `json:"estimate_email_template_name,omitempty"`
	CreditnoteEmailTemplateID   string `json:"creditnote_email_template_id,omitempty"`
	CreditnoteEmailTemplateName string `json:"creditnote_email_template_name,omitempty"`
}
