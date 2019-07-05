package crm

import (
	"fmt"

	"github.com/schmorrison/Zoho"
)

// GetProfiles will return the list of profiles in this CRM organization
// https://www.zoho.com/crm/help/api/v2/#Profiles-APIs
func (c *API) GetProfiles() (data ProfilesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "profiles",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/settings/profiles", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &ProfilesResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ProfilesResponse{}, fmt.Errorf("Failed to retrieve profiles: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ProfilesResponse); ok {
		return *v, nil
	}

	return ProfilesResponse{}, fmt.Errorf("Data retrieved was not 'ProfilesResponse'")
}

// GetProfile will return the profile specified by id
// https://www.zoho.com/crm/help/api/v2/#get-single-profile-data
func (c *API) GetProfile(id string) (data ProfilesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "profiles",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/settings/profiles/%s", c.ZohoTLD, id),
		Method:       zoho.HTTPGet,
		ResponseData: &ProfilesResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ProfilesResponse{}, fmt.Errorf("Failed to retrieve profile (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*ProfilesResponse); ok {
		return *v, nil
	}

	return ProfilesResponse{}, fmt.Errorf("Data retrieved was not 'ProfilesResponse'")
}

// ProfilesResponse is the data returned by GetProfiles and GetProfile
type ProfilesResponse struct {
	Profiles []struct {
		Name        string `json:"name,omitempty"`
		ModifiedBy  string `json:"modified_by,omitempty"`
		Description string `json:"description,omitempty"`
		ID          string `json:"id,omitempty"`
		Category    bool   `json:"category,omitempty"`
	} `json:"profiles,omitempty"`
}
