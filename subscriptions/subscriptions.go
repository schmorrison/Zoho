package subscriptions

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

type SubscriptionStatus string

// Proper names for Subscription statuses
const (
	SubscriptionStatusAll                  SubscriptionStatus = "SubscriptionStatus.All"
	SubscriptionStatusActive               SubscriptionStatus = "SubscriptionStatus.ACTIVE"
	SubscriptionStatusLive                 SubscriptionStatus = "SubscriptionStatus.LIVE"
	SubscriptionStatusFuture               SubscriptionStatus = "SubscriptionStatus.FUTURE"
	SubscriptionStatusTrial                SubscriptionStatus = "SubscriptionStatus.TRIAL"
	SubscriptionStatusPastDue              SubscriptionStatus = "SubscriptionStatus.PAST_DUE"
	SubscriptionStatusUnpaid               SubscriptionStatus = "SubscriptionStatus.UNPAID"
	SubscriptionStatusNonRenewing          SubscriptionStatus = "SubscriptionStatus.NON_RENEWING"
	SubscriptionStatusCancelledFromDunning SubscriptionStatus = "SubscriptionStatus.CANCELLED_FROM_DUNNING"
	SubscriptionStatusCancelled            SubscriptionStatus = "SubscriptionStatus.CANCELLED"
	SubscriptionStatusExpired              SubscriptionStatus = "SubscriptionStatus.EXPIRED"
	SubscriptionStatusTrialExpired         SubscriptionStatus = "SubscriptionStatus.TRIAL_EXPIRED"
	SubscriptionStatusCancelledLastMonth   SubscriptionStatus = "SubscriptionStatus.CANCELLED_LAST_MONTH"
	SubscriptionStatusCancelledThisMonth   SubscriptionStatus = "SubscriptionStatus.CANCELLED_THIS_MONTH"

	SubscriptionModeOnline  SubscriptionStatus = "SubscriptionMode.ONLINE"
	SubscriptionModeOffline SubscriptionStatus = "SubscriptionMode.OFFLINE"
)

// ListSubscriptions will return the list of subscriptions that match the given subscription status.
// https://www.zoho.com/subscriptions/api/v1/#Subscriptions_List_all_subscriptions
func (s *API) ListSubscriptions(status SubscriptionStatus) (data SubscriptionsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "subscriptions",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/subscriptions", s.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &SubscriptionsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": zoho.Parameter(status),
		},
		Headers: map[string]string{
			ZohoSubscriptionsOriganizationId: s.OrganizationId,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return SubscriptionsResponse{}, fmt.Errorf("Failed to retrieve subscriptions: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*SubscriptionsResponse); ok {
		return *v, nil
	}

	return SubscriptionsResponse{}, fmt.Errorf("Data retrieved was not 'SubscriptionsResponse'")
}

