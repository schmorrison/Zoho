package crm

import (
	"fmt"

	".."
)

func (c *API) GetNotes(params map[string]zoho.Parameter) (data NotesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          "https://www.zohoapis.com/crm/v2/Notes",
		Method:       zoho.HTTPGet,
		ResponseData: NotesResponse{},
		URLParameters: map[string]zoho.Parameter{
			"page":     "",
			"per_page": "200",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return NotesResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(NotesResponse); ok {
		return v, nil
	}

	return NotesResponse{}, fmt.Errorf("Data returned was not 'NotesResponse'")
}

func (c *API) GetNote(module crmModule, recordID string) (data NotesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.com/crm/v2/%s/%s/Notes", module, recordID),
		Method:       zoho.HTTPGet,
		ResponseData: NotesResponse{},
	}
	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return NotesResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(NotesResponse); ok {
		return v, nil
	}

	return NotesResponse{}, fmt.Errorf("Data returned was not 'NotesResponse'")
}

type NotesResponse struct {
	Data []struct {
		Owner struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Owner,omitempty"`
		SeModule string `json:"$se_module,omitempty"`
		Approval struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
		} `json:"$approval,omitempty"`
		ModifiedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		ModifiedTime Time `json:"Modified_Time,omitempty"`
		CreatedTime  Time `json:"Created_Time,omitempty"`
		Followed     bool `json:"$followed,omitempty"`
		ParentID     struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Parent_Id,omitempty"`
		ID        string `json:"id,omitempty"`
		CreatedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		NoteTitle   string `json:"Note_Title,omitempty"`
		NoteContent string `json:"Note_Content,omitempty"`
	} `json:"data,omitempty"`
	Info struct {
		PerPage     int  `json:"per_page,omitempty"`
		Count       int  `json:"count,omitempty"`
		Page        int  `json:"page,omitempty"`
		MoreRecords bool `json:"more_records,omitempty"`
	} `json:"info,omitempty"`
}

func (c *API) CreateNote(input CreateNoteData) (data CreateNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          "https://www.zohoapis.com/crm/v2/Notes",
		Method:       zoho.HTTPPost,
		ResponseData: CreateNoteResponse{},
		RequestBody:  input,
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return CreateNoteResponse{}, fmt.Errorf("Failed to create notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(CreateNoteResponse); ok {
		return v, nil
	}

	return CreateNoteResponse{}, fmt.Errorf("Data returned was not 'CreateNoteResponse'")
}

type CreateNoteData struct {
	Data []struct {
		NoteTitle   string `json:"Note_Title,omitempty"`
		NoteContent string `json:"Note_Content,omitempty"`
		ParentID    string `json:"Parent_Id,omitempty"`
		SeModule    string `json:"se_module,omitempty"`
	} `json:"data,omitempty"`
}

type CreateNoteResponse struct {
	Data []struct {
		Message string `json:"message,omitempty"`
		Details struct {
			CreatedBy struct {
				ID   string `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"created_by,omitempty"`
			ID         string `json:"id,omitempty"`
			ModifiedBy struct {
				ID   string `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"modified_by,omitempty"`
			ModifiedTime Time `json:"modified_time,omitempty"`
			CreatedTime  Time `json:"created_time,omitempty"`
		} `json:"details,omitempty"`
		Status string `json:"status,omitempty"`
		Code   string `json:"code,omitempty"`
	} `json:"data,omitempty"`
}

func (c *API) CreateRecordNote(input CreateRecordNoteData, module crmModule, recordID string) (data CreateRecordNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.com/crm/v2/%s/%s/Notes", module, recordID),
		Method:       zoho.HTTPPost,
		ResponseData: CreateRecordNoteResponse{},
		RequestBody:  input,
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return CreateRecordNoteResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(CreateRecordNoteResponse); ok {
		return v, nil
	}

	return CreateRecordNoteResponse{}, fmt.Errorf("Data returned was not 'CreateRecordNoteResponse'")
}

type CreateRecordNoteResponse = CreateNoteResponse

type CreateRecordNoteData struct {
	Data []struct {
		NoteTitle   string `json:"Note_Title,omitempty"`
		NoteContent string `json:"Note_Content,omitempty"`
	} `json:"data,omitempty"`
}

func (c *API) UpdateNote(input UpdateNoteData, module crmModule, recordID, noteID string) (data UpdateNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.com/crm/v2/%s/%s/Notes/%s", module, recordID, noteID),
		Method:       zoho.HTTPPut,
		ResponseData: UpdateNoteResponse{},
		RequestBody:  input,
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return UpdateNoteResponse{}, fmt.Errorf("Failed to update notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(UpdateNoteResponse); ok {
		return v, nil
	}

	return UpdateNoteResponse{}, fmt.Errorf("Data returned was not 'UpdateNoteResponse'")
}

type UpdateNoteResponse = CreateNoteResponse
type UpdateNoteData = CreateRecordNoteData

func (c *API) DeleteNote(module crmModule, recordID, noteID string) (data DeleteNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.com/crm/v2/%s/%s/Notes/%s", module, recordID, noteID),
		Method:       zoho.HTTPDelete,
		ResponseData: DeleteNoteResponse{},
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return DeleteNoteResponse{}, fmt.Errorf("Failed to delete note: %s", err)
	}

	if v, ok := endpoint.ResponseData.(DeleteNoteResponse); ok {
		return v, nil
	}

	return DeleteNoteResponse{}, fmt.Errorf("Data returned was not 'DeleteNoteResponse'")
}

func (c *API) DeleteNotes(IDs ...string) (data DeleteNoteResponse, err error) {
	idStr := ""
	for i, a := range IDs {
		idStr += a
		if i < len(IDs)-1 {
			idStr += ","
		}
	}
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          "https://www.zohoapis.com/crm/v2/Notes",
		Method:       zoho.HTTPDelete,
		ResponseData: DeleteNoteResponse{},
		URLParameters: map[string]zoho.Parameter{
			"ids": "",
		},
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return DeleteNoteResponse{}, fmt.Errorf("Failed to delete notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(DeleteNoteResponse); ok {
		return v, nil
	}

	return DeleteNoteResponse{}, fmt.Errorf("Data returned was not 'DeleteNoteResponse'")
}

type DeleteNoteResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ID string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}
