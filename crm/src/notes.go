package main

var NotesSchema = EndpointSchema{
	Name: "notes",
	EndpointList: []EndpointSchemaBase{
		EndpointSchemaBase{
			Comments: []string{
				"GetNotes returns a list of all notes",
				"https://www.zoho.com/crm/help/api/v2/#notes-api",
			},
			Endpoint: &Endpoint{
				Name:   "GetNotes",
				Method: "GET",
				URI:    "/crm/v2/Notes",
			},
			Params: EndpointParams{
				EndpointParam{
					Name: "page",
					Type: "",
					Map:  true,
				},
				EndpointParam{
					Name: "per_page",
					Type: "200",
					Map:  true,
				},
			},
			Response: "NotesResponse",
		},
		EndpointSchemaBase{
			Comments: []string{
				"GetNote returns the note specified by ID and module",
				"https://www.zoho.com/crm/help/api/v2/#get-spec-notes-data",
			},
			Endpoint: &Endpoint{
				Name:   "GetNote",
				Method: "GET",
				URI:    "/crm/v2/%s/%s/Notes",
			},
			Params: EndpointParams{
				EndpointParam{
					Name: "module",
					Type: "crmModule",
					Func: true,
					URI:  true,
				},
				EndpointParam{
					Name: "id",
					Type: "string",
					Func: true,
					URI:  true,
				},
			},
			Response: "NoteResponse",
		},
		EndpointSchemaBase{
			Comments: []string{
				"CreateNotes will create multiple notes provided in the request data",
				"https://www.zoho.com/crm/help/api/v2/#create-notes",
			},
			Endpoint: &Endpoint{
				Name:   "CreateNotes",
				Method: "POST",
				URI:    "/crm/v2/Notes",
			},
			Params:   EndpointParams{},
			Request:  "CreateNoteData",
			Response: "CreateNoteResponse",
		},
		EndpointSchemaBase{
			Comments: []string{
				"UpdateNote will update the note data of the specified note on the specified record of the module",
				"https://www.zoho.com/crm/help/api/v2/#update-notes",
			},
			Endpoint: &Endpoint{
				Name:   "UpdateNote",
				Method: "PUT",
				URI:    "/crm/v2/%s/%s/Notes/%s",
			},
			Params: EndpointParams{
				EndpointParam{
					Name: "module",
					Type: "crmModule",
					Func: true,
					URI:  true,
				},
				EndpointParam{
					Name: "recordID",
					Type: "string",
					Func: true,
					URI:  true,
				},
				EndpointParam{
					Name: "noteID",
					Type: "string",
					Func: true,
					URI:  true,
				},
			},
			Response: "UpdateNoteResponse",
			Request:  "UpdateNoteData",
		},
		EndpointSchemaBase{
			Comments: []string{
				"DeleteNote will delete the specified note on the specified record from the module",
				"https://www.zoho.com/crm/help/api/v2/#delete-notes",
			},
			Endpoint: &Endpoint{
				Name:   "DeleteNote",
				Method: "DELETE",
				URI:    "/crm/v2/%s/%s/Notes/%s",
			},
			Params: EndpointParams{
				EndpointParam{
					Name: "module",
					Type: "crmModule",
					Func: true,
					URI:  true,
				},
				EndpointParam{
					Name: "recordID",
					Type: "string",
					Func: true,
					URI:  true,
				},
				EndpointParam{
					Name: "noteID",
					Type: "string",
					Func: true,
					URI:  true,
				},
			},
			Response: "DeleteNoteResponse",
		},
		EndpointSchemaBase{
			Comments: []string{
				"DeleteNotes will delete all notes specified in the IDs",
				"https://www.zoho.com/crm/help/api/v2/#delete-bulk-notes",
			},
			Endpoint: &Endpoint{
				Name:   "DeleteNotes",
				Method: "DELETE",
				URI:    "/crm/v2/Notes",
			},
			Params: EndpointParams{
				EndpointParam{
					Name: "IDs",
					Type: "...string",
					Func: true,
					Map:  true,
				},
			},
			Response: "DeleteNotesResponse",
		},
	},
}
