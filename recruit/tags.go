package recruit

import (
	"fmt"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

// https://www.zoho.com/recruit/developer-guide/apiv2/create-tag.html
func (c *API) CreateTags(request CreateTagsRequest, params map[string]zoho.Parameter) (data CreateTagsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "CreateTags",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/settings/tags", c.ZohoTLD),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateTagsResponse{},
		RequestBody:  request,
		BodyFormat:   zoho.JSON,
		URLParameters: map[string]zoho.Parameter{
			"module": "", // mandatory
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	// pp.Printf("%s\n", endpoint)

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateTagsResponse{}, fmt.Errorf("failed to create Tag(s): %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateTagsResponse); ok {
		for _, resp := range v.Tags {
			if resp.Code != "SUCCESS" {
				return CreateTagsResponse{}, fmt.Errorf("failed to create Tag(s): %s: %s", resp.Code, resp.Message)
			}
		}
		return *v, nil
	}

	return CreateTagsResponse{}, fmt.Errorf("data returned was not 'CreateTagsResponse'")
}

type CreateTagsRequest struct {
	Tags []Tags `json:"tags"`
}
type Tags struct {
	Name string `json:"name"`
}

type CreateTagsResponse struct {
	Tags []struct {
		Code    string `json:"code"`
		Details struct {
			CreatedTime  time.Time `json:"created_time"`
			ModifiedTime time.Time `json:"modified_time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"modified_by"`
			ID        string `json:"id"`
			CreatedBy struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"created_by"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"tags"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/add-tags.html
func (c *API) AddTagsToIDs(module Module, params map[string]zoho.Parameter) (data AddTagsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "AddTagsToIDs",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/actions/add_tags", c.ZohoTLD, module),
		Method:       zoho.HTTPPost,
		ResponseData: &AddTagsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"tag_names": "",
			"ids":       "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AddTagsResponse{}, fmt.Errorf("failed to insert Tag(s): %s", err)
	}

	if v, ok := endpoint.ResponseData.(*AddTagsResponse); ok {
		return *v, nil
	}

	return AddTagsResponse{}, fmt.Errorf("data returned was not 'AddTagsResponse'")
}

// https://www.zoho.com/recruit/developer-guide/apiv2/add-tags.html
func (c *API) AddTagsToId(module Module, ID string, params map[string]zoho.Parameter) (data AddTagsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "AddTagsToId",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s/actions/add_tags", c.ZohoTLD, module, ID),
		Method:       zoho.HTTPPost,
		ResponseData: &AddTagsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"tag_names": "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AddTagsResponse{}, fmt.Errorf("failed to add Tag(s): %s", err)
	}

	if v, ok := endpoint.ResponseData.(*AddTagsResponse); ok {
		for _, resp := range v.Data {
			if resp.Code != "SUCCESS" {
				return AddTagsResponse{}, fmt.Errorf("failed to add Tag(s): %s: %s", resp.Code, resp.Message)
			}
		}
		return *v, nil
	}

	return AddTagsResponse{}, fmt.Errorf("data returned was not 'AddTagsResponse'")
}

type AddTagsResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ID   int64    `json:"id"`
			Tags []string `json:"tags"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/delete-tag.html
func (c *API) DeleteTagById(tagID string) (data DeleteTagResponse, err error) {
	if len(tagID) == 0 {
		return DeleteTagResponse{}, fmt.Errorf("failed to delete Tag, must provide tagID")
	}

	endpoint := zoho.Endpoint{
		Name:         "DeleteTagById",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/settings/tags/%s", c.ZohoTLD, tagID),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteTagResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteTagResponse{}, fmt.Errorf("failed to delete Tag: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteTagResponse); ok {
		return *v, nil
	}

	return DeleteTagResponse{}, fmt.Errorf("data retrieved was not 'DeleteTagResponse'")
}

type DeleteTagResponse struct {
	Tags struct {
		Code    string `json:"code"`
		Details struct {
			ID int64 `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"tags"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/get-tag-list.html
func (c *API) GetTagsList(module Module, params map[string]zoho.Parameter) (data TagsListResponse, err error) {
	if len(module) == 0 {
		return TagsListResponse{}, fmt.Errorf("failed to list Tags, module name is missing")
	}
	endpoint := zoho.Endpoint{
		Name:         "GetTagsList",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/settings/tags", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &TagsListResponse{},
		URLParameters: map[string]zoho.Parameter{
			"module":  zoho.Parameter(module), // mandatory
			"my_tags": "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	// log.Printf("endpoint: %+v", endpoint)

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return TagsListResponse{}, fmt.Errorf("failed to retrieve %s TagsList: %s", params["module"], err)
	}

	if v, ok := endpoint.ResponseData.(*TagsListResponse); ok {
		return *v, nil
	}

	return TagsListResponse{}, fmt.Errorf("data returned was not 'TagsListResponse'")
}

type TagsListResponse struct {
	Data struct {
		Tags []struct {
			CreatedTime  time.Time `json:"created_time"`
			ModifiedTime time.Time `json:"modified_time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"modified_by"`
			Name      string `json:"name"`
			ID        string `json:"id"`
			CreatedBy struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"created_by"`
		} `json:"tags"`
		Info PageInfo `json:"info"`
	} `json:"data"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/update-tags.html
func (c *API) UpdateTag(ID string, request UpdateTagRequest) (data UpdateTagResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "UpdateTag",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/settings/tags/%s", c.ZohoTLD, ID),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateTagResponse{},
		RequestBody:  request,
		BodyFormat:   zoho.JSON,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateTagResponse{}, fmt.Errorf("failed to update Tag: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateTagResponse); ok {
		return *v, nil
	}

	return UpdateTagResponse{}, fmt.Errorf("data returned was not 'UpdateTagResponse'")
}

type UpdateTagRequest struct {
	Tags []struct {
		Name string `json:"name"`
	} `json:"tags"`
}
type UpdateTagResponse struct {
	Tags []struct {
		Code    string `json:"code"`
		Details struct {
			CreatedTime  time.Time `json:"created_time"`
			ModifiedTime time.Time `json:"modified_time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"modified_by"`
			Name      string `json:"name"`
			ID        int64  `json:"id"`
			CreatedBy struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"created_by"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"tags"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/remove-tags.html
func (c *API) RemoveTagsFromIDs(module Module, params map[string]zoho.Parameter) (data RemoveTagsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "RemoveTagsFromIDs",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/actions/remove_tags", c.ZohoTLD, module),
		Method:       zoho.HTTPPost,
		ResponseData: &RemoveTagsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"tag_names": "",
			"ids":       "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return RemoveTagsResponse{}, fmt.Errorf("failed to insert Tag(s): %s", err)
	}

	if v, ok := endpoint.ResponseData.(*RemoveTagsResponse); ok {
		return *v, nil
	}

	return RemoveTagsResponse{}, fmt.Errorf("data returned was not 'RemoveTagsResponse'")
}

// https://www.zoho.com/recruit/developer-guide/apiv2/remove-tags.html
func (c *API) RemoveTagsFromId(module Module, ID string, params map[string]zoho.Parameter) (data RemoveTagsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "RemoveTagsFromId",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s/actions/remove_tags", c.ZohoTLD, module, ID),
		Method:       zoho.HTTPPost,
		ResponseData: &RemoveTagsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"tag_names": "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return RemoveTagsResponse{}, fmt.Errorf("failed to remove Tag(s): %s", err)
	}

	if v, ok := endpoint.ResponseData.(*RemoveTagsResponse); ok {
		for _, resp := range v.Data {
			if resp.Code != "SUCCESS" {
				return RemoveTagsResponse{}, fmt.Errorf("failed to remove Tag(s): %s: %s", resp.Code, resp.Message)
			}
		}
		return *v, nil
	}

	return RemoveTagsResponse{}, fmt.Errorf("data returned was not 'RemoveTagsResponse'")
}

type RemoveTagsResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ID   int64    `json:"id"`
			Tags []string `json:"tags"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}
