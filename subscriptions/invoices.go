package subscriptions

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

type InvoiceStatus string

// Proper names for Invoice statuses
const (
	InvoiceStatusAll           InvoiceStatus = "Status.All"
	InvoiceStatusSent          InvoiceStatus = "Status.Sent"
	InvoiceStatusDraft         InvoiceStatus = "Status.Draft"
	InvoiceStatusOverDue       InvoiceStatus = "Status.OverDue"
	InvoiceStatusPaid          InvoiceStatus = "Status.Paid"
	InvoiceStatusPartiallyPaid InvoiceStatus = "Status.PartiallyPaid"
	InvoiceStatusVoid          InvoiceStatus = "Status.Void"
	InvoiceStatusUnpaid        InvoiceStatus = "Status.Unpaid"
)

// ListInvoices will return the list of invoices that match the given invoice status.
// https://www.zoho.com/subscriptions/api/v1/#Invoices_List_all_invoices
func (s *API) ListInvoices(status InvoiceStatus) (data InvoicesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "invoices",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/invoices", s.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &InvoicesResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": zoho.Parameter(status),
		},
		Headers: map[string]string{
			ZohoSubscriptionsOriganizationID: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return InvoicesResponse{}, fmt.Errorf("Failed to retrieve invoices: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*InvoicesResponse); ok {
		return *v, nil
	}

	return InvoicesResponse{}, fmt.Errorf("Data retrieved was not 'InvoicesResponse'")
}

