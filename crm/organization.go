package crm

import (
	"fmt"

	"github.com/schmorrison/Zoho"
)

// GetOrganization will return the organization data related to the logged in account
// https://www.zoho.com/crm/help/api/v2/#Organization-API
func (c *API) GetOrganization() (data OrganizationResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "organization",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/org", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &OrganizationResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return OrganizationResponse{}, fmt.Errorf("Failed to retrieve organization: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*OrganizationResponse); ok {
		return *v, nil
	}

	return OrganizationResponse{}, fmt.Errorf("Data retrieved was not 'OrganizationResponse'")
}

// OrganizationResponse is the data returned by GetOrganization
type OrganizationResponse struct {
	Org []struct {
		Country        string `json:"country,omitempty"`
		McStatus       bool   `json:"mc_status,omitempty"`
		GappsEnabled   bool   `json:"gapps_enabled,omitempty"`
		ID             string `json:"id,omitempty"`
		State          string `json:"state,omitempty"`
		EmployeeCount  string `json:"employee_count,omitempty"`
		Website        string `json:"website,omitempty"`
		CurrencySymbol string `json:"currency_symbol,omitempty"`
		Mobile         string `json:"mobile,omitempty"`
		CurrencyLocale string `json:"currency_locale,omitempty"`
		PrimaryZuid    string `json:"primary_zuid,omitempty"`
		Zgid           string `json:"zgid,omitempty"`
		CountryCode    string `json:"country_code,omitempty"`
		LicenseDetails struct {
			PaidExpiry            Time   `json:"paid_expiry,omitempty"`
			UsersLicensePurchased int    `json:"users_license_purchased,omitempty"`
			TrialType             string `json:"trial_type,omitempty"`
			TrialExpiry           Time   `json:"trial_expiry,omitempty"`
			Paid                  bool   `json:"paid,omitempty"`
			PaidType              string `json:"paid_type,omitempty"`
		} `json:"license_details,omitempty"`
		CompanyName     string `json:"company_name,omitempty"`
		PrivacySettings bool   `json:"privacy_settings,omitempty"`
		PrimaryEmail    string `json:"primary_email,omitempty"`
		IsoCode         string `json:"iso_code,omitempty"`
	} `json:"org,omitempty"`
}
