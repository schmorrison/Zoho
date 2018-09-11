package crm

import (
	"fmt"

	".."
)

func (c *API) GetRoles() (data RolesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "roles",
		URL:          "https://www.zohoapis.com/crm/v2/settings/roles",
		Method:       zoho.HTTPGet,
		ResponseData: RolesResponse{},
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return RolesResponse{}, fmt.Errorf("Failed to retrieve roles: %s", err)
	}

	if v, ok := endpoint.ResponseData.(RolesResponse); ok {
		return v, nil
	}

	return RolesResponse{}, fmt.Errorf("Data retrieved was not 'RolesResponse'")
}

func (c *API) GetRole(id string) (data RolesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "roles",
		URL:          fmt.Sprintf("https://www.zohoapis.com/crm/v2/settings/roles/%s", id),
		Method:       zoho.HTTPGet,
		ResponseData: RolesResponse{},
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return RolesResponse{}, fmt.Errorf("Failed to retrieve role (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(RolesResponse); ok {
		return v, nil
	}

	return RolesResponse{}, fmt.Errorf("Data retrieved was not 'RolesResponse'")
}

type RolesResponse struct {
	Roles []struct {
		DisplayLabel string `json:"display_label,omitempty"`
		Name         string `json:"name,omitempty"`
		ID           string `json:"id,omitempty"`
		ReportingTo  string `json:"reporting_to,omitempty"`
		AdminUser    bool   `json:"admin_user,omitempty"`
	} `json:"roles,omitempty"`
}
