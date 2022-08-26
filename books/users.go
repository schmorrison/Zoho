package books

import (
	"fmt"

	zoho "github.com/recap-technologies/Zoho"
)

// GetCurrentUser will return the currently authenticated users
// https://www.zoho.com/books/api/v3/users/#get-current-user
func (c *API) GetCurrentUser(id string) (data CurrentUserResponse, err error) {
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

// CurrentUserResponse is the data returned by GetCurrentUser
type CurrentUserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	User    struct {
		UserID   string `json:"user_id"`
		Name     string `json:"name"`
		EmailIds []struct {
			IsSelected bool   `json:"is_selected"`
			Email      string `json:"email"`
		} `json:"email_ids"`
		Status   string `json:"status"`
		UserRole string `json:"user_role"`
		UserType string `json:"user_type"`
		RoleID   string `json:"role_id"`
		PhotoURL string `json:"photo_url"`
		Role     struct {
			Role struct {
				RoleName    string `json:"role_name"`
				Permissions []struct {
					FullAccess      bool   `json:"full_access"`
					Entity          string `json:"entity"`
					MorePermissions []struct {
						IsEnabled           bool   `json:"is_enabled"`
						PermissionFormatted string `json:"permission_formatted"`
						Permission          string `json:"permission"`
					} `json:"more_permissions,omitempty"`
					ReportPermissions []struct {
						Reports []struct {
							FullAccess          bool   `json:"full_access"`
							CanSchedule         bool   `json:"can_schedule"`
							ReportConstant      string `json:"report_constant"`
							CanShare            bool   `json:"can_share"`
							IsExportEnabled     bool   `json:"is_export_enabled"`
							ReportNameFormatted string `json:"report_name_formatted"`
							IsScheduleEnabled   bool   `json:"is_schedule_enabled"`
							CanExport           bool   `json:"can_export"`
							CanAccess           bool   `json:"can_access"`
						} `json:"reports"`
						ReportGroupFormatted string `json:"report_group_formatted"`
						ReportGroup          string `json:"report_group"`
					} `json:"report_permissions,omitempty"`
				} `json:"permissions"`
				Description string `json:"description"`
				DisplayName string `json:"display_name"`
			} `json:"role"`
			RoleID string `json:"role_id"`
			Name   string `json:"name"`
			Email  string `json:"email"`
			Zuid   string `json:"zuid"`
		} `json:"role"`
		IsClaimant          bool          `json:"is_claimant"`
		IsEmployee          bool          `json:"is_employee"`
		Email               string        `json:"email"`
		IsCustomerSegmented bool          `json:"is_customer_segmented"`
		IsVendorSegmented   bool          `json:"is_vendor_segmented"`
		IsAccountant        bool          `json:"is_accountant"`
		CreatedTime         string        `json:"created_time"`
		CustomFields        []interface{} `json:"custom_fields"`
		CustomFieldHash     struct {
		} `json:"custom_field_hash"`
		IsAssociatedForApproval  bool          `json:"is_associated_for_approval"`
		IsAssociatedWithOrgEmail bool          `json:"is_associated_with_org_email"`
		Branches                 []interface{} `json:"branches"`
		DefaultBranchID          string        `json:"default_branch_id"`
	} `json:"user"`
}
