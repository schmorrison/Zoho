package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) FetchWorkspaces(request interface{}, params map[string]zoho.Parameter) (data WorkspaceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         FetchWorkspacesModule,
		URL:          fmt.Sprintf(BookingsAPIEndpoint+"%s", FetchWorkspacesModule),
		Method:       zoho.HTTPGet,
		ResponseData: &WorkspaceResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}
	if len(params) != 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
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
	response struct {
		returnValue struct {
			data []struct {
				name string `json:"name"`
				id string `json:"id"`
			} `json:"data"`
		} `json:"returnvalue"`
	} `json:"response"`
}
