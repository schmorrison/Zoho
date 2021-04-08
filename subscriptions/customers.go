package subscriptions

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetCustomer will return customer specified by id
// https://www.zoho.com/subscriptions/api/v1/#Invoices_Retrieve_a_subscription
func (s *API) GetCustomer(id string) (data CustomerResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "customers",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/customers/%s", s.ZohoTLD, id),
		Method:       zoho.HTTPGet,
		ResponseData: &CustomerResponse{},
		Headers: map[string]string{
			ZohoSubscriptionsEndpointHeader: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CustomerResponse{}, fmt.Errorf("Failed to retrieve customer (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*CustomerResponse); ok {
		return *v, nil
	}

	return CustomerResponse{}, fmt.Errorf("Data retrieved was not 'CustomerResponse'")
}

type CustomerResponse struct {
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	Customer struct {
		CustomerID  string `json:"customer_id"`
		DisplayName string `json:"display_name"`
		Salutation  string `json:"salutation"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		Tags        []struct {
			TagOptionID    string `json:"tag_option_id"`
			IsTagMandatory bool   `json:"is_tag_mandatory"`
			TagName        string `json:"tag_name"`
			TagID          string `json:"tag_id"`
			TagOptionName  string `json:"tag_option_name"`
		} `json:"tags"`
		CompanyName     string `json:"company_name"`
		Phone           string `json:"phone"`
		Mobile          string `json:"mobile"`
		Website         string `json:"website"`
		Designation     string `json:"designation"`
		Department      string `json:"department"`
		IsPortalEnabled bool   `json:"is_portal_enabled"`
		BillingAddress  struct {
			Attention   string `json:"attention"`
			Street      string `json:"street"`
			City        string `json:"city"`
			State       string `json:"state"`
			Zip         string `json:"zip"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			StateCode   string `json:"state_code"`
			Fax         string `json:"fax"`
		} `json:"billing_address"`
		ShippingAddress struct {
			Attention   string `json:"attention"`
			Street      string `json:"street"`
			City        string `json:"city"`
			State       string `json:"state"`
			Zip         string `json:"zip"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			StateCode   string `json:"state_code"`
			Fax         string `json:"fax"`
		} `json:"shipping_address"`
		CurrencyCode   string  `json:"currency_code"`
		CurrencyID     string  `json:"currency_id"`
		AchSupported   bool    `json:"ach_supported"`
		GstNo          string  `json:"gst_no"`
		GstTreatment   string  `json:"gst_treatment"`
		PlaceOfContact string  `json:"place_of_contact"`
		PricePrecision int64   `json:"price_precision"`
		UnusedCredits  float64 `json:"unused_credits"`
		Outstanding    float64 `json:"outstanding"`
		Notes          string  `json:"notes"`
		Status         string  `json:"status"`
		CustomFields   []struct {
			Index    int64  `json:"index"`
			Value    string `json:"value"`
			DataType string `json:"data_type"`
			Label    string `json:"label"`
		} `json:"custom_fields"`
		ZcrmAccountID          string `json:"zcrm_account_id"`
		ZcrmContactID          string `json:"zcrm_contact_id"`
		UpdatedTime            string `json:"updated_time"`
		CreatedTime            string `json:"created_time"`
		Source                 string `json:"source"`
		PaymentTermsLabel      string `json:"payment_terms_label"`
		IsLinkedWithZohocrm    bool   `json:"is_linked_with_zohocrm"`
		PrimaryContactpersonID string `json:"primary_contactperson_id"`
		CanAddCard             bool   `json:"can_add_card"`
		CanAddBankAccount      bool   `json:"can_add_bank_account"`
		DefaultTemplates       struct {
			InvoiceTemplateID    string `json:"invoice_template_id"`
			CreditnoteTemplateID string `json:"creditnote_template_id"`
		} `json:"default_templates"`
		Documents []struct {
			CanShowInPortal   bool   `json:"can_show_in_portal"`
			FileName          string `json:"file_name"`
			FileType          string `json:"file_type"`
			FileSize          int64  `json:"file_size"`
			FileSizeFormatted string `json:"file_size_formatted"`
			DocumentID        string `json:"document_id"`
			AttachmentOrder   int64  `json:"attachment_order"`
		} `json:"documents"`
	} `json:"customer"`
}
