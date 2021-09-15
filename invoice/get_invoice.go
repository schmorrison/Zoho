package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Invoices_Get_an_invoice
//func (c *API) GetInvoice(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data GetInvoiceResponse, err error) {
func (c *API) GetInvoice(invoiceId string) (data GetInvoiceResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         InvoicesModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s/%s", c.ZohoTLD, InvoicesModule, invoiceId),
		Method:       zoho.HTTPGet,
		ResponseData: &GetInvoiceResponse{},
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
		return GetInvoiceResponse{}, fmt.Errorf("Failed to retrieve invoice: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*GetInvoiceResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to retrieve invoice: %s", v.Message)
		}
		return *v, nil
	}
	return GetInvoiceResponse{}, fmt.Errorf("Data retrieved was not 'GetInvoiceResponse'")
}

type GetInvoiceResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Invoice struct {
		InvoiceId           string   `json:"invoice_id"`
		AchPaymentInitiated bool     `json:"ach_payment_initiated"`
		InvoiceNumber       string   `json:"invoice_number"`
		IsPreGst            bool     `json:"is_pre_gst"`
		PlaceOfSupply       string   `json:"place_of_supply"`
		GstNo               string   `json:"gst_no"`
		GstTreatment        string   `json:"gst_treatment"`
		Date                string   `json:"date"`
		Status              string   `json:"status"`
		PaymentTerms        int64    `json:"payment_terms"`
		PaymentTermsLabel   string   `json:"payment_terms_label"`
		DueDate             string   `json:"due_date"`
		PaymentExpectedDate string   `json:"payment_expected_date"`
		LastPaymentDate     string   `json:"last_payment_date"`
		ReferenceNumber     string   `json:"reference_number"`
		CustomerId          string   `json:"customer_id"`
		CustomerName        string   `json:"customer_name"`
		ContactPersons      []string `json:"contact_persons"`
		CurrencyId          string   `json:"currency_id"`
		CurrencyCode        string   `json:"currency_code"`
		ExchangeRate        float64  `json:"exchange_rate"`
		Discount            float64  `json:"discount"`
		IsDiscountBeforeTax bool     `json:"is_discount_before_tax"`
		DiscountType        string   `json:"discount_type"`
		IsInclusiveTax      bool     `json:"is_inclusive_tax"`
		RecurringInvoiceId  string   `json:"recurring_invoice_id"`
		IsViewedByClient    bool     `json:"is_viewed_by_client"`
		HasAttachment       bool     `json:"has_attachment"`
		ClientViewedTime    string   `json:"client_viewed_time"`
		LineItems           []struct {
			LineItemId     string   `json:"line_item_id"`
			ItemId         string   `json:"item_id"`
			ProjectId      string   `json:"project_id"`
			ProjectName    string   `json:"project_name"`
			ItemType       string   `json:"item_type"`
			ProductType    string   `json:"product_type"`
			TimeEntryIds   []string `json:"time_entry_ids"`
			ExpenseId      string   `json:"expense_id"`
			Name           string   `json:"name"`
			ItemOrder      float64  `json:"item_order"`
			BcyRate        float64  `json:"bcy_rate"`
			Rate           float64  `json:"rate"`
			Quantity       int64    `json:"quantity"`
			Unit           string   `json:"unit"`
			DiscountAmount float64  `json:"discount_amount"`
			Discount       float64  `json:"discount"`
			TaxId          string   `json:"tax_id"`
			TaxName        string   `json:"tax_name"`
			TaxType        string   `json:"tax_type"`
			TaxPercentage  float64  `json:"tax_percentage"`
			ItemTotal      float64  `json:"item_total"`
		} `json:"line_items"`
		ShippingCharge        float64 `json:"shipping_charge"`
		Adjustment            float64 `json:"adjustment"`
		AdjustmentDescription string  `json:"adjustment_description"`
		SubTotal              float64 `json:"sub_total"`
		TaxTotal              float64 `json:"tax_total"`
		Total                 float64 `json:"total"`
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
		BillingAddress         struct {
			Address string `json:"address"`
			Street2 string `json:"street2"`
			City    string `json:"city"`
			State   string `json:"state"`
			Zip     string `json:"zip"`
			Country string `json:"country"`
			Fax     string `json:"fax"`
		} `json:"billing_address"`
		ShippingAddress struct {
			Address string `json:"address"`
			City    string `json:"city"`
			State   string `json:"state"`
			Zip     string `json:"zip"`
			Country string `json:"country"`
			Fax     string `json:"fax"`
		} `json:"shipping_address"`
		Notes string `json:"notes"`
		Terms string `json:"terms"`
		/*CustomFields []struct {
		    CustomfieldId int64  `json:"customfield_id"`
		    DataType      string `json:"data_type"`
		    Index         int64  `json:"index"`
		    Label         string `json:"label"`
		    ShowOnPdf     bool   `json:"show_on_pdf"`
		    ShowInAllPdf  bool   `json:"show_in_all_pdf"`
		    Value         int64  `json:"value"`
		} `json:"custom_fields"`*/
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
