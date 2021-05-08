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
	if status == "" {
		status = SubscriptionStatusAll
	}
	endpoint := zoho.Endpoint{
		Name:         "subscriptions",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/subscriptions", s.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &SubscriptionsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": zoho.Parameter(status),
		},
		Headers: map[string]string{
			ZohoSubscriptionsEndpointHeader: s.OrganizationID,
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
			ZohoSubscriptionsEndpointHeader: s.OrganizationID,
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
func (s *API) CreateSubscription(request SubscriptionCreate) (data SubscriptionResponse, err error) {
	if request.CustomerID == "" {
		if request.Customer.DisplayName == "" || request.Customer.Email == "" {
			err = fmt.Errorf("CustomerID is a required field if subscription is created for existen customer. For new customer Customer.DisplayName and Customer.Email fields are required")
			return
		}
	}

	endpoint := zoho.Endpoint{
		Name:         "subscriptions",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/subscriptions", s.ZohoTLD),
		Method:       zoho.HTTPPost,
		ResponseData: &SubscriptionResponse{},
		RequestBody:  request,
		Headers: map[string]string{
			ZohoSubscriptionsEndpointHeader: s.OrganizationID,
		},
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
func (s *API) UpdateSubscription(request SubscriptionUpdate, ID string) (data SubscriptionResponse, err error) {
	if request.Plan.PlanCode == "" {
		return SubscriptionResponse{}, fmt.Errorf("Plan.PlanCode is a required field")
	}

	endpoint := zoho.Endpoint{
		Name:         "subscriptions",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/subscriptions/%s", s.ZohoTLD, ID),
		Method:       zoho.HTTPPut,
		ResponseData: &SubscriptionResponse{},
		RequestBody:  request,
		Headers: map[string]string{
			ZohoSubscriptionsEndpointHeader: s.OrganizationID,
		},
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

// DeleteSubscription will delete subscription by id
// https://www.zoho.com/subscriptions/api/v1/#Subscriptions_Delete_a_subscription
func (s *API) DeleteSubscription(ID string) (data SubscriptionDeleteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "subscriptions",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/subscriptions/%s", s.ZohoTLD, ID),
		Method:       zoho.HTTPDelete,
		ResponseData: &SubscriptionDeleteResponse{},
		Headers: map[string]string{
			ZohoSubscriptionsEndpointHeader: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return SubscriptionDeleteResponse{}, fmt.Errorf("Failed to delete subscription %s: %s", ID, err)
	}

	if v, ok := endpoint.ResponseData.(*SubscriptionDeleteResponse); ok {
		return *v, nil
	}

	return SubscriptionDeleteResponse{}, fmt.Errorf("Data returned was nil")
}

// AddChargeToSubscription charges a one-time amount for the subscription
// https://www.zoho.com/subscriptions/api/v1/#Subscriptions_Update_a_subscription
func (s *API) AddChargeToSubscription(request SubscriptionAddCharge, ID string) (data AddChargeResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "subscriptions",
		URL:          fmt.Sprintf("https://subscriptions.zoho.%s/api/v1/subscriptions/%s/charge", s.ZohoTLD, ID),
		Method:       zoho.HTTPPost,
		ResponseData: &AddChargeResponse{},
		RequestBody:  request,
		Headers: map[string]string{
			ZohoSubscriptionsEndpointHeader: s.OrganizationID,
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AddChargeResponse{}, fmt.Errorf("Failed to charge subscription: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*AddChargeResponse); ok {
		return *v, nil
	}

	return AddChargeResponse{}, fmt.Errorf("Data returned was nil")
}

type AddChargeResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Invoice struct {
		AchPaymentInitiated     bool    `json:"ach_payment_initiated"`
		Adjustment              float64 `json:"adjustment"`
		AdjustmentDescription   string  `json:"adjustment_description"`
		AllowPartialPayments    bool    `json:"allow_partial_payments"`
		ApproverID              string  `json:"approver_id"`
		AutoRemindersConfigured bool    `json:"auto_reminders_configured"`
		Balance                 float64 `json:"balance"`
		BcyAdjustment           float64 `json:"bcy_adjustment"`
		BcyDiscountTotal        float64 `json:"bcy_discount_total"`
		BcyShippingCharge       float64 `json:"bcy_shipping_charge"`
		BcySubTotal             float64 `json:"bcy_sub_total"`
		BcyTaxTotal             float64 `json:"bcy_tax_total"`
		BcyTotal                float64 `json:"bcy_total"`
		BillingAddress          struct {
			Address   string `json:"address"`
			Attention string `json:"attention"`
			City      string `json:"city"`
			Country   string `json:"country"`
			Fax       string `json:"fax"`
			Phone     string `json:"phone"`
			State     string `json:"state"`
			Street    string `json:"street"`
			Street2   string `json:"street2"`
			Zip       string `json:"zip"`
		} `json:"billing_address"`
		CanEditItems       bool   `json:"can_edit_items"`
		CanSendInMail      bool   `json:"can_send_in_mail"`
		CanSendInvoiceSms  bool   `json:"can_send_invoice_sms"`
		CanSkipPaymentInfo bool   `json:"can_skip_payment_info"`
		ClientViewedTime   string `json:"client_viewed_time"`
		Contactpersons     []struct {
			ContactpersonID string `json:"contactperson_id"`
			Email           string `json:"email"`
			Mobile          string `json:"mobile"`
			Phone           string `json:"phone"`
			ZcrmContactID   string `json:"zcrm_contact_id"`
		} `json:"contactpersons"`
		Coupons         []Coupon      `json:"coupons"`
		CreatedByID     string        `json:"created_by_id"`
		CreatedDate     string        `json:"created_date"`
		CreatedTime     string        `json:"created_time"`
		Credits         []interface{} `json:"credits"`
		CreditsApplied  float64       `json:"credits_applied"`
		CurrencyCode    string        `json:"currency_code"`
		CurrencyID      string        `json:"currency_id"`
		CurrencySymbol  string        `json:"currency_symbol"`
		CustomFieldHash struct {
		} `json:"custom_field_hash"`
		CustomFields            []CustomField `json:"custom_fields"`
		CustomerCustomFieldHash struct {
		} `json:"customer_custom_field_hash"`
		CustomerCustomFields        []CustomField `json:"customer_custom_fields"`
		CustomerID                  string        `json:"customer_id"`
		CustomerName                string        `json:"customer_name"`
		Date                        string        `json:"date"`
		DiscountPercent             float64       `json:"discount_percent"`
		DiscountTotal               float64       `json:"discount_total"`
		Documents                   []interface{} `json:"documents"`
		DueDate                     string        `json:"due_date"`
		Email                       string        `json:"email"`
		ExchangeRate                float64       `json:"exchange_rate"`
		InprocessTransactionPresent bool          `json:"inprocess_transaction_present"`
		InvoiceDate                 string        `json:"invoice_date"`
		InvoiceID                   string        `json:"invoice_id"`
		InvoiceItems                []struct {
			AccountID        string        `json:"account_id"`
			AccountName      string        `json:"account_name"`
			Code             string        `json:"code"`
			Description      string        `json:"description"`
			DiscountAmount   float64       `json:"discount_amount"`
			ItemCustomFields []CustomField `json:"item_custom_fields"`
			ItemID           string        `json:"item_id"`
			ItemTotal        float64       `json:"item_total"`
			Name             string        `json:"name"`
			Price            float64       `json:"price"`
			ProductID        string        `json:"product_id"`
			ProductType      string        `json:"product_type"`
			Quantity         float64       `json:"quantity"`
			Tags             []Tag         `json:"tags"`
			TaxID            string        `json:"tax_id"`
			TaxName          string        `json:"tax_name"`
			TaxPercentage    float64       `json:"tax_percentage"`
			TaxType          string        `json:"tax_type"`
			Unit             string        `json:"unit"`
		} `json:"invoice_items"`
		InvoiceNumber          string `json:"invoice_number"`
		InvoiceURL             string `json:"invoice_url"`
		IsInclusiveTax         bool   `json:"is_inclusive_tax"`
		IsReverseChargeApplied bool   `json:"is_reverse_charge_applied"`
		IsViewedByClient       bool   `json:"is_viewed_by_client"`
		LastModifiedByID       string `json:"last_modified_by_id"`
		Notes                  string `json:"notes"`
		Number                 string `json:"number"`
		PageWidth              string `json:"page_width"`
		PaymentExpectedDate    string `json:"payment_expected_date"`
		PaymentGateways        []struct {
			PaymentGateway string `json:"payment_gateway"`
		} `json:"payment_gateways"`
		PaymentMade            float64 `json:"payment_made"`
		PaymentReminderEnabled bool    `json:"payment_reminder_enabled"`
		PaymentTerms           int64   `json:"payment_terms"`
		PaymentTermsLabel      string  `json:"payment_terms_label"`
		Payments               []struct {
			Amount               float64 `json:"amount"`
			AmountRefunded       float64 `json:"amount_refunded"`
			BankCharges          float64 `json:"bank_charges"`
			CardType             string  `json:"card_type"`
			Date                 string  `json:"date"`
			Description          string  `json:"description"`
			ExchangeRate         float64 `json:"exchange_rate"`
			GatewayTransactionID string  `json:"gateway_transaction_id"`
			InvoicePaymentID     string  `json:"invoice_payment_id"`
			LastFourDigits       string  `json:"last_four_digits"`
			PaymentID            string  `json:"payment_id"`
			PaymentMode          string  `json:"payment_mode"`
			ReferenceNumber      string  `json:"reference_number"`
			SettlementStatus     string  `json:"settlement_status"`
		} `json:"payments"`
		PricePrecision  int64  `json:"price_precision"`
		PricebookID     string `json:"pricebook_id"`
		ReferenceID     string `json:"reference_id"`
		ReferenceNumber string `json:"reference_number"`
		SalespersonID   string `json:"salesperson_id"`
		SalespersonName string `json:"salesperson_name"`
		ShippingAddress struct {
			Address   string `json:"address"`
			Attention string `json:"attention"`
			City      string `json:"city"`
			Country   string `json:"country"`
			Fax       string `json:"fax"`
			Phone     string `json:"phone"`
			State     string `json:"state"`
			Street    string `json:"street"`
			Street2   string `json:"street2"`
			Zip       string `json:"zip"`
		} `json:"shipping_address"`
		ShippingCharge                        float64 `json:"shipping_charge"`
		ShippingChargeExclusiveOfTax          float64 `json:"shipping_charge_exclusive_of_tax"`
		ShippingChargeExclusiveOfTaxFormatted string  `json:"shipping_charge_exclusive_of_tax_formatted"`
		ShippingChargeInclusiveOfTax          float64 `json:"shipping_charge_inclusive_of_tax"`
		ShippingChargeInclusiveOfTaxFormatted string  `json:"shipping_charge_inclusive_of_tax_formatted"`
		ShippingChargeTax                     string  `json:"shipping_charge_tax"`
		ShippingChargeTaxFormatted            string  `json:"shipping_charge_tax_formatted"`
		ShippingChargeTaxID                   string  `json:"shipping_charge_tax_id"`
		ShippingChargeTaxName                 string  `json:"shipping_charge_tax_name"`
		ShippingChargeTaxPercentage           string  `json:"shipping_charge_tax_percentage"`
		ShippingChargeTaxType                 string  `json:"shipping_charge_tax_type"`
		Status                                string  `json:"status"`
		StopReminderUntilPaymentExpectedDate  bool    `json:"stop_reminder_until_payment_expected_date"`
		SubTotal                              float64 `json:"sub_total"`
		SubmitterID                           string  `json:"submitter_id"`
		Subscriptions                         []struct {
			SubscriptionID string `json:"subscription_id"`
		} `json:"subscriptions"`
		TaxRounding                   string        `json:"tax_rounding"`
		TaxTotal                      float64       `json:"tax_total"`
		Taxes                         []interface{} `json:"taxes"`
		TemplateID                    string        `json:"template_id"`
		TemplateName                  string        `json:"template_name"`
		TemplateType                  string        `json:"template_type"`
		Terms                         string        `json:"terms"`
		Total                         float64       `json:"total"`
		TransactionType               string        `json:"transaction_type"`
		UnbilledChargesID             string        `json:"unbilled_charges_id"`
		UnusedCreditsReceivableAmount float64       `json:"unused_credits_receivable_amount"`
		UpdatedTime                   string        `json:"updated_time"`
		VatTreatment                  string        `json:"vat_treatment"`
		WriteOffAmount                float64       `json:"write_off_amount"`
		ZcrmPotentialID               string        `json:"zcrm_potential_id"`
	} `json:"invoice"`
	UnbilledCharge struct {
		Balance        float64 `json:"balance"`
		BillingAddress struct {
			Attention string `json:"attention"`
			City      string `json:"city"`
			Country   string `json:"country"`
			Fax       string `json:"fax"`
			Phone     string `json:"phone"`
			State     string `json:"state"`
			Street    string `json:"street"`
			Street2   string `json:"street2"`
			Zip       string `json:"zip"`
		} `json:"billing_address"`
		Coupons         []Coupon `json:"coupons"`
		CreatedTime     string   `json:"created_time"`
		CurrencyCode    string   `json:"currency_code"`
		CurrencySymbol  string   `json:"currency_symbol"`
		CustomFieldHash struct {
		} `json:"custom_field_hash"`
		CustomFields            []CustomField `json:"custom_fields"`
		CustomerCustomFieldHash struct {
		} `json:"customer_custom_field_hash"`
		CustomerCustomFields   []CustomField `json:"customer_custom_fields"`
		CustomerID             string        `json:"customer_id"`
		CustomerName           string        `json:"customer_name"`
		Email                  string        `json:"email"`
		IsInclusiveTax         bool          `json:"is_inclusive_tax"`
		IsOnlineSubscription   bool          `json:"is_online_subscription"`
		IsReverseChargeApplied bool          `json:"is_reverse_charge_applied"`
		Number                 string        `json:"number"`
		PricePrecision         int64         `json:"price_precision"`
		SalespersonID          string        `json:"salesperson_id"`
		SalespersonName        string        `json:"salesperson_name"`
		ShippingAddress        struct {
			Attention string `json:"attention"`
			City      string `json:"city"`
			Country   string `json:"country"`
			Fax       string `json:"fax"`
			Phone     string `json:"phone"`
			State     string `json:"state"`
			Street    string `json:"street"`
			Street2   string `json:"street2"`
			Zip       string `json:"zip"`
		} `json:"shipping_address"`
		Status              string        `json:"status"`
		SubTotal            float64       `json:"sub_total"`
		SubscriptionID      string        `json:"subscription_id"`
		TaxRounding         string        `json:"tax_rounding"`
		TaxTotal            float64       `json:"tax_total"`
		Taxes               []interface{} `json:"taxes"`
		Total               float64       `json:"total"`
		TransactionType     string        `json:"transaction_type"`
		UnbilledChargeDate  string        `json:"unbilled_charge_date"`
		UnbilledChargeID    string        `json:"unbilled_charge_id"`
		UnbilledChargeItems []struct {
			AccountID            string  `json:"account_id"`
			Code                 string  `json:"code"`
			Description          string  `json:"description"`
			DiscountAmount       float64 `json:"discount_amount"`
			ItemTotal            float64 `json:"item_total"`
			Name                 string  `json:"name"`
			Price                float64 `json:"price"`
			ProductID            string  `json:"product_id"`
			ProductType          string  `json:"product_type"`
			Quantity             float64 `json:"quantity"`
			TaxID                string  `json:"tax_id"`
			TaxName              string  `json:"tax_name"`
			TaxPercentage        float64 `json:"tax_percentage"`
			TaxType              string  `json:"tax_type"`
			UnbilledChargeItemID string  `json:"unbilled_charge_item_id"`
		} `json:"unbilled_charge_items"`
		UnusedCreditsReceivableAmount float64 `json:"unused_credits_receivable_amount"`
		UpdatedTime                   string  `json:"updated_time"`
	} `json:"unbilled_charge"`
}

type SubscriptionsResponse struct {
	Subscriptions []Subscription `json:"subscriptions"`
	Code          int64          `json:"code"`
	Message       string         `json:"message"`
}

type SubscriptionResponse struct {
	Subscription Subscription `json:"subscription"`
	Code         int64        `json:"code"`
	Message      string       `json:"message"`
}

type SubscriptionDeleteResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type SubscriptionCreate struct {
	AddToUnbilledCharges bool   `json:"add_to_unbilled_charges,omitempty"`
	CustomerID           string `json:"customer_id,omitempty"`
	Customer             struct {
		DisplayName       string        `json:"display_name,omitempty"`
		Salutation        string        `json:"salutation,omitempty"`
		FirstName         string        `json:"first_name,omitempty"`
		LastName          string        `json:"last_name,omitempty"`
		Email             string        `json:"email,omitempty"`
		CompanyName       string        `json:"company_name,omitempty"`
		BillingAddress    Address       `json:"billing_address,omitempty"`
		ShippingAddress   Address       `json:"shipping_address,omitempty"`
		PaymentTerms      int64         `json:"payment_terms,omitempty"`
		PaymentTermsLabel string        `json:"payment_terms_label,omitempty"`
		CustomFields      []CustomField `json:"custom_fields,omitempty"`
		PlaceOfContact    string        `json:"place_of_contact,omitempty"`
		GstNo             string        `json:"gst_no,omitempty"`
		GstTreatment      string        `json:"gst_treatment,omitempty"`
		VatTreatment      string        `json:"vat_treatment,omitempty"`
		VatRegNo          int64         `json:"vat_reg_no,omitempty"`
		CountryCode       string        `json:"country_code,omitempty"`
		IsTaxable         bool          `json:"is_taxable,omitempty"`
		TaxID             string        `json:"tax_id,omitempty"`
		TaxAuthorityID    string        `json:"tax_authority_id,omitempty"`
		TaxAuthorityName  string        `json:"tax_authority_name,omitempty"`
		TaxExemptionID    string        `json:"tax_exemption_id,omitempty"`
		TaxExemptionCode  string        `json:"tax_exemption_code,omitempty"`
	} `json:"customer,omitempty"`
	PaymentTerms      int64         `json:"payment_terms,omitempty"`
	PaymentTermsLabel string        `json:"payment_terms_label,omitempty"`
	CustomFields      []CustomField `json:"custom_fields,omitempty"`
	Contactpersons    []struct {
		ContactpersonID string `json:"contactperson_id,omitempty"`
	} `json:"contactpersons,omitempty"`
	StartsAt      string `json:"starts_at,omitempty"`
	ExchangeRate  int64  `json:"exchange_rate,omitempty"`
	PlaceOfSupply string `json:"place_of_supply,omitempty"`
	Plan          struct {
		PlanCode                 string        `json:"plan_code,omitempty"`
		PlanDescription          string        `json:"plan_description,omitempty"`
		Price                    float64       `json:"price,omitempty"`
		SetupFee                 float64       `json:"setup_fee,omitempty"`
		SetupFeeTaxID            string        `json:"setup_fee_tax_id,omitempty"`
		Tags                     []Tag         `json:"tags,omitempty"`
		ItemCustomFields         []CustomField `json:"item_custom_fields,omitempty"`
		Quantity                 float64       `json:"quantity,omitempty"`
		TaxExemptionID           string        `json:"tax_exemption_id,omitempty"`
		TaxExemptionCode         string        `json:"tax_exemption_code,omitempty"`
		SetupFeeTaxExemptionID   string        `json:"setup_fee_tax_exemption_id,omitempty"`
		SetupFeeTaxExemptionCode string        `json:"setup_fee_tax_exemption_code,omitempty"`
		ExcludeTrial             bool          `json:"exclude_trial,omitempty"`
		ExcludeSetupFee          bool          `json:"exclude_setup_fee,omitempty"`
		BillingCycles            int64         `json:"billing_cycles,omitempty"`
		TrialDays                int64         `json:"trial_days,omitempty"`
	} `json:"plan,omitempty"`
	Addons []struct {
		AddonCode        string        `json:"addon_code,omitempty"`
		AddonDescription string        `json:"addon_description,omitempty"`
		Price            float64       `json:"price,omitempty"`
		Tags             []Tag         `json:"tags,omitempty"`
		ItemCustomFields []CustomField `json:"item_custom_fields,omitempty"`
		TaxExemptionID   string        `json:"tax_exemption_id,omitempty"`
		TaxExemptionCode string        `json:"tax_exemption_code,omitempty"`
	} `json:"addons,omitempty"`
	CouponCode             string           `json:"coupon_code,omitempty"`
	AutoCollect            bool             `json:"auto_collect,omitempty"`
	ReferenceID            string           `json:"reference_id,omitempty"`
	SalespersonName        string           `json:"salesperson_name,omitempty"`
	PaymentGateways        []PaymentGateway `json:"payment_gateways,omitempty"`
	CreateBackdatedInvoice bool             `json:"create_backdated_invoice,omitempty"`
	TemplateID             int64            `json:"template_id,omitempty"`
}

type SubscriptionUpdate struct {
	ExchangeRate float64 `json:"exchange_rate,omitempty,omitempty"`
	Plan         struct {
		PlanCode                 string        `json:"plan_code,omitempty"`
		PlanDescription          string        `json:"plan_description,omitempty"`
		Price                    float64       `json:"price,omitempty"`
		SetupFee                 float64       `json:"setup_fee,omitempty"`
		Quantity                 float64       `json:"quantity,omitempty"`
		Tags                     []Tag         `json:"tags,omitempty"`
		ItemCustomFields         []CustomField `json:"item_custom_fields,omitempty"`
		TaxID                    string        `json:"tax_id,omitempty"`
		TaxExemptionID           string        `json:"tax_exemption_id,omitempty"`
		TaxExemptionCode         string        `json:"tax_exemption_code,omitempty"`
		SetupFeeTaxExemptionID   string        `json:"setup_fee_tax_exemption_id,omitempty"`
		SetupFeeTaxExemptionCode string        `json:"setup_fee_tax_exemption_code,omitempty"`
		ExcludeTrial             bool          `json:"exclude_trial,omitempty"`
		ExcludeSetupFee          bool          `json:"exclude_setup_fee,omitempty"`
		BillingCycles            int64         `json:"billing_cycles,omitempty"`
		TrialDays                int64         `json:"trial_days,omitempty"`
	} `json:"plan,omitempty"`
	Addons []struct {
		AddonCode        string        `json:"addon_code,omitempty"`
		AddonDescription string        `json:"addon_description,omitempty"`
		Price            float64       `json:"price,omitempty"`
		Tags             []Tag         `json:"tags,omitempty"`
		ItemCustomFields []CustomField `json:"item_custom_fields,omitempty"`
		TaxExemptionID   string        `json:"tax_exemption_id,omitempty"`
		TaxExemptionCode string        `json:"tax_exemption_code,omitempty"`
	} `json:"addons,omitempty"`
	AutoCollect     bool   `json:"auto_collect,omitempty"`
	CouponCode      string `json:"coupon_code,omitempty"`
	ReferenceID     string `json:"reference_id,omitempty"`
	SalespersonID   string `json:"salesperson_id,omitempty"`
	SalespersonName string `json:"salesperson_name,omitempty"`
	EndOfTerm       bool   `json:"end_of_term,omitempty"`
	Contactpersons  []struct {
		ContactpersonID string `json:"contactperson_id,omitempty"`
	} `json:"contactpersons,omitempty"`
	PaymentTerms      int64            `json:"payment_terms,omitempty"`
	PaymentTermsLabel string           `json:"payment_terms_label,omitempty"`
	PaymentGateways   []PaymentGateway `json:"payment_gateways,omitempty"`
	CustomFields      []CustomField    `json:"custom_fields,omitempty"`
	TemplateID        int64            `json:"template_id,omitempty"`
}

type SubscriptionAddCharge struct {
	Amount               float64       `json:"amount,omitempty"`
	Description          string        `json:"description,omitempty"`
	Tags                 []Tag         `json:"tags,omitempty"`
	ItemCustomFields     []CustomField `json:"item_custom_fields,omitempty"`
	AccountID            string        `json:"account_id,omitempty"`
	AddToUnbilledCharges bool          `json:"add_to_unbilled_charges,omitempty"`
}

type Subscription struct {
	SubscriptionID      string  `json:"subscription_id,omitempty"`
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
	ReferenceID         string  `json:"reference_id,omitempty"`
	SalespersonID       string  `json:"salesperson_id,omitempty"`
	SalespersonName     string  `json:"salesperson_name,omitempty"`
	ChildInvoiceID      string  `json:"child_invoice_id,omitempty"`
	CurrencyCode        string  `json:"currency_code,omitempty"`
	CurrencySymbol      string  `json:"currency_symbol,omitempty"`
	EndOfTerm           bool    `json:"end_of_term,omitempty"`
	ProductID           string  `json:"product_id,omitempty"`
	ProductName         string  `json:"product_name,omitempty"`
	Plan                Plan    `json:"plan,omitempty"`
	Addons              []Addon `json:"addons,omitempty"`
	Coupon              struct {
		CouponCode     string  `json:"coupon_code,omitempty"`
		DiscountAmount float64 `json:"discount_amount,omitempty"`
	} `json:"coupon,omitempty"`
	Card struct {
		CardID         string `json:"card_id,omitempty"`
		LastFourDigits string `json:"last_four_digits,omitempty"`
		PaymentGateway string `json:"payment_gateway,omitempty"`
		ExpiryMonth    int64  `json:"expiry_month,omitempty"`
		ExpiryYear     int64  `json:"expiry_year,omitempty"`
	} `json:"card,omitempty"`
	PaymentTerms      int64           `json:"payment_terms,omitempty"`
	PaymentTermsLabel string          `json:"payment_terms_label,omitempty"`
	CanAddBankAccount bool            `json:"can_add_bank_account,omitempty"`
	Customer          Customer        `json:"customer,omitempty"`
	CustomFields      []CustomField   `json:"custom_fields,omitempty"`
	Contactpersons    []ContactPerson `json:"contactpersons,omitempty"`
	Notes             []struct {
		NoteID        string `json:"note_id,omitempty"`
		Description   string `json:"description,omitempty"`
		CommentedBy   string `json:"commented_by,omitempty"`
		CommentedTime string `json:"commented_time,omitempty"`
	} `json:"notes,omitempty"`
	PaymentGateways        []PaymentGateway `json:"payment_gateways,omitempty"`
	CreateBackdatedInvoice bool             `json:"create_backdated_invoice,omitempty"`
	TemplateID             string           `json:"template_id,omitempty"`
}

type Addon struct {
	AddonCode        string  `json:"addon_code,omitempty"`
	Name             string  `json:"name,omitempty"`
	AddonDescription string  `json:"addon_description,omitempty"`
	Quantity         float64 `json:"quantity,omitempty"`
	Price            float64 `json:"price,omitempty"`
	Discount         float64 `json:"discount,omitempty"`
	Total            float64 `json:"total,omitempty"`
	TaxID            string  `json:"tax_id,omitempty"`
}

type Plan struct {
	PlanCode        string  `json:"plan_code,omitempty"`
	Name            string  `json:"name,omitempty"`
	Quantity        float64 `json:"quantity,omitempty"`
	Price           float64 `json:"price,omitempty"`
	Discount        float64 `json:"discount,omitempty"`
	Total           float64 `json:"total,omitempty"`
	SetupFee        float64 `json:"setup_fee,omitempty"`
	PlanDescription string  `json:"plan_description,omitempty"`
	TaxID           string  `json:"tax_id,omitempty"`
	TrialDays       int64   `json:"trial_days,omitempty"`
}

type Customer struct {
	CustomerID        string  `json:"customer_id,omitempty"`
	DisplayName       string  `json:"display_name,omitempty"`
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
	Street2   string `json:"street2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Country   string `json:"country,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Fax       string `json:"fax,omitempty"`
}

type CustomField struct {
	Index    int64       `json:"index,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Label    string      `json:"label,omitempty"`
	DataType string      `json:"data_type,omitempty"`
}

type ContactPerson struct {
	ContactpersonID string `json:"contactperson_id,omitempty"`
}

type PaymentGateway struct {
	PaymentGateway string `json:"payment_gateway,omitempty"`
}
