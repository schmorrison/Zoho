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
			ZohoSubscriptionsOriganizationId: s.OrganizationId,
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
			ZohoSubscriptionsOriganizationId: s.OrganizationId,
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

type InvoicesResponse struct {
	Invoices []Invoice `json:"subscriptions"`
	Code     int64     `json:"code"`
	Message  string    `json:success"`
}

type InvoiceResponse struct {
	Invoice Invoice `json:"subscription"`
	Code    int64   `json:"code"`
	Message string  `json:success"`
}

type Invoice struct {
	InvoiceId            string        `json:"invoice_id"`
	Number               string        `json:"number"`
	Status               string        `json:"status"`
	InvoiceDate          string        `json:"invoice_date"`
	DueDate              string        `json:"due_date"`
	CustomerId           string        `json:"customer_id"`
	CustomerName         string        `json:"customer_name"`
	Email                string        `json:"email"`
	Balance              float64       `json:"balance"`
	Total                float64       `json:"total"`
	PaymentMade          float64       `json:"payment_made"`
	CreditsApplied       float64       `json:"credits_applied"`
	WriteOffAmount       float64       `json:"write_off_amount"`
	CurrencyCode         string        `json:"currency_code"`
	CurrencySymbol       string        `json:"currency_symbol"`
	HasAttachment        bool          `json:"has_attachment"`
	CreatedTime          string        `json:"created_time"`
	UpdatedTime          string        `json:"updated_time"`
	SalespersonId        string        `json:"salesperson_id"`
	SalespersonName      string        `json:"salesperson_name"`
	InvoiceUrl           string        `json:"invoice_url"`
	PaymentExpectedDate  string        `json:"payment_expected_date"`
	ArchPaymentInitiated bool          `json:"ach_payment_initiated"`
	TransactionType      string        `json:"transaction_type"`
	InvoiceItems         []InvoiceItem `json:"invoice_items"`
	Coupons              []Coupon      `json:"coupons"`
	Credits              []Credit      `json:"credits"`
	Payments             []Payment     `json:"payments"`
	BillingAddress       Address       `json:"billing_address"`
	ShippingAddress      Address       `json:"shipping_address"`
	Comments             []Comment     `json:"comments"`
	CustomFields         []CustomField `json:"custom_fields"`
	CanSendInEmail       bool          `json:"can_send_in_mail"`
	Documents            []Document    `json:"documents"`
}

type InvoiceItem struct {
	ItemId           string        `json:"item_id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Code             string        `json:"code"`
	Tags             []Tag         `json:"tags"`
	ItemCustomFields []CustomField `json:"item_custom_fields"`
	Price            float64       `json:"price"`
	Quantity         int64         `json:"quantity"`
	DiscountAmount   float64       `json:"discount_amount"`
	ItemTotal        int64         `json:"item_total"`
	TaxId            string        `json:"tax_id"`
	TaxExemptionId   string        `json:"tax_exemption_id"`
	TaxExemptionCode string        `json:"tax_exemption_code"`
}

type Coupon struct {
	CouponCode     string  `json:"coupon_code"`
	CouponName     string  `json:"coupon_name"`
	DiscountAmount float64 `json:"discount_amount"`
}

type Credit struct {
	CreditnoteId      string  `json:"creditnote_id"`
	CreditnotesNumber string  `json:"creditnotes_number"`
	CreditedDate      string  `json:"credited_date"`
	CreditedAmount    float64 `json:"credited_amount"`
}

type Payment struct {
	PaymentId            string  `json:"payment_id"`
	PaymentMode          string  `json:"payment_mode"`
	InvoicePaymentId     string  `json:"invoice_payment_id"`
	AmountRefunded       float64 `json:"amount_refunded"`
	GatewayTransactionId string  `json:"gateway_transaction_id"`
	Description          string  `json:"description"`
	Date                 string  `json:"date"`
	ReferenceNumber      string  `json:"reference_number"`
	Amount               float64 `json:"amount"`
	BankCharges          float64 `json:"bank_charges"`
	ExchangeRate         float64 `json:"exchange_rate"`
}

type Comment struct {
	CommentId       string `json:"comment_id"`
	Description     string `json:"description"`
	CommentedById   string `json:"commented_by_id"`
	CommentedBy     string `json:"commented_by"`
	CommentType     string `json:"comment_type"`
	Time            string `json:"time"`
	OperationType   string `json:"operation_type"`
	TransactionId   string `json:"transaction_id"`
	TransactionType string `json:"transaction_type"`
}

type Document struct {
	FileName          string `json:"file_name"`
	FileType          string `json:"file_type"`
	FileSize          int64  `json:"file_size"`
	FileSizeFormatted string `json:"file_size_formatted"`
	DocumentId        string `json:"document_id"`
	AttachmentOrder   int64  `json:"attachment_order"`
}

type Tag struct {
	TagId       string `json:"tag_id"`
	TagOptionId string `json:"tag_option_id"`
}
