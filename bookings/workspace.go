package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) FetchWorkspaces(workspacesID zoho.Parameter) (data WorkspaceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         FetchWorkspacesModule,
		URL:          fmt.Sprintf("https://www.zohoapis.%s/bookings/v1/json/%s", c.ZohoTLD,FetchWorkspacesModule),
		Method:       zoho.HTTPGet,
		ResponseData: &WorkspaceResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}

	if workspacesID != "" {
		endpoint.URLParameters["workspace_id"] = workspacesID
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return WorkspaceResponse{}, fmt.Errorf("Failed to retrieve workspaces: %s", err)
	}

	if v,ok := endpoint.ResponseData.(*WorkspaceResponse); ok {
		return *v, nil
	}
	return WorkspaceResponse{}, fmt.Errorf("Data retrieved was not 'Workspace Response'")
}

type WorkspaceResponse struct {
	Response struct {
		ReturnValue struct {
			Data []struct {
				Name string `json:"name"`
				Id string `json:"id"`
			} `json:"data"`
		} `json:"returnvalue"`
		Status string `json:"status"`
	} `json:"response"`
}
