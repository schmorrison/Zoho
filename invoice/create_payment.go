package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Customer_Payments_Create_a_payment
//func (c *API) CreatePayment(request interface{}, OrganizationID string, params map[string]zoho.Parameter) (data ListContactsResponse, err error) {
func (c *API) CreatePayment(request interface{}) (data CreatePaymentResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         ContactsModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s", c.ZohoTLD, ContactsModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreatePaymentResponse{},
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
		return CreatePaymentResponse{}, fmt.Errorf("Failed to create payment: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreatePaymentResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to create payment: %s", v.Message)
		}
		return *v, nil
	}
	return CreatePaymentResponse{}, fmt.Errorf("Data retrieved was not 'CreatePaymentResponse'")
}

type CreatePaymentRequest struct {
	CustomerId      string                 `json:"customer_id"`
	PaymentMode     string                 `json:"payment_mode"`
	Amount          float64                `json:"amount"`
	Date            string                 `json:"date"`
	ReferenceNumber string                 `json:"reference_number"`
	Description     string                 `json:"description"`
	Invoices        []CreatePaymentInvoice `json:"invoices"`
	ExchangeRate    float64                `json:"exchange_rate"`
	BankCharges     float64                `json:"bank_charges"`
	CustomFields    []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}
}

type CreatePaymentInvoice struct {
	CustomerId    string  `json:"invoice_id"`
	AmountApplied float64 `json:"amount_applied"`
}

type CreatePaymentResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Payment struct {
		PaymentId       string  `json:"payment_id"`
		PaymentMode     string  `json:"payment_mode"`
		Amount          float64 `json:"amount"`
		AmountRefunded  float64 `json:"amount_refunded"`
		BankCharges     float64 `json:"bank_charges"`
		Date            float64 `json:"date"`
		Status          string  `json:"status"`
		ReferenceNumber string  `json:"reference_number"`
		CustomerId      string  `json:"customer_id"`
		CustomerName    string  `json:"customer_name"`
		Email           string  `json:"email"`
		Invoices        []struct {
			InvoiceId        string  `json:"invoice_id"`
			InvoicePaymentId string  `json:"invoice_payment_id"`
			InvoiceNumber    string  `json:"invoice_number"`
			Date             string  `json:"date"`
			InvoiceAmount    float64 `json:"invoice_amount"`
			AmountApplied    float64 `json:"amount_applied"`
			BalanceAmount    float64 `json:"balance_amount"`
		}
		CurrencyCode   string `json:"currency_code"`
		CurrencySymbol string `json:"currency_symbol"`
		CustomFields   []struct {
			CustomfieldId int64  `json:"customfield_id"`
			DataType      string `json:"data_type"`
			Index         int64  `json:"index"`
			Label         string `json:"label"`
			ShowOnPdf     bool   `json:"show_on_pdf"`
			ShowInAllPdf  bool   `json:"show_in_all_pdf"`
			Value         string `json:"value"`
		}
	}
}
