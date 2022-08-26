package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Invoices_List_invoices
//func (c *API) ListInvoices(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data ListInvoicesResponse, err error) {
func (c *API) ListInvoices() (data ListInvoicesResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:          InvoicesModule,
		URL:           fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s", c.ZohoTLD, InvoicesModule),
		Method:        zoho.HTTPGet,
		ResponseData:  &ListInvoicesResponse{},
		URLParameters: map[string]zoho.Parameter{
			//"filter_by": "",
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
		return ListInvoicesResponse{}, fmt.Errorf("Failed to retrieve expense reports: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ListInvoicesResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to list invoices: %s", v.Message)
		}
		return *v, nil
	}
	return ListInvoicesResponse{}, fmt.Errorf("Data retrieved was not 'ListInvoicesResponse'")
}

// ListContactsResponse is the data returned by GetExpenseReports
type ListInvoicesResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Invoices []struct {
		InvoiceID            string  `json:"invoice_id"`
		AchPaymentInitiated  bool    `json:"ach_payment_initiated"`
		CustomerName         string  `json:"customer_name"`
		CustomerID           string  `json:"customer_id"`
		Status               string  `json:"status"`
		InvoiceNumber        string  `json:"invoice_number"`
		ReferenceNumber      string  `json:"reference_number"`
		Date                 string  `json:"date"`
		DueDate              string  `json:"due_date"`
		DueDays              string  `json:"due_days"`
		CurrencyID           string  `json:"currency_id"`
		ScheduleTime         string  `json:"schedule_time"`
		CurrencyCode         string  `json:"currency_code"`
		IsViewedByClient     bool    `json:"is_viewed_by_client"`
		HasAttachment        bool    `json:"has_attachment"`
		ClientViewedTime     string  `json:"client_viewed_time"`
		Total                float64 `json:"total"`
		Balance              float64 `json:"balance"`
		CreatedTime          string  `json:"created_time"`
		LastModifiedTime     string  `json:"last_modified_time"`
		IsEmailed            bool    `json:"is_emailed"`
		RemindersSent        int64   `json:"reminders_sent"`
		LastReminderSentDate string  `json:"last_reminder_sent_date"`
		PaymentExpectedDate  string  `json:"payment_expected_date"`
		LastPaymentDate      string  `json:"last_payment_date"`
		/*CustomFields  []struct {
			CustomfieldID string `json:"customfield_id"`
			Label         string `json:"label"`
			Value         string `json:"value"`
		} `json:"custom_fields"`*/
		Documents       string  `json:"documents"`
		SalespersonID   string  `json:"salesperson_id"`
		SalespersonName string  `json:"salesperson_name"`
		ShippingCharge  float32 `json:"shipping_charge"`
		Adjustment      float32 `json:"adjustment"`
		WriteOffAmount  float32 `json:"write_off_amount"`
		ExchangeRate    float32 `json:"exchange_rate"`
	} `json:"invoices"`
}
