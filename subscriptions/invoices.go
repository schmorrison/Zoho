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

//type InvoiceResponse map[string]interface{}

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
