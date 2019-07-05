package crm

import (
	"fmt"

	"github.com/schmorrison/Zoho"
)

// GetNotes returns a list of all notes
// https://www.zoho.com/crm/help/api/v2/#notes-api
func (c *API) GetNotes(params map[string]zoho.Parameter) (data NotesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/Notes", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &NotesResponse{},
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

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return NotesResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*NotesResponse); ok {
		return *v, nil
	}

	return NotesResponse{}, fmt.Errorf("Data returned was not 'NotesResponse'")
}

// GetNote returns the note specified by ID and module
// https://www.zoho.com/crm/help/api/v2/#get-spec-notes-data
func (c *API) GetNote(module crmModule, id string) (data NotesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/Notes", c.ZohoTLD, module, id),
		Method:       zoho.HTTPGet,
		ResponseData: &NotesResponse{},
	}
	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return NotesResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*NotesResponse); ok {
		return *v, nil
	}

	return NotesResponse{}, fmt.Errorf("Data returned was not 'NotesResponse'")
}

// NotesResponse is the data returned by GetNotes and GetNote
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
	Info PageInfo `json:"info,omitempty"`
}

// CreateNotes will create multiple notes provided in the request data
// https://www.zoho.com/crm/help/api/v2/#create-notes
func (c *API) CreateNotes(request CreateNoteData) (data CreateNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/Notes", c.ZohoTLD),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateNoteResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateNoteResponse{}, fmt.Errorf("Failed to create notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateNoteResponse); ok {
		return *v, nil
	}

	return CreateNoteResponse{}, fmt.Errorf("Data returned was not 'CreateNoteResponse'")
}

// CreateNoteData is the data provided to create 1 or more notes
type CreateNoteData struct {
	Data []struct {
		NoteTitle   string `json:"Note_Title,omitempty"`
		NoteContent string `json:"Note_Content,omitempty"`
		ParentID    string `json:"Parent_Id,omitempty"`
		SeModule    string `json:"se_module,omitempty"`
	} `json:"data,omitempty"`
}

// CreateNoteResponse is the data returned by CreateNotes
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

// CreateRecordNote will create a note on the specified record of the specified module
// https://www.zoho.com/crm/help/api/v2/#create-spec-notes
func (c *API) CreateRecordNote(request CreateRecordNoteData, module crmModule, recordID string) (data CreateRecordNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/Notes", c.ZohoTLD, module, recordID),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateRecordNoteResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateRecordNoteResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateRecordNoteResponse); ok {
		return *v, nil
	}

	return CreateRecordNoteResponse{}, fmt.Errorf("Data returned was not 'CreateRecordNoteResponse'")
}

// CreateRecordNoteResponse is the data returned by CreateRecordNote, it is the same as the data returned by CreateNote
type CreateRecordNoteResponse = CreateNoteResponse

// CreateRecordNoteData is the data returned by CreateRecordNote
type CreateRecordNoteData struct {
	Data []struct {
		NoteTitle   string `json:"Note_Title,omitempty"`
		NoteContent string `json:"Note_Content,omitempty"`
	} `json:"data,omitempty"`
}

// UpdateNote will update the note data of the specified note on the specified record of the module
// https://www.zoho.com/crm/help/api/v2/#update-notes
func (c *API) UpdateNote(request UpdateNoteData, module crmModule, recordID, noteID string) (data UpdateNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/Notes/%s", c.ZohoTLD, module, recordID, noteID),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateNoteResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateNoteResponse{}, fmt.Errorf("Failed to update notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateNoteResponse); ok {
		return *v, nil
	}

	return UpdateNoteResponse{}, fmt.Errorf("Data returned was not 'UpdateNoteResponse'")
}

// UpdateNoteResponse is the data returned by UpdateNote
type UpdateNoteResponse = CreateNoteResponse

// UpdateNoteData is the data required by UpdateNote
type UpdateNoteData = CreateRecordNoteData

// DeleteNote will delete the specified note on the specified record from the module
// https://www.zoho.com/crm/help/api/v2/#delete-notes
func (c *API) DeleteNote(module crmModule, recordID, noteID string) (data DeleteNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/Notes/%s", c.ZohoTLD, module, recordID, noteID),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteNoteResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteNoteResponse{}, fmt.Errorf("Failed to delete note: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteNoteResponse); ok {
		return *v, nil
	}

	return DeleteNoteResponse{}, fmt.Errorf("Data returned was not 'DeleteNoteResponse'")
}

// DeleteNotes will delete all notes specified in the IDs
// https://www.zoho.com/crm/help/api/v2/#delete-bulk-notes
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
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/Notes", c.ZohoTLD),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteNoteResponse{},
		URLParameters: map[string]zoho.Parameter{
			"ids": "",
		},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteNoteResponse{}, fmt.Errorf("Failed to delete notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteNoteResponse); ok {
		return *v, nil
	}

	return DeleteNoteResponse{}, fmt.Errorf("Data returned was not 'DeleteNoteResponse'")
}

// DeleteNoteResponse is the data returned when deleting a note
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
