package crm

import (
	"fmt"

	"github.com/schmorrison/Zoho"
)

// GetUsers will return the list of users in the CRM organization. The list can be filtered using the
// 'kind' parameter
// https://www.zoho.com/crm/help/api/v2/#Users-APIs
func (c *API) GetUsers(kind UserType) (data UsersResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "users",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/users", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &UsersResponse{},
		URLParameters: map[string]zoho.Parameter{
			"type": kind,
		},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UsersResponse{}, fmt.Errorf("Failed to retrieve users: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UsersResponse); ok {
		return *v, nil
	}

	return UsersResponse{}, fmt.Errorf("Data retrieved was not 'UsersResponse'")
}

// GetUser will return the user specified by id
// https://www.zoho.com/crm/help/api/v2/#get-single-user-data
func (c *API) GetUser(id string) (data UsersResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "users",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/users/%s", c.ZohoTLD, id),
		Method:       zoho.HTTPGet,
		ResponseData: &UsersResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UsersResponse{}, fmt.Errorf("Failed to retrieve user (%s): %s", id, err)
	}

	if v, ok := endpoint.ResponseData.(*UsersResponse); ok {
		return *v, nil
	}

	return UsersResponse{}, fmt.Errorf("Data retrieved was not 'UsersResponse'")
}

// UserType is the 'kind' parameter in the GetUsers function
type UserType = zoho.Parameter

const (
	// None - Do not filter the Users list
	None UserType = ""
	// AllUsers - To list all users in your organization (both active and inactive users)
	AllUsers UserType = "AllUsers"
	// ActiveUsers - To get the list of all Active Users
	ActiveUsers UserType = "ActiveUsers"
	// DeactiveUsers - To get the list of all users who were deactivated
	DeactiveUsers UserType = "DeactiveUsers"
	// ConfirmedUsers - To get the list of confirmed users
	ConfirmedUsers UserType = "ConfirmedUsers"
	// NotConfirmedUsers - To get the list of non-confirmed users
	NotConfirmedUsers UserType = "NotConfirmedUsers"
	// DeletedUsers - To get the list of deleted users
	DeletedUsers UserType = "DeletedUsers"
	// ActiveConfirmedUsers - To get the list of active users who are also confirmed
	ActiveConfirmedUsers UserType = "ActiveConfirmedUsers"
	// AdminUsers - To get the list of admin users.
	AdminUsers UserType = "AdminUsers"
	// ActiveConfirmedAdmins - To get the list of active users with the administrative privileges and are also confirmed
	ActiveConfirmedAdmins UserType = "ActiveConfirmedAdmins"
	// CurrentUser - To get the list of current CRM users
	CurrentUser UserType = "CurrentUser"
)

// UsersResponse is the data returned by GetUsers and GetUser
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
		Territories []struct {
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
	Info PageInfo `json:"info,omitempty"`
}
