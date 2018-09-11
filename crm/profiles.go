package crm

import (
	"fmt"

	".."
)

func (c *API) GetProfiles() (data ProfilesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "profiles",
		URL:          "https://www.zohoapis.com/crm/v2/settings/profiles",
		Method:       zoho.HTTPGet,
		ResponseData: ProfilesResponse{},
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return ProfilesResponse{}, fmt.Errorf("Failed to retrieve profiles: %s", err)
	}

	if v, ok := endpoint.ResponseData.(ProfilesResponse); ok {
		return v, nil
	}

	return ProfilesResponse{}, fmt.Errorf("Data retrieved was not 'ProfilesResponse'")
}

func (c *API) GetProfile(id string) (data ProfilesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "profiles",
		URL:          fmt.Sprintf("https://www.zohoapis.com/crm/v2/settings/profiles/%s", id),
		Method:       zoho.HTTPGet,
		ResponseData: ProfilesResponse{},
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return ProfilesResponse{}, fmt.Errorf("Failed to retrieve profile (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(ProfilesResponse); ok {
		return v, nil
	}

	return ProfilesResponse{}, fmt.Errorf("Data retrieved was not 'ProfilesResponse'")
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
