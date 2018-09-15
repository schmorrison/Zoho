package crm

import (
	"fmt"

	"github.com/schmorrison/Zoho"
)

func (c *API) GetUsers(kind UserType) (data UsersResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "users",
		URL:          "https://www.zohoapis.com/crm/v2/users",
		Method:       zoho.HTTPGet,
		ResponseData: &UsersResponse{},
		URLParameters: map[string]zoho.Parameter{
			"type": kind,
		},
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return UsersResponse{}, fmt.Errorf("Failed to retrieve users: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UsersResponse); ok {
		return *v, nil
	}

	return UsersResponse{}, fmt.Errorf("Data retrieved was not 'UsersResponse'")
}

func (c *API) GetUser(id string) (data UsersResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "users",
		URL:          fmt.Sprintf("https://www.zohoapis.com/crm/v2/users/%s", id),
		Method:       zoho.HTTPGet,
		ResponseData: UsersResponse{},
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return UsersResponse{}, fmt.Errorf("Failed to retrieve user (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(UsersResponse); ok {
		return v, nil
	}

	return UsersResponse{}, fmt.Errorf("Data retrieved was not 'UsersResponse'")
}

type UserType = zoho.Parameter

const (
	None                  UserType = ""
	AllUsers              UserType = "AllUsers"
	ActiveUsers           UserType = "ActiveUsers"
	DeactiveUsers         UserType = "DeactiveUsers"
	ConfirmedUsers        UserType = "ConfirmedUsers"
	NotConfirmedUsers     UserType = "NotConfirmedUsers"
	DeletedUsers          UserType = "DeletedUsers"
	ActiveConfirmedUsers  UserType = "ActiveConfirmedUsers"
	AdminUsers            UserType = "AdminUsers"
	ActiveConfirmedAdmins UserType = "ActiveConfirmedAdmins"
	CurrentUser           UserType = "CurrentUser"
)

type UsersResponse struct {
	Users []struct {
		Country string `json:"country,omitempty"`
		Role    struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"role,omitempty"`
		City       string `json:"city,omitempty"`
		Language   string `json:"language,omitempty"`
		Locale     string `json:"locale,omitempty"`
		ModifiedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		Street        string `json:"street,omitempty"`
		Currency      string `json:"Currency,omitempty"`
		Alias         string `json:"alias,omitempty"`
		ID            string `json:"id,omitempty"`
		State         string `json:"state,omitempty"`
		Fax           string `json:"fax,omitempty"`
		CountryLocale string `json:"country_locale,omitempty"`
		FirstName     string `json:"first_name,omitempty"`
		Email         string `json:"email,omitempty"`
		ReportingTo   string `json:"Reporting_To,omitempty"`
		Zip           string `json:"zip,omitempty"`
		CreatedTime   string `json:"created_time,omitempty"`
		ModifiedTime  string `json:"modified_time,omitempty"`
		Website       string `json:"website,omitempty"`
		TimeFormat    string `json:"time_format,omitempty"`
		Offset        int64  `json:"offset,omitempty"`
		Profile       struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"profile,omitempty"`
		Mobile    string `json:"mobile,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		CreatedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"created_by,omitempty"`
		Zuid        string `json:"zuid,omitempty"`
		Confirm     bool   `json:"confirm,omitempty"`
		FullName    string `json:"full_name,omitempty"`
		Territories struct {
			Manager bool   `json:"manager,omitempty"`
			Name    string `json:"name,omitempty"`
			ID      string `json:"id,omitempty"`
		} `json:"territories,omitempty"`
		Phone         string `json:"phone,omitempty"`
		Dob           string `json:"dob,omitempty"`
		DateFormat    string `json:"date_format,omitempty"`
		Status        string `json:"status,omitempty"`
		CustomizeInfo struct {
			NotesDesc       string `json:"notes_desc,omitempty"`
			ShowRightPanel  bool   `json:"show_right_panel,omitempty"`
			BcView          string `json:"bc_view,omitempty"`
			ShowHome        bool   `json:"show_home,omitempty"`
			UnpinRecentItem bool   `json:"unpin_recent_item,omitempty"`
		} `json:"customize_info,omitempty,omitempty"`
		Signature           string `json:"signature,omitempty,omitempty"`
		NameFormat          string `json:"name_format,omitempty,omitempty"`
		PersonalAccount     bool   `json:"personal_account,omitempty,omitempty"`
		NtcNotificationType []int  `json:"ntc_notification_type,omitempty,omitempty"`
		DefaultTabGroup     string `json:"default_tab_group,omitempty,omitempty"`
		Theme               struct {
			NormalTab struct {
				FontColor  string `json:"font_color,omitempty"`
				Background string `json:"background,omitempty"`
			} `json:"normal_tab,omitempty"`
			SelectedTab struct {
				FontColor  string `json:"font_color,omitempty"`
				Background string `json:"background,omitempty"`
			} `json:"selected_tab,omitempty"`
			NewBackground string `json:"new_background,omitempty"`
			Background    string `json:"background,omitempty"`
			Screen        string `json:"screen,omitempty"`
			Type          string `json:"type,omitempty"`
		} `json:"theme,omitempty,omitempty"`
		TelephonyEnabled bool   `json:"telephony_enabled,omitempty,omitempty"`
		ImapStatus       bool   `json:"imap_status,omitempty,omitempty"`
		DecimalSeparator string `json:"decimal_separator,omitempty,omitempty"`
		TimeZone         string `json:"time_zone,omitempty"`
		RtlEnabled       bool   `json:"rtl_enabled,omitempty,omitempty"`
		NtcEnabled       bool   `json:"ntc_enabled,omitempty,omitempty"`
	} `json:"users,omitempty"`
	Info struct {
		PerPage     int  `json:"per_page,omitempty"`
		Count       int  `json:"count,omitempty"`
		Page        int  `json:"page,omitempty"`
		MoreRecords bool `json:"more_records,omitempty"`
	} `json:"info,omitempty"`
}
