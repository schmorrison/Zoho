package recruit

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetUsers returns a list of users. Users are those who are allowed to access and manage records.
// https://www.zoho.com/recruit/developer-guide/apiv2/get-users.html
// https://recruit.zoho.eu/recruit/v2/users?type={AllUsers,ActiveUsers,DeactiveUsers,ConfirmedUsers,NotConfirmedUsers,DeletedUsers,ActiveConfirmedUsers,AdminUsers,ActiveConfirmedAdmins,CurrentUser}
func (c *API) GetUsers(params map[string]zoho.Parameter) (data UsersResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetUsers",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/users", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &UsersResponse{},
		URLParameters: map[string]zoho.Parameter{
			"type": "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UsersResponse{}, fmt.Errorf("failed to retrieve users: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UsersResponse); ok {
		return *v, nil
	}

	return UsersResponse{}, fmt.Errorf("data returned was not 'UsersResponse'")
}

// UsersResponse is the data returned by GetUsers
type UsersResponse struct {
	Users []struct {
		DecimalSeparator string `json:"decimal_separator,omitempty"`
		Role             struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"role,omitempty"`
		CustomizeInfo struct {
			NotesDesc       interface{} `json:"notes_desc,omitempty"`
			ShowRightPanel  interface{} `json:"show_right_panel,omitempty"`
			BcView          interface{} `json:"bc_view,omitempty"`
			ShowHome        bool        `json:"show_home,omitempty"`
			ShowDetailView  interface{} `json:"show_detail_view,omitempty"`
			UnpinRecentItem interface{} `json:"unpin_recent_item,omitempty"`
		} `json:"customize_info,omitempty"`
		Signature interface{} `json:"signature,omitempty"`
		Profile   struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"profile,omitempty"`
		LastName        string        `json:"last_name,omitempty"`
		NameFormat      string        `json:"name_format,omitempty"`
		Language        string        `json:"language,omitempty"`
		Locale          string        `json:"locale,omitempty"`
		TimeZone        string        `json:"time_zone,omitempty"`
		PersonalAccount bool          `json:"personal_account,omitempty"`
		Zuid            string        `json:"zuid,omitempty"`
		Confirm         bool          `json:"confirm,omitempty"`
		DefaultTabGroup string        `json:"default_tab_group,omitempty"`
		FullName        string        `json:"full_name,omitempty"`
		Territories     []interface{} `json:"territories,omitempty"`
		Dob             string        `json:"dob,omitempty"`
		DateFormat      string        `json:"date_format,omitempty"`
		Theme           struct {
			NormalTab struct {
				FontColor  string `json:"font_color,omitempty"`
				Background string `json:"background,omitempty"`
			} `json:"normal_tab,omitempty"`
			SelectedTab struct {
				FontColor  string `json:"font_color,omitempty"`
				Background string `json:"background,omitempty"`
			} `json:"selected_tab,omitempty"`
			NewBackground interface{} `json:"new_background,omitempty"`
			Background    string      `json:"background,omitempty"`
			Screen        string      `json:"screen,omitempty"`
			Type          string      `json:"type,omitempty"`
		} `json:"theme,omitempty"`
		ID            string      `json:"id,omitempty"`
		CountryLocale string      `json:"country_locale,omitempty"`
		FirstName     interface{} `json:"first_name,omitempty"`
		Email         string      `json:"email,omitempty"`
		Status        string      `json:"status,omitempty"`
	} `json:"users,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}
