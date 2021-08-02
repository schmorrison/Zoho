package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Recurring_Invoices_Stop_a_Recurring_Invoice
//func (c *API) StopRecurringInvoice(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data StopRecurringInvoiceResponse, err error) {
func (c *API) StopRecurringInvoice(recurringInvoiceId string) (data StopRecurringInvoiceResponse, err error) {

	endpoint := zoho.Endpoint{
		Name: RecurringInvoicesModule,
		URL: fmt.Sprintf(
			"https://invoice.zoho.%s/api/v3/%s/status/stop", c.ZohoTLD, recurringInvoiceId,
		),
		Method:       zoho.HTTPPost,
		ResponseData: &StopRecurringInvoiceResponse{},
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
	}*/

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return StopRecurringInvoiceResponse{}, fmt.Errorf("Failed to stop recurring invoice: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*StopRecurringInvoiceResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to stop recurring invoice: %s", v.Message)
		}
		return *v, nil
	}
	return StopRecurringInvoiceResponse{}, fmt.Errorf("Data retrieved was not 'StopRecurringInvoiceResponse'")
}

type StopRecurringInvoiceResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
