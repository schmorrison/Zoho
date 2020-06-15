package expense

import (
	"fmt"
	"go-zoho/zoho"
)

// GetOrganization will return the organization data related to the logged in account
// Parse this response to get organization_id field and pass this value in each expense apis
// https://www.zoho.com/expense/api/v1/#organization-id
//
// Alternatively organization_id can also be known after login to zoho web page at
// https://expense.zoho.com/app#/organizations
func (c *API) GetOrganization() (data OrganizationResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         OrganizationsModule,
		URL:          fmt.Sprintf(ExpenseAPIEndpoint+"%s", OrganizationsModule),
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
	Code          int    `json:"code"`
	Message       string `json:"message"`
	Organizations []struct {
		AccountCreatedDate   string `json:"account_created_date"`
		ContactName          string `json:"contact_name"`
		CurrencyCode         string `json:"currency_code"`
		CurrencyFormat       string `json:"currency_format"`
		CurrencyID           string `json:"currency_id"`
		CurrencySymbol       string `json:"currency_symbol"`
		Email                string `json:"email"`
		FiscalYearStartMonth int    `json:"fiscal_year_start_month"`
		IsDefaultOrg         bool   `json:"is_default_org"`
		IsOrgActive          bool   `json:"is_org_active"`
		LanguageCode         string `json:"language_code"`
		Name                 string `json:"name"`
		OrganizationID       string `json:"organization_id"`
		PricePrecision       int    `json:"price_precision"`
		TimeZone             string `json:"time_zone"`
	} `json:"organizations"`
}
