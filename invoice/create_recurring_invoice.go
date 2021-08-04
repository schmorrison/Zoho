package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Recurring_Invoices_Create_a_Recurring_Invoice
//func (c *API) CreateRecurringInvoice(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data ListContactsResponse, err error) {
func (c *API) CreateRecurringInvoice(request interface{}) (data CreateRecurringInvoiceResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         RecurringInvoicesModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s", c.ZohoTLD, RecurringInvoicesModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateRecurringInvoiceResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
		RequestBody: request,
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
		return CreateRecurringInvoiceResponse{}, fmt.Errorf("Failed to create recurring invoice: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateRecurringInvoiceResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to create recurring invoice: %s", v.Message)
		}
		return *v, nil
	}
	return CreateRecurringInvoiceResponse{}, fmt.Errorf("Data retrieved was not 'CreateRecurringInvoiceResponse'")
}

type CreateRecurringInvoiceRequest struct {
	RecurrenceName      string               `json:"recurrence_name"`
	ReferenceNumber     string               `json:"reference_number,omitempty"`
	CustomerId          string               `json:"customer_id"`
	TemplateId          string               `json:"template_id"`
	SalespersonId       string               `json:"salesperson_id,omitempty"`
	IsInclusiveTax      bool                 `json:"is_inclusive_tax,omitempty"`
	ContactPersons      []string             `json:"contact_persons,omitempty"`
	StartDate           string               `json:"start_date"`
	EndDate             string               `json:"end_date,omitempty"`
	PlaceOfSupply       string               `json:"place_of_supply,omitempty"`
	GstTreatment        string               `json:"gst_treatment,omitempty"`
	GstNo               string               `json:"gst_no,omitempty"`
	RecurrenceFrequency string               `json:"recurrence_frequency"`
	RepeatEvery         int64                `json:"repeat_every,omitempty"`
	PaymentTerms        int64                `json:"payment_terms,omitempty"`
	PaymentTermsLabel   string               `json:"payment_terms_label,omitempty"`
	CustomFields        []CustomFieldRequest `json:"custom_fields,omitempty"`
	LineItems           []InvoiceLineItem    `json:"line_items,omitempty"`
	TaxId               string               `json:"tax_id,omitempty"`
	Email               string               `json:"email,omitempty"`
	PaymentOptions      PaymentOptions       `json:"payment_options,omitempty"`
	TaxAuthorityId      string               `json:"tax_authority_id,omitempty"`
	TaxExemptionId      string               `json:"tax_exemption_id,omitempty"`
}

/*
type RecurringInvoiceLineItem struct {
	ItemId         string  `json:"item_id"`
	Name           string  `json:"name,omitempty"`
	Description    string  `json:"description,omitempty"`
	Rate           float64 `json:"rate,omitempty"`
	Quantity       int64   `json:"quantity"`
	Discount       float64 `json:"discount,omitempty"`
	TaxId          string  `json:"tax_id,omitempty"`
	TaxExemptionId string  `json:"tax_exemption_id,omitempty"`
	ItemTotal      float64 `json:"item_total,omitempty"`
	ProductType    string  `json:"product_type,omitempty"`
	HsnOrSac       int64   `json:"hsn_or_sac,omitempty"`
	ProjectId      string  `json:"project_id,omitempty"`
}
*/

type CreateRecurringInvoiceResponse struct {
	Code             int64  `json:"code"`
	Message          string `json:"message"`
	RecurringInvoice struct {
		RecurringInvoiceId string `json:"recurring_invoice_id"`
		RecurrenceName     string `json:"recurrence_name"`
		ReferenceNumber    string `json:"reference_number"`
		IsPreGst           bool   `json:"is_pre_gst"`
		GstNo              string `json:"gst_no"`
		GstTreatment       string `json:"gst_treatment"`
		PlaceOfSupply      string `json:"place_of_supply"`
		CustomerName       string `json:"customer_name"`
		CustomerId         string `json:"customer_id"`
		CurrencyId         string `json:"currency_id"`
		CurrencyCode       string `json:"currency_code"`
		StartDate          string `json:"start_date"`
		EndDate            string `json:"end_date"`
		LastSentDate       string `json:"last_sent_date"`
		NextInvoiceDate    string `json:"next_invoice_date"`
		LineItems          []struct {
			LineItemId  string  `json:"line_item_id"`
			Quantity    int64   `json:"quantity"`
			Name        string  `json:"name"`
			ItemTotal   float64 `json:"item_total"`
			Sku         string  `json:"sku"`
			ProductType string  `json:"product_type"`
			ProjectId   string  `json:"project_id"`
			ProjectName string  `json:"project_name"`
		} `json:"line_items"`
		BillingAddress  ContactAddress `json:"billing_address"`
		ShippingAddress ContactAddress `json:"shipping_address"`
		CustomFields    []struct {
			CustomfieldId string `json:"customfield_id"`
			DataType      string `json:"data_type"`
			Index         int64  `json:"index"`
			IsActive      bool   `json:"is_active"`
			Label         string `json:"label"`
			ShowInAllPdf  bool   `json:"show_in_all_pdf"`
			ShowOnPdf     bool   `json:"show_on_pdf"`
			Value         string `json:"value"`
		} `json:"custom_fields"`
		PaymentOptions PaymentOptions `json:"payment_options"`
	} `json:"recurring_invoice"`
}

type PaymentOptions struct {
	PaymentGateways []PaymentGateway `json:"payment_gateways,omitempty"`
}

type PaymentGateway struct {
	AdditionalField1     string `json:"additional_field1,omitempty"`
	Configured           bool   `json:"configured,omitempty"`
	GatewayName          string `json:"gateway_name,omitempty"`
	GatewayNameFormatted string `json:"gateway_name_formatted,omitempty"`
}
