package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Recurring_Invoices_Get_a_Recurring_Invoice
//func (c *API) GetRecurringInvoice(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data ListContactsResponse, err error) {
func (c *API) GetRecurringInvoice(recurringInvoiceId string) (data RecurringInvoiceResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         RecurringInvoicesModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s/%s", c.ZohoTLD, RecurringInvoicesModule, recurringInvoiceId),
		Method:       zoho.HTTPGet,
		ResponseData: &RecurringInvoiceResponse{},
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
		return RecurringInvoiceResponse{}, fmt.Errorf("Failed to retrieve recurring invoice: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*RecurringInvoiceResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to retrieve recurring invoice: %s", v.Message)
		}
		return *v, nil
	}
	return RecurringInvoiceResponse{}, fmt.Errorf("Data retrieved was not 'RecurringInvoiceResponse'")
}

type RecurringInvoiceResponse struct {
	Code             int64  `json:"code"`
	Message          string `json:"message"`
	RecurringInvoice struct {
		RecurringInvoiceId  string `json:"recurring_invoice_id"`
		RecurrenceName      string `json:"recurrence_name"`
		ReferenceNumber     string `json:"reference_number"`
		CustomerName        string `json:"customer_name"`
		CustomerId          string `json:"customer_id"`
		IsPreGst            bool   `json:"is_pre_gst"`
		GstNo               string `json:"gst_no"`
		GstTreatment        string `json:"gst_treatment"`
		PlaceOfSupply       string `json:"place_of_supply"`
		RecurrenceFrequency string `json:"recurrence_frequency"`
		CompanyName         string `json:"company_name"`
		CustomerEmail       string `json:"customer_email"`
		CustomerMobilePhone string `json:"customer_mobile_phone"`
		CustomerPhone       string `json:"customer_phone"`
		PhotoUrl            string `json:"photo_url"`
		CurrencyId          string `json:"currency_id"`
		CurrencyCode        string `json:"currency_code"`
		StartDate           string `json:"start_date"`
		EndDate             string `json:"end_date"`
		LastSentDate        string `json:"last_sent_date"`
		NextInvoiceDate     string `json:"next_invoice_date"`
		LineItems           []struct {
			LineItemId       string  `json:"line_item_id"`
			ItemId           string  `json:"item_id"`
			ItemOrder        float64 `json:"item_order"`
			DiscountAmount   float64 `json:"discount_amount"`
			Quantity         int64   `json:"quantity"`
			Rate             float64 `json:"rate"`
			Discount         float64 `json:"discount"`
			Name             string  `json:"name"`
			ItemTotal        float64 `json:"item_total"`
			Sku              string  `json:"sku"`
			ProductType      string  `json:"product_type"`
			ProjectId        string  `json:"project_id"`
			ProjectName      string  `json:"project_name"`
			ItemCustomFields []struct {
				CustomfieldID string `json:"customfield_id,omitempty"`
				Label         string `json:"label"`
				Value         string `json:"value,omitempty"`
			} `json:"item_custom_fields"`
		} `json:"line_items"`
		PaidInvoicesTotal     float64        `json:"paid_invoices_total"`
		UnpaidInvoicesBalance float64        `json:"unpaid_invoices_balance"`
		BillingAddress        ContactAddress `json:"billing_address"`
		ShippingAddress       ContactAddress `json:"shipping_address"`
		/*CustomFields []struct {
		    CustomfieldId int64  `json:"customfield_id"`
		    DataType      string `json:"data_type"`
		    Index         int64  `json:"index"`
		    IsActive      bool   `json:"is_active"`
		    Label         string `json:"label"`
		    ShowInAllPdf  bool   `json:"show_in_all_pdf"`
		    ShowOnPdf     bool   `json:"show_on_pdf"`
		    Value         string `json:"value"`
		} `json:"custom_fields"`*/
		PaymentOptions PaymentOptions `json:"payment_options"`
	} `json:"recurring_invoice"`
}
