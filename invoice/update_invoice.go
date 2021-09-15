package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Invoices_Update_an_invoice
//func (c *API) UpdateRecurringInvoice(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data UpdateInvoiceResponse, err error) {
func (c *API) UpdateInvoice(request interface{}, invoiceId string) (data UpdateInvoiceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         ContactsModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s/%s", c.ZohoTLD, InvoicesModule, invoiceId),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateInvoiceResponse{},
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
		return UpdateInvoiceResponse{}, fmt.Errorf("Failed to update invoice: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateInvoiceResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to update invoice: %s", v.Message)
		}
		return *v, nil
	}
	return UpdateInvoiceResponse{}, fmt.Errorf("Data retrieved was not 'UpdateInvoiceResponse'")
}

type UpdateInvoiceRequest struct {
	CustomerId string `json:"customer_id"`
	//ContactName          string   `json:"contact_name,omitempty"`
	ContactPersons        []string             `json:"contact_persons,omitempty"`
	InvoiceNumber         string               `json:"invoice_number,omitempty"`
	ReferenceNumber       string               `json:"reference_number,omitempty"`
	PlaceOfSupply         string               `json:"place_of_supply,omitempty"`
	GstTreatment          string               `json:"gst_treatment,omitempty"`
	GstNo                 string               `json:"gst_no,omitempty"`
	TemplateId            string               `json:"template_id,omitempty"`
	Date                  string               `json:"date,omitempty"`
	PaymentTerms          int64                `json:"payment_terms,omitempty"`
	PaymentTermsLabel     string               `json:"payment_terms_label,omitempty"`
	DueDate               string               `json:"due_date,omitempty"`
	Discount              float64              `json:"discount,omitempty"`
	IsDiscountBeforeTax   bool                 `json:"is_discount_before_tax,omitempty"`
	DiscountType          string               `json:"discount_type,omitempty"`
	IsInclusiveTax        bool                 `json:"is_inclusive_tax,omitempty"`
	ExchangeRate          float64              `json:"exchange_rate,omitempty"`
	RecurringInvoiceId    string               `json:"recurring_invoice_id,omitempty"`
	InvoicedEstimateId    string               `json:"invoiced_estimate_id,omitempty"`
	SalespersonId         string               `json:"salesperson_id,omitempty"`
	CustomFields          []CustomFieldRequest `json:"custom_fields,omitempty"`
	ProjectId             string               `json:"project_id,omitempty"`
	LineItems             []InvoiceLineItem    `json:"line_items"`
	PaymentOptions        PaymentOptions       `json:"payment_options"`
	AllowPartialPayments  bool                 `json:"allow_partial_payments"`
	CustomBody            string               `json:"custom_body,omitempty"`
	CustomSubject         string               `json:"custom_subject,omitempty"`
	Notes                 string               `json:"notes,omitempty"`
	Terms                 string               `json:"terms,omitempty"`
	ShippingCharge        float64              `json:"shipping_charge,omitempty"`
	Adjustment            float64              `json:"adjustment,omitempty"`
	AdjustmentDescription string               `json:"adjustment_description"`
	Reason                string               `json:"reason,omitempty"`
	TaxAuthorityId        string               `json:"tax_authority_id,omitempty"`
	TaxExemptionId        string               `json:"tax_exemption_id,omitempty"`
}

type UpdateInvoiceResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Invoice struct {
		InvoiceId             string            `json:"invoice_id"`
		AchPaymentInitiated   bool              `json:"ach_payment_initiated"`
		InvoiceNumber         string            `json:"invoice_number"`
		IsPreGst              bool              `json:"is_pre_gst"`
		PlaceOfSupply         string            `json:"place_of_supply"`
		GstNo                 string            `json:"gst_no"`
		GstTreatment          string            `json:"gst_treatment"`
		Date                  string            `json:"date"`
		Status                string            `json:"status"`
		PaymentTerms          int64             `json:"payment_terms"`
		PaymentTermsLabel     string            `json:"payment_terms_label"`
		DueDate               string            `json:"due_date"`
		PaymentExpectedDate   string            `json:"payment_expected_date"`
		LastPaymentDate       string            `json:"last_payment_date"`
		ReferenceNumber       string            `json:"reference_number"`
		CustomerId            string            `json:"customer_id"`
		CustomerName          string            `json:"customer_name"`
		ContactPersons        []string          `json:"contact_persons"`
		CurrencyId            string            `json:"currency_id"`
		CurrencyCode          string            `json:"currency_code"`
		ExchangeRate          float64           `json:"exchange_rate"`
		Discount              float64           `json:"discount"`
		IsDiscountBeforeTax   bool              `json:"is_discount_before_tax"`
		DiscountType          string            `json:"discount_type"`
		IsInclusiveTax        bool              `json:"is_inclusive_tax"`
		RecurringInvoiceId    string            `json:"recurring_invoice_id"`
		IsViewedByClient      bool              `json:"is_viewed_by_client"`
		HasAttachment         bool              `json:"has_attachment"`
		ClientViewedTime      string            `json:"client_viewed_time"`
		LineItems             []InvoiceLineItem `json:"line_items"`
		ShippingCharge        float64           `json:"shipping_charge"`
		Adjustment            float64           `json:"adjustment"`
		AdjustmentDescription string            `json:"adjustment_description"`
		SubTotal              float64           `json:"sub_total"`
		TaxTotal              float64           `json:"tax_total"`
		Total                 float64           `json:"total"`
		Taxes                 []struct {
			TaxName   string  `json:"tax_name"`
			TaxAmount float64 `json:"tax_amount"`
		} `json:"taxes"`
		PaymentReminderEnabled bool           `json:"payment_reminder_enabled"`
		PaymentMade            float64        `json:"payment_made"`
		CreditsApplied         float64        `json:"credits_applied"`
		TaxAmountWithheld      float64        `json:"tax_amount_withheld"`
		Balance                float64        `json:"balance"`
		WriteOffAmount         float64        `json:"write_off_amount"`
		AllowPartialPayments   bool           `json:"allow_partial_payments"`
		PricePrecision         int64          `json:"price_precision"`
		PaymentOptions         PaymentOptions `json:"payment_options"`
		IsEmailed              bool           `json:"is_emailed"`
		RemindersSent          int64          `json:"reminders_sent"`
		LastReminderSentDate   string         `json:"last_reminder_sent_date"`
		BillingAddress         ContactAddress `json:"billing_address"`
		ShippingAddress        ContactAddress `json:"shipping_address"`
		Notes                  string         `json:"notes"`
		Terms                  string         `json:"terms"`
		CustomFields           []struct {
			CustomfieldId string `json:"customfield_id"`
			DataType      string `json:"data_type"`
			Index         int64  `json:"index"`
			Label         string `json:"label"`
			ShowOnPdf     bool   `json:"show_on_pdf"`
			ShowInAllPdf  bool   `json:"show_in_all_pdf"`
			Value         string `json:"value"`
		} `json:"custom_fields"`
		TemplateId       string `json:"template_id"`
		TemplateName     string `json:"template_name"`
		CreatedTime      string `json:"created_time"`
		LastModifiedTime string `json:"last_modified_time"`
		AttachmentName   string `json:"attachment_name"`
		CanSendInMail    bool   `json:"can_send_in_mail"`
		SalespersonId    string `json:"salesperson_id"`
		SalespersonName  string `json:"salesperson_name"`
		InvoiceUrl       string `json:"invoice_url"`
	} `json:"invoice"`
}
