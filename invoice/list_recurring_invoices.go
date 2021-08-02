package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Recurring_Invoices_List_Recurring_Invoice
//func (c *API) ListRecurringInvoices(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data ListRecurringInvoicesResponse, err error) {
func (c *API) ListRecurringInvoices() (data ListRecurringInvoicesResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:          RecurringInvoicesModule,
		URL:           fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s", c.ZohoTLD, RecurringInvoicesModule),
		Method:        zoho.HTTPGet,
		ResponseData:  &ListRecurringInvoicesResponse{},
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
		return ListRecurringInvoicesResponse{}, fmt.Errorf("Failed to retrieve expense reports: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ListRecurringInvoicesResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to list recurring invoices: %s", v.Message)
		}
		return *v, nil
	}
	return ListRecurringInvoicesResponse{}, fmt.Errorf("Data retrieved was not 'ListRecurringInvoicesResponse'")
}

type ListRecurringInvoicesResponse struct {
	Code              int    `json:"code"`
	Message           string `json:"message"`
	RecurringInvoices []struct {
		RecurringInvoiceId  string  `json:"recurring_invoice_id"`
		RecurrenceName      string  `json:"recurrence_name"`
		ReferenceNumber     string  `json:"reference_number"`
		Status              string  `json:"status"`
		Total               float64 `json:"total"`
		CustomerId          string  `json:"customer_id"`
		CustomerName        string  `json:"customer_name"`
		StartDate           string  `json:"start_date"`
		EndDate             string  `json:"end_date"`
		LastSentDate        string  `json:"last_sent_date"`
		NextInvoiceDate     string  `json:"next_invoice_date"`
		RecurrenceFrequency string  `json:"recurrence_frequency"`
		RepeatEvery         int64   `json:"repeat_every"`
	} `json:"recurring_invoices"`
}
