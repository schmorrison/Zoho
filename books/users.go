package books

import (
	"encoding/json"
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetCurrentUser will return the currently authenticated users
// https://www.zoho.com/books/api/v3/users/#get-current-user
func (c *API) GetCurrentUser() (data CurrentUserResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "users",
		URL:          fmt.Sprintf("https://books.zoho.%s/api/v3/users/me", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &CurrentUserResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CurrentUserResponse{}, fmt.Errorf("Failed to retrieve current user: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CurrentUserResponse); ok {
		return *v, nil
	}

	return CurrentUserResponse{}, fmt.Errorf("Data retrieved was not 'UsersResponse'")
}

func (m *MorePermissions) UnmarshalJSON(data []byte) error {
	if string(data) == `""` {
		return nil
	}

	type tmp MorePermissions
	return json.Unmarshal(data, (*tmp)(m))
}

type MorePermissions []Permission

type Permission struct {
	IsEnabled           bool   `json:"is_enabled,omitempty"`
	PermissionFormatted string `json:"permission_formatted,omitempty"`
	Permission          string `json:"permission,omitempty"`
}

// CurrentUserResponse is the data returned by GetCurrentUser
type CurrentUserResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	User    struct {
		UserID   string `json:"user_id,omitempty"`
		Name     string `json:"name,omitempty"`
		EmailIds []struct {
			IsSelected bool   `json:"is_selected,omitempty"`
			Email      string `json:"email,omitempty"`
		} `json:"email_ids,omitempty"`
		Status   string `json:"status,omitempty"`
		UserRole string `json:"user_role,omitempty"`
		UserType string `json:"user_type,omitempty"`
		RoleID   string `json:"role_id,omitempty"`
		PhotoURL string `json:"photo_url,omitempty"`
		Role     struct {
			Role struct {
				RoleName    string `json:"role_name,omitempty"`
				Permissions []struct {
					FullAccess        bool            `json:"full_access,omitempty"`
					Entity            string          `json:"entity,omitempty"`
					MorePermissions   MorePermissions `json:"more_permissions,omitempty"`
					ReportPermissions []struct {
						Reports []struct {
							FullAccess          bool   `json:"full_access,omitempty"`
							CanSchedule         bool   `json:"can_schedule,omitempty"`
							ReportConstant      string `json:"report_constant,omitempty"`
							CanShare            bool   `json:"can_share,omitempty"`
							IsExportEnabled     bool   `json:"is_export_enabled,omitempty"`
							ReportNameFormatted string `json:"report_name_formatted,omitempty"`
							IsScheduleEnabled   bool   `json:"is_schedule_enabled,omitempty"`
							CanExport           bool   `json:"can_export,omitempty"`
							CanAccess           bool   `json:"can_access,omitempty"`
						} `json:"reports,omitempty"`
						ReportGroupFormatted string `json:"report_group_formatted,omitempty"`
						ReportGroup          string `json:"report_group,omitempty"`
					} `json:"report_permissions,omitempty"`
				} `json:"permissions,omitempty"`
				Description string `json:"description,omitempty"`
				DisplayName string `json:"display_name,omitempty"`
			} `json:"role,omitempty"`
			RoleID string `json:"role_id,omitempty"`
			Name   string `json:"name,omitempty"`
			Email  string `json:"email,omitempty"`
			Zuid   string `json:"zuid,omitempty"`
		} `json:"role,omitempty"`
		IsClaimant          bool          `json:"is_claimant,omitempty"`
		IsEmployee          bool          `json:"is_employee,omitempty"`
		Email               string        `json:"email,omitempty"`
		IsCustomerSegmented bool          `json:"is_customer_segmented,omitempty"`
		IsVendorSegmented   bool          `json:"is_vendor_segmented,omitempty"`
		IsAccountant        bool          `json:"is_accountant,omitempty"`
		CreatedTime         string        `json:"created_time,omitempty"`
		CustomFields        []interface{} `json:"custom_fields,omitempty"`
		CustomFieldHash     struct {
		} `json:"custom_field_hash,omitempty"`
		IsAssociatedForApproval  bool          `json:"is_associated_for_approval,omitempty"`
		IsAssociatedWithOrgEmail bool          `json:"is_associated_with_org_email,omitempty"`
		Branches                 []interface{} `json:"branches,omitempty"`
		DefaultBranchID          string        `json:"default_branch_id,omitempty"`
	} `json:"user,omitempty"`
}
