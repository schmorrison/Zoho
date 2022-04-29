package recruit

import (
	"encoding/json"

	zoho "github.com/schmorrison/Zoho"
)

// https://www.zoho.com/recruit/developer-guide/apiv2/upload-attachment.html
func (c *API) UploadAttachment(file string, params map[string]zohoutils.Parameter, module Module, recordId string) (data UploadAttachmentResponse, err error) {
	endpoint := zohoutils.Endpoint{
		Name:         "filesandattachments",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s/Attachments", c.ZohoTLD, module, recordId),
		Method:       zohoutils.HTTPPost,
		ResponseData: &UploadAttachmentResponse{},
		Attachment:   file,
		BodyFormat:   zohoutils.FILE,
		URLParameters: map[string]zohoutils.Parameter{
			"attachments_category_id": "",
			"attachments_category":    "",
			"attachment_url":          "",
		},
	}

	for k, v := range params {
		endpoint.URLParameters[k] = v
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UploadAttachmentResponse{}, fmt.Errorf("failed to upload Attachment: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UploadAttachmentResponse); ok {
		return *v, nil
	}

	return UploadAttachmentResponse{}, fmt.Errorf("data returned was nil")
}

type UploadAttachmentResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ModifiedTime time.Time `json:"Modified_Time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"Modified_By"`
			CreatedTime time.Time `json:"Created_Time"`
			ID          string    `json:"id"`
			CreatedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"Created_By"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}