// GetSubscription will return the subscription specified by id
// https://www.zoho.com/subscriptions/api/v1/#Subscriptions_Retrieve_a_subscription
func (s *API) GetSubscription(id string) (data SubscriptionResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "subscriptions",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/subscriptions/%s", s.ZohoTLD, id),
		Method:       zoho.HTTPGet,
		ResponseData: &SubscriptionResponse{},
		Headers: map[string]string{
			ZohoSubscriptionsOriganizationId: s.OrganizationId,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return SubscriptionResponse{}, fmt.Errorf("Failed to retrieve subscription (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*SubscriptionResponse); ok {
		return *v, nil
	}

	return SubscriptionResponse{}, fmt.Errorf("Data retrieved was not 'SubscriptionResponse'")
}

// CreateSubscription creates new subscription
// https://www.zoho.com/subscriptions/api/v1/#Subscriptions_Create_a_subscription
func (s *API) CreateSubscription(request Subscription) (data SubscriptionResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "subscriptions",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/subscriptions", s.ZohoTLD),
		Method:       zoho.HTTPPost,
		ResponseData: &SubscriptionResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return SubscriptionResponse{}, fmt.Errorf("Failed to create subscription: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*SubscriptionResponse); ok {
		return *v, nil
	}

	return SubscriptionResponse{}, fmt.Errorf("Data returned was nil")
}

// UpdateSubscription will modify subscription by the data provided to request
// https://www.zoho.com/subscriptions/api/v1/#Subscriptions_Update_a_subscription
func (s *API) UpdateSubscription(request Subscription) (data SubscriptionResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "subscriptions",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/subscriptions/%s", s.ZohoTLD, request.SubscriptionId),
		Method:       zoho.HTTPPut,
		ResponseData: &SubscriptionResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return SubscriptionResponse{}, fmt.Errorf("Failed to update subscription: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*SubscriptionResponse); ok {
		return *v, nil
	}

	return SubscriptionResponse{}, fmt.Errorf("Data returned was nil")
}

type SubscriptionsResponse struct {
	Subscriptions []Subscription `json:"subscriptions"`
	Code          int64          `json:"code"`
	Message       string         `json:"success"`
}

type SubscriptionResponse struct {
	Subscription Subscription `json:"subscription"`
	Code         int64        `json:"code"`
	Message      string       `json:"success"`
}

type Subscription struct {
	SubscriptionId      string  `json:"subscription_id,omitempty"`
	Name                string  `json:"name,omitempty"`
	Status              string  `json:"status,omitempty"`
	Amount              float64 `json:"amount,omitempty"`
	CreatedAt           string  `json:"created_at,omitempty"`
	ActivatedAt         string  `json:"activated_at,omitempty"`
	CurrentTermStartsAt string  `json:"current_term_starts_at,omitempty"`
	CurrentTermEndsAt   string  `json:"current_term_ends_at,omitempty"`
	LastBillingAt       string  `json:"last_billing_at,omitempty"`
	NextBillingAt       string  `json:"next_billing_at,omitempty"`
	ExpiresAt           string  `json:"expires_at,omitempty"`
	Interval            int64   `json:"interval,omitempty"`
	IntervalUnit        string  `json:"interval_unit,omitempty"`
	AutoCollect         bool    `json:"auto_collect,omitempty"`
	CreatedTime         string  `json:"created_time,omitempty"`
	UpdatedTime         string  `json:"updated_time,omitempty"`
	ReferenceId         string  `json:"reference_id,omitempty"`
	SalespersonId       string  `json:"salesperson_id,omitempty"`
	SalespersonName     string  `json:"salesperson_name,omitempty"`
	ChildInvoiceId      string  `json:"child_invoice_id,omitempty"`
	CurrencyCode        string  `json:"currency_code,omitempty"`
	CurrencySymbol      string  `json:"currency_symbol,omitempty"`
	EndOfTerm           bool    `json:"end_of_term,omitempty"`
	ProductId           string  `json:"product_id,omitempty"`
	ProductName         string  `json:"product_name,omitempty"`
	Plan                struct {
		PlanCode        string  `json:"plan_code,omitempty"`
		Name            string  `json:"name,omitempty"`
		Quantity        int64   `json:"quantity,omitempty"`
		Price           float64 `json:"price,omitempty"`
		Discount        float64 `json:"discount,omitempty"`
		Total           float64 `json:"total,omitempty"`
		SetupFee        float64 `json:"setup_fee,omitempty"`
		PlanDescription string  `json:"plan_description,omitempty"`
		TaxId           string  `json:"tax_id,omitempty"`
		TrialDays       int64   `json:"trial_days,omitempty"`
	} `json:"plan,omitempty"`
	Addons []struct {
		AddonCode        string  `json:"addon_code,omitempty"`
		Name             string  `json:"name,omitempty"`
		AddonDescription string  `json:"addon_description,omitempty"`
		Quantity         int64   `json:"quantity,omitempty"`
		Price            float64 `json:"price,omitempty"`
		Discount         float64 `json:"discount,omitempty"`
		Total            float64 `json:"total,omitempty"`
		TaxId            string  `json:"tax_id,omitempty"`
	} `json:"addons,omitempty"`
	Coupon struct {
		CouponCode     string  `json:"coupon_code,omitempty"`
		DiscountAmount float64 `json:"discount_amount,omitempty"`
	} `json:"coupon,omitempty"`
	Card struct {
		CardId         string `json:"card_id,omitempty"`
		LastFourDigits string `json:"last_four_digits,omitempty"`
		PaymentGateway string `json:"payment_gateway,omitempty"`
		ExpiryMonth    int64  `json:"expiry_month,omitempty"`
		ExpiryYear     int64  `json:"expiry_year,omitempty"`
	} `json:"card,omitempty"`
	PaymentTerms      int64         `json:"payment_terms,omitempty"`
	PaymentTermsLabel string        `json:"payment_terms_label,omitempty"`
	CanAddBankAccount bool          `json:"can_add_bank_account,omitempty"`
	Customer          Customer      `json:"customer,omitempty"`
	CustomFields      []CustomField `json:"custom_fields,omitempty"`
	Contactpersons    []struct {
		ContactpersonId string `json:"contactperson_id,omitempty"`
	} `json:"contactpersons,omitempty"`
	Notes []struct {
		NoteId        string `json:"note_id,omitempty"`
		Description   string `json:"description,omitempty"`
		CommentedBy   string `json:"commented_by,omitempty"`
		CommentedTime string `json:"commented_time,omitempty"`
	} `json:"notes,omitempty"`
	PaymentGateways []struct {
		PaymentGateway string `json:"payment_gateway,omitempty"`
	} `json:"payment_gateways,omitempty"`
}

type Customer struct {
	SubscriptionId    string  `json:"customer_id,omitempty"`
	Name              string  `json:"display_name,omitempty"`
	Salutation        string  `json:"salutation,omitempty"`
	FirstName         string  `json:"first_name,omitempty"`
	LastName          string  `json:"last_name,omitempty"`
	Email             string  `json:"email,omitempty"`
	CompanyName       string  `json:"company_name,omitempty"`
	BillingAddress    Address `json:"billing_address,omitempty"`
	ShippingAddress   Address `json:"shipping_address,omitempty"`
	PaymentTerms      int64   `json:"payment_terms,omitempty"`
	PaymentTermsLabel string  `json:"payment_terms_label,omitempty"`
}

type Address struct {
	Attention string `json:"attention,omitempty"`
	Street    string `json:"street,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Country   string `json:"country,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Fax       string `json:"fax,omitempty"`
}

type CustomField struct {
	Index    int64  `json:"index,omitempty"`
	Value    string `json:"value,omitempty"`
	Label    string `json:"label,omitempty"`
	DataType string `json:"data_type,omitempty"`
}
