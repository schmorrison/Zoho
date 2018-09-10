package crm

import (
	".."
)

var UsersEndpoint = zoho.Endpoint{
	Name:         "users",
	URL:          "https://www.zohoapis.com/crm/v2/users/${id}",
	Methods:      []zoho.HttpMethod{zoho.HTTPGet},
	ResponseData: UsersResponse{},
	URLParameters: map[string]zoho.Parameter{
		"type": None,
	},
	OptionalSegments: map[string]string{
		"id": "",
	},
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
		CustomizeInfo struct {
			NotesDesc       string `json:"notes_desc,omitempty"`
			ShowRightPanel  bool   `json:"show_right_panel,omitempty"`
			BcView          string `json:"bc_view,omitempty"`
			ShowHome        bool   `json:"show_home,omitempty"`
			UnpinRecentItem bool   `json:"unpin_recent_item,omitempty"`
		} `json:"customize_info,omitempty,omitempty"`
		City                string `json:"city,omitempty"`
		Signature           string `json:"signature,omitempty,omitempty"`
		NameFormat          string `json:"name_format,omitempty,omitempty"`
		Language            string `json:"language,omitempty"`
		Locale              string `json:"locale,omitempty"`
		PersonalAccount     bool   `json:"personal_account,omitempty,omitempty"`
		NtcNotificationType []int  `json:"ntc_notification_type,omitempty,omitempty"`
		DefaultTabGroup     string `json:"default_tab_group,omitempty,omitempty"`
		Street              string `json:"street,omitempty"`
		Alias               string `json:"alias,omitempty"`
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
		ID               string `json:"id,omitempty"`
		State            string `json:"state,omitempty"`
		CountryLocale    string `json:"country_locale,omitempty"`
		Fax              string `json:"fax,omitempty"`
		FirstName        string `json:"first_name,omitempty"`
		Email            string `json:"email,omitempty"`
		TelephonyEnabled bool   `json:"telephony_enabled,omitempty,omitempty"`
		ImapStatus       bool   `json:"imap_status,omitempty,omitempty"`
		Zip              string `json:"zip,omitempty"`
		DecimalSeparator string `json:"decimal_separator,omitempty,omitempty"`
		Website          string `json:"website,omitempty"`
		TimeFormat       string `json:"time_format,omitempty"`
		Profile          struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"profile,omitempty"`
		Mobile     string `json:"mobile,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		TimeZone   string `json:"time_zone,omitempty"`
		Zuid       string `json:"zuid,omitempty"`
		Confirm    bool   `json:"confirm,omitempty"`
		RtlEnabled bool   `json:"rtl_enabled,omitempty,omitempty"`
		FullName   string `json:"full_name,omitempty"`
		Phone      string `json:"phone,omitempty"`
		Dob        string `json:"dob,omitempty"`
		DateFormat string `json:"date_format,omitempty"`
		NtcEnabled bool   `json:"ntc_enabled,omitempty,omitempty"`
		Status     string `json:"status,omitempty"`
	} `json:"users,omitempty"`
	Info struct {
		PerPage     int  `json:"per_page,omitempty"`
		Count       int  `json:"count,omitempty"`
		Page        int  `json:"page,omitempty"`
		MoreRecords bool `json:"more_records,omitempty"`
	} `json:"info,omitempty"`
}
