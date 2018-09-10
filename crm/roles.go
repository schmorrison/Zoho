package crm

import (
	".."
)

var RolesEndpoint = zoho.Endpoint{
	Name:         "roles",
	URL:          "https://www.zohoapis.com/crm/v2/settings/roles/${id}",
	Methods:      []zoho.HttpMethod{zoho.HTTPGet},
	ResponseData: RolesResponse{},
	OptionalSegments: map[string]string{
		"id": "",
	},
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