// GetInvoice will return the subscription specified by id
// https://www.zoho.com/subscriptions/api/v1/#Invoices_Retrieve_a_subscription
func (s *API) GetInvoice(id string) (data InvoiceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "invoices",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/invoices/%s", s.ZohoTLD, id),
		Method:       zoho.HTTPGet,
		ResponseData: &InvoiceResponse{},
		Headers: map[string]string{
			ZohoSubscriptionsOriganizationID: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return InvoiceResponse{}, fmt.Errorf("Failed to retrieve invoice (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*InvoiceResponse); ok {
		return *v, nil
	}

	return InvoiceResponse{}, fmt.Errorf("Data retrieved was not 'InvoiceResponse'")
}

// AddAttachment attaches a file to an invoice
// https://www.zoho.com/subscriptions/api/v1/#Invoices_Add_attachment_to_an_invoice
func (s *API) AddAttachment(id, file string, canSendInEmail bool) (data AttachementResponse, err error) {

	endpoint := zoho.Endpoint{
		Name:         "invoices",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/invoices/%s/attachment", s.ZohoTLD, id),
		Method:       zoho.HTTPPost,
		ResponseData: &AttachementResponse{},
		RequestBody:  AttachmentRequest{CanSendInEmail: canSendInEmail},
		Headers: map[string]string{
			ZohoSubscriptionsOriganizationID: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AttachementResponse{}, fmt.Errorf("Failed to attach file to invoice (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*AttachementResponse); ok {
		return *v, nil
	}

	return AttachementResponse{}, fmt.Errorf("Data retrieved was not 'AttachementResponse'")
}

// EmailInvoice sends an invoice in email
// https://www.zoho.com/subscriptions/api/v1/#Invoices_Email_an_invoice
func (s *API) EmailInvoice(id string, request EmailInvoiceRequest) (data EmailInvoiceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "invoices",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/invoices/%s/email", s.ZohoTLD, id),
		Method:       zoho.HTTPPost,
		ResponseData: &EmailInvoiceResponse{},
		RequestBody:  request,
		Headers: map[string]string{
			ZohoSubscriptionsOriganizationID: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return EmailInvoiceResponse{}, fmt.Errorf("Failed to email invoice (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*EmailInvoiceResponse); ok {
		return *v, nil
	}

	return EmailInvoiceResponse{}, fmt.Errorf("Data retrieved was not 'EmailInvoiceResponse'")
}

// AddItems adds items to pending invoice
// https://www.zoho.com/subscriptions/api/v1/#Invoices_Add_items_to_a_pending_invoice
func (s *API) AddItems(id string, request AddItemsRequest) (data AddItemsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "invoices",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/invoices/%s/lineitems", s.ZohoTLD, id),
		Method:       zoho.HTTPPost,
		ResponseData: &AddItemsResponse{},
		RequestBody:  request,
		Headers: map[string]string{
			ZohoSubscriptionsOriganizationID: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AddItemsResponse{}, fmt.Errorf("Failed to add items to invoice (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*AddItemsResponse); ok {
		return *v, nil
	}

	return AddItemsResponse{}, fmt.Errorf("Data retrieved was not 'AddItemsResponse'")
}

// CollectChargeViaCreditCard collects charge via credit card
// https://www.zoho.com/subscriptions/api/v1/#Invoices_Collect_charge_via_credit_card
func (s *API) CollectChargeViaCreditCard(id string, request CollectChangeViaCreditCardRequest) (data CollectChangeViaCreditCardResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "invoices",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/invoices/%s/collect", s.ZohoTLD, id),
		Method:       zoho.HTTPPost,
		ResponseData: &CollectChangeViaCreditCardResponse{},
		RequestBody:  request,
		Headers: map[string]string{
			ZohoSubscriptionsOriganizationID: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CollectChangeViaCreditCardResponse{}, fmt.Errorf("Failed to collect charge via credit card (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*CollectChangeViaCreditCardResponse); ok {
		return *v, nil
	}

	return CollectChangeViaCreditCardResponse{}, fmt.Errorf("Data retrieved was not 'CollectChangeViaBankCreditCardResponse'")
}

// CollectChargeViaBankAccount collects charge via bank account
// https://www.zoho.com/subscriptions/api/v1/#Invoices_Collect_charge_via_bank_account
func (s *API) CollectChargeViaBankAccount(id string, request CollectChangeViaBankAccountRequest) (data CollectChangeViaBankAccountResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "invoices",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/invoices/%s/collect", s.ZohoTLD, id),
		Method:       zoho.HTTPPost,
		ResponseData: &CollectChangeViaBankAccountResponse{},
		RequestBody:  request,
		Headers: map[string]string{
			ZohoSubscriptionsOriganizationID: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CollectChangeViaBankAccountResponse{}, fmt.Errorf("Failed to collect charge via bank account (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*CollectChangeViaBankAccountResponse); ok {
		return *v, nil
	}

	return CollectChangeViaBankAccountResponse{}, fmt.Errorf("Data retrieved was not 'CollectChangeViaBankAccountResponse'")
}

type CollectChangeViaBankAccountRequest struct {
	AccountID string `json:"account_id"`
}

type CollectChangeViaBankAccountResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payment struct {
		PaymentID       string `json:"payment_id"`
		PaymentMode     string `json:"payment_mode"`
		Amount          int    `json:"amount"`
		AmountRefunded  int    `json:"amount_refunded"`
		BankCharges     int    `json:"bank_charges"`
		Date            string `json:"date"`
		Status          string `json:"status"`
		ReferenceNumber string `json:"reference_number"`
		DueDate         string `json:"due_date"`
		AmountDue       int    `json:"amount_due"`
		Description     string `json:"description"`
		CustomerID      string `json:"customer_id"`
		CustomerName    string `json:"customer_name"`
		Email           string `json:"email"`
		Autotransaction struct {
			AutotransactionID    string `json:"autotransaction_id"`
			PaymentGateway       string `json:"payment_gateway"`
			GatewayTransactionID string `json:"gateway_transaction_id"`
			GatewayErrorMessage  string `json:"gateway_error_message"`
			AccountID            string `json:"account_id"`
		} `json:"autotransaction"`
		Invoices []struct {
			InvoiceID     string `json:"invoice_id"`
			InvoiceNumber string `json:"invoice_number"`
			Date          string `json:"date"`
			InvoiceAmount int    `json:"invoice_amount"`
			AmountApplied int    `json:"amount_applied"`
			BalanceAmount int    `json:"balance_amount"`
		} `json:"invoices"`
		CurrencyCode   string `json:"currency_code"`
		CurrencySymbol string `json:"currency_symbol"`
		CustomFields   []struct {
			Index    int    `json:"index"`
			DataType string `json:"data_type"`
		} `json:"custom_fields"`
		CreatedTime string `json:"created_time"`
		UpdatedTime string `json:"updated_time"`
	} `json:"payment"`
}

type CollectChangeViaCreditCardRequest struct {
	CardID string `json:"card_id"`
}

type CollectChangeViaCreditCardResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payment struct {
		PaymentID       string `json:"payment_id"`
		PaymentMode     string `json:"payment_mode"`
		Amount          int    `json:"amount"`
		AmountRefunded  int    `json:"amount_refunded"`
		BankCharges     int    `json:"bank_charges"`
		Date            string `json:"date"`
		Status          string `json:"status"`
		ReferenceNumber string `json:"reference_number"`
		Description     string `json:"description"`
		CustomerID      string `json:"customer_id"`
		CustomerName    string `json:"customer_name"`
		Email           string `json:"email"`
		Autotransaction struct {
			AutotransactionID    string `json:"autotransaction_id"`
			PaymentGateway       string `json:"payment_gateway"`
			GatewayTransactionID string `json:"gateway_transaction_id"`
			GatewayErrorMessage  string `json:"gateway_error_message"`
			CardID               string `json:"card_id"`
			LastFourDigits       int    `json:"last_four_digits"`
			ExpiryMonth          int    `json:"expiry_month"`
			ExpiryYear           int    `json:"expiry_year"`
		} `json:"autotransaction"`
		Invoices []struct {
			InvoiceID     string `json:"invoice_id"`
			InvoiceNumber string `json:"invoice_number"`
			Date          string `json:"date"`
			InvoiceAmount int    `json:"invoice_amount"`
			AmountApplied int    `json:"amount_applied"`
			BalanceAmount int    `json:"balance_amount"`
		} `json:"invoices"`
		CurrencyCode   string `json:"currency_code"`
		CurrencySymbol string `json:"currency_symbol"`
		CustomFields   []struct {
			Index    int    `json:"index"`
			DataType string `json:"data_type"`
		} `json:"custom_fields"`
		CreatedTime string `json:"created_time"`
		UpdatedTime string `json:"updated_time"`
	} `json:"payment"`
}

type AddItemsRequest struct {
	InvoiceItems []struct {
		Code           string `json:"code"`
		ProductID      string `json:"product_id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		Price          int    `json:"price"`
		Quantity       int    `json:"quantity"`
		TaxID          string `json:"tax_id"`
		TaxExemptionID string `json:"tax_exemption_id"`
	} `json:"invoice_items"`
}

type AddItemsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Invoice struct {
		InvoiceID           string `json:"invoice_id"`
		Number              string `json:"number"`
		Status              string `json:"status"`
		InvoiceDate         string `json:"invoice_date"`
		DueDate             string `json:"due_date"`
		PaymentExpectedDate string `json:"payment_expected_date"`
		AchPaymentInitiated bool   `json:"ach_payment_initiated"`
		TransactionType     string `json:"transaction_type"`
		CustomerID          string `json:"customer_id"`
		CustomerName        string `json:"customer_name"`
		Email               string `json:"email"`
		InvoiceItems        []struct {
			ItemID           string        `json:"item_id"`
			Name             string        `json:"name"`
			Description      string        `json:"description"`
			Tags             []Tag         `json:"tags"`
			ItemCustomFields []CustomField `json:"item_custom_fields"`
			Code             string        `json:"code"`
			Price            int           `json:"price"`
			Quantity         int           `json:"quantity"`
			DiscountAmount   int           `json:"discount_amount"`
			ItemTotal        int           `json:"item_total"`
			TaxID            string        `json:"tax_id"`
			ProductType      string        `json:"product_type"`
			HsnOrSac         string        `json:"hsn_or_sac"`
			TaxExemptionID   string        `json:"tax_exemption_id"`
			TaxExemptionCode string        `json:"tax_exemption_code"`
		} `json:"invoice_items"`
		Coupons []struct {
			CouponCode     string `json:"coupon_code"`
			CouponName     string `json:"coupon_name"`
			DiscountAmount int    `json:"discount_amount"`
		} `json:"coupons"`
		Credits []struct {
			CreditnoteID      string `json:"creditnote_id"`
			CreditnotesNumber string `json:"creditnotes_number"`
			CreditedDate      string `json:"credited_date"`
			CreditedAmount    int    `json:"credited_amount"`
		} `json:"credits"`
		Total          int `json:"total"`
		PaymentMade    int `json:"payment_made"`
		Balance        int `json:"balance"`
		CreditsApplied int `json:"credits_applied"`
		WriteOffAmount int `json:"write_off_amount"`
		Payments       []struct {
			PaymentID            string `json:"payment_id"`
			PaymentMode          string `json:"payment_mode"`
			InvoicePaymentID     string `json:"invoice_payment_id"`
			GatewayTransactionID string `json:"gateway_transaction_id"`
			Description          string `json:"description"`
			Date                 string `json:"date"`
			ReferenceNumber      string `json:"reference_number"`
			Amount               int    `json:"amount"`
			BankCharges          int    `json:"bank_charges"`
			ExchangeRate         int    `json:"exchange_rate"`
		} `json:"payments"`
		CurrencyCode    string  `json:"currency_code"`
		CurrencySymbol  string  `json:"currency_symbol"`
		CreatedTime     string  `json:"created_time"`
		UpdatedTime     string  `json:"updated_time"`
		SalespersonID   string  `json:"salesperson_id"`
		SalespersonName string  `json:"salesperson_name"`
		InvoiceURL      string  `json:"invoice_url"`
		BillingAddress  Address `json:"billing_address"`
		ShippingAddress Address `json:"shipping_address"`
		Comments        []struct {
			CommentID       string `json:"comment_id"`
			Description     string `json:"description"`
			CommentedByID   string `json:"commented_by_id"`
			CommentedBy     string `json:"commented_by"`
			CommentType     string `json:"comment_type"`
			Time            string `json:"time"`
			OperationType   string `json:"operation_type"`
			TransactionID   string `json:"transaction_id"`
			TransactionType string `json:"transaction_type"`
		} `json:"comments"`
		CustomFields []CustomField `json:"custom_fields"`
	} `json:"invoice"`
}

type EmailInvoiceRequest struct {
	ToMailIds []string `json:"to_mail_ids"`
	CcMailIds []string `json:"cc_mail_ids"`
	Subject   string   `json:"subject"`
	Body      string   `json:"body"`
}

type EmailInvoiceResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type AttachmentRequest struct {
	CanSendInEmail bool `json:"can_send_in_mail"`
}

type AttachementResponse struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	Documents []struct {
		FileName          string `json:"file_name"`
		FileType          string `json:"file_type"`
		FileSize          int64  `json:"file_size"`
		FileSizeFormatted string `json:"file_size_formatted"`
		DocumentID        string `json:"document_id"`
		AttachmentOrder   int64  `json:"attachment_order"`
	} `json:"documents"`
}

type InvoicesResponse struct {
	Invoices []Invoice `json:"invoices"`
	Code     int64     `json:"code"`
	Message  string    `json:"message"`
}

type InvoiceResponse struct {
	Invoice Invoice `json:"invoice"`
	Code    int64   `json:"code"`
	Message string  `json:"message"`
}

type Invoice struct {
	InvoiceID            string        `json:"invoice_id,omitempty"`
	Number               string        `json:"number,omitempty"`
	Status               string        `json:"status,omitempty"`
	InvoiceDate          string        `json:"invoice_date,omitempty"`
	DueDate              string        `json:"due_date,omitempty"`
	CustomerID           string        `json:"customer_id,omitempty"`
	CustomerName         string        `json:"customer_name,omitempty"`
	Email                string        `json:"email,omitempty"`
	Balance              float64       `json:"balance,omitempty"`
	Total                float64       `json:"total,omitempty"`
	PaymentMade          float64       `json:"payment_made,omitempty"`
	CreditsApplied       float64       `json:"credits_applied,omitempty"`
	WriteOffAmount       float64       `json:"write_off_amount,omitempty"`
	CurrencyCode         string        `json:"currency_code,omitempty"`
	CurrencySymbol       string        `json:"currency_symbol,omitempty"`
	HasAttachment        bool          `json:"has_attachment,omitempty"`
	CreatedTime          string        `json:"created_time,omitempty"`
	UpdatedTime          string        `json:"updated_time,omitempty"`
	SalespersonID        string        `json:"salesperson_id,omitempty"`
	SalespersonName      string        `json:"salesperson_name,omitempty"`
	InvoiceUrl           string        `json:"invoice_url,omitempty"`
	PaymentExpectedDate  string        `json:"payment_expected_date,omitempty"`
	ArchPaymentInitiated bool          `json:"ach_payment_initiated,omitempty"`
	TransactionType      string        `json:"transaction_type,omitempty"`
	InvoiceItems         []InvoiceItem `json:"invoice_items,omitempty"`
	Coupons              []Coupon      `json:"coupons,omitempty"`
	Credits              []Credit      `json:"credits,omitempty"`
	Payments             []Payment     `json:"payments,omitempty"`
	BillingAddress       Address       `json:"billing_address,omitempty"`
	ShippingAddress      Address       `json:"shipping_address,omitempty"`
	Comments             []Comment     `json:"comments,omitempty"`
	CustomFields         []CustomField `json:"custom_fields,omitempty"`
	CanSendInEmail       bool          `json:"can_send_in_mail,omitempty"`
	Documents            []Document    `json:"documents,omitempty"`
}

type InvoiceItem struct {
	ItemID           string        `json:"item_id,omitempty"`
	Name             string        `json:"name,omitempty"`
	Description      string        `json:"description,omitempty"`
	Code             string        `json:"code,omitempty"`
	Tags             []Tag         `json:"tags,omitempty"`
	ItemCustomFields []CustomField `json:"item_custom_fields,omitempty"`
	Price            float64       `json:"price,omitempty"`
	Quantity         int64         `json:"quantity,omitempty"`
	DiscountAmount   float64       `json:"discount_amount,omitempty"`
	ItemTotal        float64       `json:"item_total,omitempty"`
	TaxID            string        `json:"tax_id,omitempty"`
	TaxExemptionID   string        `json:"tax_exemption_id,omitempty"`
	TaxExemptionCode string        `json:"tax_exemption_code,omitempty"`
}

type Coupon struct {
	CouponCode     string  `json:"coupon_code,omitempty"`
	CouponName     string  `json:"coupon_name,omitempty"`
	DiscountAmount float64 `json:"discount_amount,omitempty"`
}

type Credit struct {
	CreditnoteID      string  `json:"creditnote_id,omitempty"`
	CreditnotesNumber string  `json:"creditnotes_number,omitempty"`
	CreditedDate      string  `json:"credited_date,omitempty"`
	CreditedAmount    float64 `json:"credited_amount,omitempty"`
}

type Payment struct {
	PaymentID            string  `json:"payment_id,omitempty"`
	PaymentMode          string  `json:"payment_mode,omitempty"`
	InvoicePaymentID     string  `json:"invoice_payment_id,omitempty"`
	AmountRefunded       float64 `json:"amount_refunded,omitempty"`
	GatewayTransactionID string  `json:"gateway_transaction_id,omitempty"`
	Description          string  `json:"description,omitempty"`
	Date                 string  `json:"date,omitempty"`
	ReferenceNumber      string  `json:"reference_number,omitempty"`
	Amount               float64 `json:"amount,omitempty"`
	BankCharges          float64 `json:"bank_charges,omitempty"`
	ExchangeRate         float64 `json:"exchange_rate,omitempty"`
}

type Comment struct {
	CommentID       string `json:"comment_id,omitempty"`
	Description     string `json:"description,omitempty"`
	CommentedByID   string `json:"commented_by_id,omitempty"`
	CommentedBy     string `json:"commented_by,omitempty"`
	CommentType     string `json:"comment_type,omitempty"`
	Time            string `json:"time,omitempty"`
	OperationType   string `json:"operation_type,omitempty"`
	TransactionID   string `json:"transaction_id,omitempty"`
	TransactionType string `json:"transaction_type,omitempty"`
}

type Document struct {
	FileName          string `json:"file_name,omitempty"`
	FileType          string `json:"file_type,omitempty"`
	FileSize          int64  `json:"file_size,omitempty"`
	FileSizeFormatted string `json:"file_size_formatted,omitempty"`
	DocumentID        string `json:"document_id,omitempty"`
	AttachmentOrder   int64  `json:"attachment_order,omitempty"`
}

type Tag struct {
	TagID       string `json:"tag_id,omitempty"`
	TagOptionID string `json:"tag_option_id,omitempty"`
}
