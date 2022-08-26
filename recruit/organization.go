package recruit

import (
	"fmt"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

// GetOrganizationDetails returns organization's data
// https://www.zoho.com/recruit/developer-guide/apiv2/get-org-data.html
// https://recruit.zoho.eu/recruit/v2/org
func (c *API) GetOrganizationDetails() (data OrganizationResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetOrganizationDetails",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/org", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &OrganizationResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return OrganizationResponse{}, fmt.Errorf("failed to retrieve organization's data: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*OrganizationResponse); ok {
		return *v, nil
	}

	return OrganizationResponse{}, fmt.Errorf("data returned was not 'OrganizationResponse'")
}

// OrganizationResponse is the data returned by GetOrganizationDetails
type OrganizationResponse struct {
	Org []struct {
		EmployeeCount  string      `json:"employee_count"`
		Zip            interface{} `json:"zip"`
		Country        interface{} `json:"country"`
		Website        string      `json:"website"`
		City           interface{} `json:"city"`
		Mobile         interface{} `json:"mobile"`
		Description    interface{} `json:"description"`
		Type           string      `json:"type"`
		TimeZone       string      `json:"time_zone"`
		McStatus       bool        `json:"mc_status"`
		CountryCode    string      `json:"country_code"`
		LicenseDetails struct {
			TrialType   string    `json:"trial_type"`
			TrialExpiry time.Time `json:"trial_expiry"`
			Paid        bool      `json:"paid"`
			PaidType    string    `json:"paid_type"`
			NoOfUsers   int       `json:"no_of_users"`
		} `json:"license_details"`
		Zgid          string      `json:"zgid"`
		Phone         string      `json:"phone"`
		Street        interface{} `json:"street"`
		CompanyName   string      `json:"company_name"`
		PrimaryUserid string      `json:"primary_userid"`
		Alias         interface{} `json:"alias"`
		PrimaryEmail  string      `json:"primary_email"`
		EditorMode    string      `json:"editor_mode"`
		State         interface{} `json:"state"`
		Fax           interface{} `json:"fax"`
	} `json:"org"`
}
