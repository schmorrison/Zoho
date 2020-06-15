package expense

import (
	"fmt"
	"go-zoho/zoho"
)

// GetExpenseReports will return a list of all submitted expense reports as specified by
// https://www.zoho.com/expense/api/v1/#Expense_Reports_List_of_all_expense_reports
func (c *API) GetExpenseReports(request interface{}, organizationId string, params map[string]zoho.Parameter) (data ExpenseReportResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         ExpenseReportModule,
		URL:          fmt.Sprintf(ExpenseAPIEndpoint+"%s", ExpenseReportModule),
		Method:       zoho.HTTPGet,
		ResponseData: &ExpenseReportResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
		Headers: map[string]string{
			ExpenseAPIEndpointHeader: organizationId,
		},
	}

	for k, v := range params {
		endpoint.URLParameters[k] = v
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ExpenseReportResponse{}, fmt.Errorf("Failed to retrieve expense reports: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ExpenseReportResponse); ok {
		return *v, nil
	}
	return ExpenseReportResponse{}, fmt.Errorf("Data retrieved was not 'ExpenseReportResponse'")
}

// ExpenseReportResponse is the data returned by GetExpenseReports
type ExpenseReportResponse struct {
	Code           int `json:"code"`
	ExpenseReports []struct {
		ApprovedDate  string `json:"approved_date"`
		ApproverEmail string `json:"approver_email"`
		ApproverID    string `json:"approver_id"`
		ApproverName  string `json:"approver_name"`
		CommentsCount int    `json:"comments_count"`
		CreatedByID   string `json:"created_by_id"`
		CreatedByName string `json:"created_by_name"`
		CreatedTime   string `json:"created_time"`
		CurrencyCode  string `json:"currency_code"`
		CurrencyID    string `json:"currency_id"`
		CustomFields  []struct {
			CustomfieldID string `json:"customfield_id"`
			Label         string `json:"label"`
			Value         string `json:"value"`
		} `json:"custom_fields"`
		CustomerID                string  `json:"customer_id"`
		CustomerName              string  `json:"customer_name"`
		Description               string  `json:"description"`
		DueDate                   string  `json:"due_date"`
		DueDays                   string  `json:"due_days"`
		EndDate                   string  `json:"end_date"`
		IsArchived                bool    `json:"is_archived"`
		LastModifiedTime          string  `json:"last_modified_time"`
		LastSubmittedDate         string  `json:"last_submitted_date"`
		NonReimbursableTotal      float64 `json:"non_reimbursable_total"`
		PolicyID                  string  `json:"policy_id"`
		PolicyName                string  `json:"policy_name"`
		PolicyViolated            bool    `json:"policy_violated"`
		ProjectID                 string  `json:"project_id"`
		ProjectName               string  `json:"project_name"`
		ReimbursableTotal         float64 `json:"reimbursable_total"`
		ReimbursementDate         string  `json:"reimbursement_date"`
		ReportID                  string  `json:"report_id"`
		ReportName                string  `json:"report_name"`
		ReportNumber              string  `json:"report_number"`
		StartDate                 string  `json:"start_date"`
		Status                    string  `json:"status"`
		SubmittedBy               string  `json:"submitted_by"`
		SubmittedDate             string  `json:"submitted_date"`
		SubmittedToEmail          string  `json:"submitted_to_email"`
		SubmittedToID             string  `json:"submitted_to_id"`
		SubmittedToName           string  `json:"submitted_to_name"`
		SubmitterEmail            string  `json:"submitter_email"`
		SubmitterName             string  `json:"submitter_name"`
		Total                     float64 `json:"total"`
		UncategorizedExpenseCount float64 `json:"uncategorized_expense_count"`
	} `json:"expense_reports"`
	Message string `json:"message"`
}
