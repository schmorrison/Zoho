package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Customer_Payments_Retrieve_a_payment
//func (c *API) RetrievePayment(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data RetrievePaymentResponse, err error) {
func (c *API) RetrievePayment(paymentId string) (data RetrievePaymentResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         CustomerPaymentsModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s/%s", c.ZohoTLD, CustomerPaymentsModule, paymentId),
		Method:       zoho.HTTPGet,
		ResponseData: &RetrievePaymentResponse{},
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
		return RetrievePaymentResponse{}, fmt.Errorf("Failed to retrieve payments: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*RetrievePaymentResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to retrieve payment: %s", v.Message)
		}
		return *v, nil
	}
	return RetrievePaymentResponse{}, fmt.Errorf("Data retrieved was not 'RetrievePaymentResponse'")
}

type RetrievePaymentResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Payment struct {
		PaymentId           string  `json:"payment_id"`
		PaymentMode         string  `json:"payment_mode"`
		Amount              float64 `json:"amount"`
		AmountRefunded      float64 `json:"amount_refunded"`
		BankCharges         float64 `json:"bank_charges"`
		Date                string  `json:"date"`
		Status              string  `json:"status"`
		ReferenceNumber     string  `json:"reference_number"`
		OnlineTransactionId string  `json:"online_transaction_id"`
		CustomerId          string  `json:"customer_id"`
		CustomerName        string  `json:"customer_name"`
		Email               string  `json:"email"`
		Invoices            []struct {
			InvoiceId        string  `json:"invoice_id"`
			InvoicePaymentId string  `json:"invoice_payment_id"`
			InvoiceNumber    string  `json:"invoice_number"`
			Date             string  `json:"date"`
			InvoiceAmount    float64 `json:"invoice_amount"`
			AmountApplied    float64 `json:"amount_applied"`
			BalanceAmount    float64 `json:"balance_amount"`
		} `json:"invoices"`
		CurrencyCode   string `json:"currency_code"`
		CurrencySymbol string `json:"currency_symbol"`
		/*CustomFields   []struct {
		    CustomfieldId int64  `json:"customfield_id"`
		    DataType      string `json:"data_type"`
		    Index         int64  `json:"index"`
		    Label         string `json:"label"`
		    ShowOnPdf     bool   `json:"show_on_pdf"`
		    ShowInAllPdf  bool   `json:"show_in_all_pdf"`
		    Value         int64  `json:"value"`
		} `json:"custom_fields"`*/
	} `json:"payment"`
}
