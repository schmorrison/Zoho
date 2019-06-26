package crm

import (
	"fmt"

	"github.com/schmorrison/Zoho"
)

// GetRoles will return the list of roles in this CRM organization
// https://www.zoho.com/crm/help/api/v2/#Roles-APIs
func (c *API) GetRoles() (data RolesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "roles",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/settings/roles", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &RolesResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return RolesResponse{}, fmt.Errorf("Failed to retrieve roles: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*RolesResponse); ok {
		return *v, nil
	}

	return RolesResponse{}, fmt.Errorf("Data retrieved was not 'RolesResponse'")
}

// GetRole will return the role specified by the id
// https://www.zoho.com/crm/help/api/v2/#get-single-role-data
func (c *API) GetRole(id string) (data RolesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "roles",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/settings/roles/%s", c.ZohoTLD, id),
		Method:       zoho.HTTPGet,
		ResponseData: &RolesResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return RolesResponse{}, fmt.Errorf("Failed to retrieve role (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*RolesResponse); ok {
		return *v, nil
	}

	return RolesResponse{}, fmt.Errorf("Data retrieved was not 'RolesResponse'")
}

// RolesResponse is the data returned by GetRoles and GetRole
type RolesResponse struct {
	Roles []struct {
		DisplayLabel string `json:"display_label,omitempty"`
		Name         string `json:"name,omitempty"`
		ID           string `json:"id,omitempty"`
		ReportingTo  string `json:"reporting_to,omitempty"`
		AdminUser    bool   `json:"admin_user,omitempty"`
	} `json:"roles,omitempty"`
}
