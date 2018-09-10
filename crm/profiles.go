package crm

import (
	".."
)

var ProfilesEndpoint = zoho.Endpoint{
	Name:         "profiles",
	URL:          "https://www.zohoapis.com/crm/v2/settings/profiles/${id}",
	Methods:      []zoho.HttpMethod{zoho.HTTPGet},
	ResponseData: ProfilesResponse{},
	OptionalSegments: map[string]string{
		"id": "",
	},
}

type ProfilesResponse struct {
	Profiles []struct {
		Name        string `json:"name,omitempty"`
		ModifiedBy  string `json:"modified_by,omitempty"`
		Description string `json:"description,omitempty"`
		ID          string `json:"id,omitempty"`
		Category    bool   `json:"category,omitempty"`
	} `json:"profiles,omitempty"`
}
