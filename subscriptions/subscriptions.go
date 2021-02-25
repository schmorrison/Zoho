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
		return SubscriptionResponse{}, fmt.Errorf("Failed to retrieve user (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*SubscriptionResponse); ok {
		return *v, nil
	}

	return SubscriptionResponse{}, fmt.Errorf("Data retrieved was not 'SubscriptionResponse'")
}

//type SubscriptionsResponse map[string]interface{}

type SubscriptionsResponse struct {
	Subscriptions []Subscription `json:"subscriptions"`
	Code          int64          `json:"code"`
	Message       string         `json:success"`
}

type SubscriptionResponse struct {
	Subscription Subscription `json:"subscription"`
	Code         int64        `json:"code"`
	Message      string       `json:success"`
}

type Subscription struct {
	SubscriptionId      string  `json:"subscription_id"`
	Name                string  `json:"name"`
	Status              string  `json:"status"`
	Amount              float64 `json:"amount"`
	CreatedAt           string  `json:"created_at"`
	ActivatedAt         string  `json:"activated_at"`
	CurrentTermStartsAt string  `json:"current_term_starts_at"`
	CurrentTermEndsAt   string  `json:"current_term_ends_at"`
	LastBillingAt       string  `json:"last_billing_at"`
	NextBillingAt       string  `json:"next_billing_at"`
	ExpiresAt           string  `json:"expires_at"`
	Interval            int64   `json:"interval"`
	IntervalUnit        string  `json:"interval_unit"`
	AutoCollect         bool    `json:"auto_collect"`
	CreatedTime         string  `json:"created_time"`
	UpdatedTime         string  `json:"updated_time"`
	ReferenceId         string  `json:"reference_id"`
	SalespersonId       string  `json:"salesperson_id"`
	SalespersonName     string  `json:"salesperson_name"`
	ChildInvoiceId      string  `json:"child_invoice_id"`
	CurrencyCode        string  `json:"currency_code"`
	CurrencySymbol      string  `json:"currency_symbol"`
	EndOfTerm           bool    `json:"end_of_term"`
	ProductId           string  `json:"product_id"`
	ProductName         string  `json:"product_name"`
	Plan                struct {
		PlanCode        string  `json:"plan_code"`
		Name            string  `json:"name"`
		Quantity        int64   `json:"quantity"`
		Price           float64 `json:"price"`
		Discount        float64 `json:"discount"`
		Total           float64 `json:"total"`
		SetupFee        float64 `json:"setup_fee"`
		PlanDescription string  `json:"plan_description"`
		TaxId           string  `json:"tax_id"`
		TrialDays       int64   `json:"trial_days"`
	} `json:"plan"`
	Addons []struct {
		AddonCode        string  `json:"addon_code"`
		Name             string  `json:"name"`
		AddonDescription string  `json:"addon_description"`
		Quantity         int64   `json:"quantity"`
		Price            float64 `json:"price"`
		Discount         float64 `json:"discount"`
		Total            float64 `json:"total"`
		TaxId            string  `json:"tax_id"`
	} `json:"addons"`
	Coupon struct {
		CouponCode     string  `json:"coupon_code"`
		DiscountAmount float64 `json:"discount_amount"`
	} `json:"coupon"`
	Card struct {
		CardId         string `json:"card_id"`
		LastFourDigits string `json:"last_four_digits"`
		PaymentGateway string `json:"payment_gateway"`
		ExpiryMonth    int64  `json:"expiry_month"`
		ExpiryYear     int64  `json:"expiry_year"`
	} `json:"card"`
	PaymentTerms      int64    `json:"payment_terms"`
	PaymentTermsLabel string   `json:"payment_terms_label"`
	CanAddBankAccount bool     `json:"can_add_bank_account"`
	Customer          Customer `json:"customer"`
	CustomFields      []struct {
		Value    string `json:"value"`
		Label    string `json:"label"`
		DataType string `json:"data_type"`
	} `json:"custom_fields"`
	Contactpersons []struct {
		ContactpersonId string `json:"contactperson_id"`
	} `json:"contactpersons"`
	Notes []struct {
		NoteId        string `json:"note_id"`
		Description   string `json:"description"`
		CommentedBy   string `json:"commented_by"`
		CommentedTime string `json:"commented_time"`
	} `json:"notes"`
	PaymentGateways []struct {
		PaymentGateway string `json:"payment_gateway"`
	} `json:"payment_gateways"`
}

type Customer struct {
	SubscriptionId    string  `json:"customer_id"`
	Name              string  `json:"display_name"`
	Salutation        string  `json:"salutation"`
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Email             string  `json:"email"`
	CompanyName       string  `json:"company_name"`
	BillingAddress    Address `json:"billing_address"`
	ShippingAddress   Address `json:"shipping_address"`
	PaymentTerms      int64   `json:"payment_terms"`
	PaymentTermsLabel string  `json:"payment_terms_label"`
}

type Address struct {
	Attention string `json:"attention"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Zip       string `json:"zip"`
	Fax       string `json:"fax"`
}
