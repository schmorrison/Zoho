package recruit

import (
	"encoding/json"

	zoho "github.com/schmorrison/Zoho"
)

// GetNotes returns a list of all notes
// https://www.zoho.com/recruit/developer-guide/apiv2/get-notes.html
// https://recruit.zoho.%s/recruit/v2/Notes
func (c *API) GetNotes(params map[string]zohoutils.Parameter) (data NotesResponse, err error) {
	endpoint := zohoutils.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/Notes", c.ZohoTLD),
		Method:       zohoutils.HTTPGet,
		ResponseData: &NotesResponse{},
		URLParameters: map[string]zohoutils.Parameter{
			"page":     "",
			"per_page": "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return NotesResponse{}, fmt.Errorf("failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*NotesResponse); ok {
		return *v, nil
	}

	return NotesResponse{}, fmt.Errorf("data returned was not 'NotesResponse'")
}

// NotesResponse is the data returned by GetNotes
type NotesResponse struct {
	Data []struct {
		IsStatusSplitDone bool   `json:"isStatusSplitDone,omitempty"`
		Modified_Time     Time   `json:"Modified_Time,omitempty"`
		Attachments       string `json:"$attachments,omitempty"`
		Created_Time      Time   `json:"Created_Time,omitempty"`
		ParentID          struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Parent_Id,omitempty"`
		Editable  bool   `json:"$editable,omitempty"`
		SeModule  string `json:"$se_module,omitempty"`
		NoteOwner struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Note_Owner,omitempty"`
		ModifiedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		Size      string `json:"$size,omitempty"`
		VoiceNote bool   `json:"$voice_note,omitempty"`
		ID        string `json:"id,omitempty"`
		CreatedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		NoteTitle      string `json:"Note_Title,omitempty"`
		NoteContent    string `json:"Note_Content,omitempty"`
		IsSystemAction bool   `json:"is_system_action,omitempty"`
	} `json:"data,omitempty"`
	Info struct {
		PerPage     int  `json:"per_page,omitempty"`
		Count       int  `json:"count,omitempty"`
		Page        int  `json:"page,omitempty"`
		MoreRecords bool `json:"more_records,omitempty"`
	} `json:"info,omitempty"`
}
