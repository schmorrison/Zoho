package recruit

import (
	"encoding/json"

	zoho "github.com/schmorrison/Zoho"
)

// GetUsers returns a list of users. Users are those who are allowed to access and manage records.
// https://www.zoho.com/recruit/developer-guide/apiv2/get-users.html
// https://recruit.zoho.eu/recruit/v2/users?type={AllUsers,ActiveUsers,DeactiveUsers,ConfirmedUsers,NotConfirmedUsers,DeletedUsers,ActiveConfirmedUsers,AdminUsers,ActiveConfirmedAdmins,CurrentUser}
func (c *API) GetUsers(params map[string]zohoutils.Parameter) (data UsersResponse, err error) {
	endpoint := zohoutils.Endpoint{
		Name:         "user",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/users", c.ZohoTLD),
		Method:       zohoutils.HTTPGet,
		ResponseData: &UsersResponse{},
		URLParameters: map[string]zohoutils.Parameter{
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
		DecimalSeparator string `json:"decimal_separator"`
		Role             struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"role"`
		CustomizeInfo struct {
			NotesDesc       interface{} `json:"notes_desc"`
			ShowRightPanel  interface{} `json:"show_right_panel"`
			BcView          interface{} `json:"bc_view"`
			ShowHome        bool        `json:"show_home"`
			ShowDetailView  interface{} `json:"show_detail_view"`
			UnpinRecentItem interface{} `json:"unpin_recent_item"`
		} `json:"customize_info"`
		Signature interface{} `json:"signature"`
		Profile   struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"profile"`
		LastName        string        `json:"last_name"`
		NameFormat      string        `json:"name_format"`
		Language        string        `json:"language"`
		Locale          string        `json:"locale"`
		TimeZone        string        `json:"time_zone"`
		PersonalAccount bool          `json:"personal_account"`
		Zuid            string        `json:"zuid"`
		Confirm         bool          `json:"confirm"`
		DefaultTabGroup string        `json:"default_tab_group"`
		FullName        string        `json:"full_name"`
		Territories     []interface{} `json:"territories"`
		Dob             string        `json:"dob"`
		DateFormat      string        `json:"date_format"`
		Theme           struct {
			NormalTab struct {
				FontColor  string `json:"font_color"`
				Background string `json:"background"`
			} `json:"normal_tab"`
			SelectedTab struct {
				FontColor  string `json:"font_color"`
				Background string `json:"background"`
			} `json:"selected_tab"`
			NewBackground interface{} `json:"new_background"`
			Background    string      `json:"background"`
			Screen        string      `json:"screen"`
			Type          string      `json:"type"`
		} `json:"theme"`
		ID            string      `json:"id"`
		CountryLocale string      `json:"country_locale"`
		FirstName     interface{} `json:"first_name"`
		Email         string      `json:"email"`
		Status        string      `json:"status"`
	} `json:"users"`
	Info struct {
		PerPage     int  `json:"per_page"`
		Count       int  `json:"count"`
		Page        int  `json:"page"`
		MoreRecords bool `json:"more_records"`
	} `json:"info"`
}
