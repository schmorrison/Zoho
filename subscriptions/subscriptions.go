package subscriptions

type Subscription struct {
	SubscriptionId    string  `json:"subscription_id"`
	Name              string  `json:"name"`
	Status            string  `json:"status"`
	Amount            float64 `json:"amount"`
	CreatedAt         string  `json:"created_at"`
	ActivatedAt       string  `json:"activated_at"`
	CurrentTermEndsAt string  `json:"current_term_ends_at"`
	LastBillingAt     string  `json:"last_billing_at"`
	NextBillingAt     string  `json:"next_billing_at"`
	ExpiresAt         string  `json:"expires_at"`
	Interval          int64   `json:"interval"`
	IntervalUnit      string  `json:"interval_unit"`
	AutoCollect       bool    `json:"auto_collect"`
	CreatedTime       string  `json:"created_time"`
	UpdatedTime       string  `json:"updated_time"`
	ReferenceId       string  `json:"reference_id"`
	SalespersonId     string  `json:"salesperson_id"`
	SalespersonName   string  `json:"salesperson_name"`
	ChildInvoiceId    string  `json:"child_invoice_id"`
	CurrencyCode      string  `json:"currency_code"`
	CurrencySymbol    string  `json:"currency_symbol"`
	EndOfTerm         bool    `json:"end_of_term"`
	ProductId         string  `json:"product_id"`
	ProductName       string  `json:"product_name"`
	Plan              struct {
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
	Card              struct{}                 `json:"card"`
	PaymentTerms      int64                    `json:"payment_terms"`
	PaymentTermsLabel string                   `json:"payment_terms_label"`
	CanAddBankAccount bool                     `json:"can_add_bank_account"`
	Customer          Customer                 `json:"customer"`
	CustomFields      []map[string]interface{} `json:"custom_fields"`
	Contactpersons    []struct {
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
	Status            string  `json:"salutation"`
	Status            string  `json:"first_name"`
	Status            string  `json:"last_name"`
	Status            string  `json:"email"`
	Status            string  `json:"company_name"`
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
	Zip       int64  `json:"zip"`
	Fax       int64  `json:"fax"`
}
