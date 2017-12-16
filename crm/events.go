package crm

import (
	"encoding/xml"
	"time"
)

type EventsModule CrmModule

func (a *API) Events() *EventsModule {
	return &EventsModule{
		id:     a.id,
		api:    a,
		module: eventsModule,
	}
}

func (E *EventsModule) GetMyRecords(o GetRecordsOptions) (Events, error) {
	v, err := E.api.getMyRecords(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) GetRecords(o GetRecordsOptions) (Events, error) {
	v, err := E.api.getRecords(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) GetRecordsByID(o GetRecordsByIdOptions) (Events, error) {
	v, err := E.api.getRecordById(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Events, error) {
	v, err := E.api.getDeletedRecordIds(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) InsertRecords(o InsertRecordsOptions) (Events, error) {
	v, err := E.api.insertRecords(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) UpdateRecords(o UpdateRecordsOptions) (Events, error) {
	v, err := E.api.updateRecords(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) DeleteRecords(id string) (Events, error) {
	v, err := E.api.deleteRecords(E.module, id)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Events, error) {
	v, err := E.api.getSearchRecordsByPDC(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Events, error) {
	v, err := E.api.getRelatedRecords(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Events, error) {
	v, err := E.api.updateRelatedRecord(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) GetFields(kind int) (Events, error) {
	v, err := E.api.getFields(E.module, kind)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) UploadFile(o UploadFileOptions) (Events, error) {
	v, err := E.api.uploadFile(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) DownloadFile(id string) (Events, error) {
	v, err := E.api.downloadFile(E.module, id)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) DeleteFile(id string) (Events, error) {
	v, err := E.api.deleteFile(E.module, id)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) Delink(o DelinkOptions) (Events, error) {
	v, err := E.api.delink(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) UploadPhoto(o UploadPhotoOptions) (Events, error) {
	v, err := E.api.uploadPhoto(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) DownloadPhoto(id string) (Events, error) {
	v, err := E.api.downloadPhoto(E.module, id)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) DeletePhoto(id string) (Events, error) {
	v, err := E.api.deletePhoto(E.module, id)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

func (E *EventsModule) SearchRecords(o SearchRecordsOptions) (Events, error) {
	v, err := E.api.searchRecords(E.module, o)
	if v == nil {
		return Events{}, err
	}
	return v.(Events), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Events
type Event struct {
	EventOwner            string    `zoho:"Event Owner"`             //
	Subject               string    `zoho:"Subject"`                 // 50 chars
	StartDateTime         time.Time `zoho:"Start DateTime"`          //
	EndDateTime           time.Time `zoho:"End DateTime"`            //
	Venue                 string    `zoho:"Venue"`                   //
	ContactLead           string    `zoho:"Contact Lead"`            //
	Accounts              string    `zoho:"Accounts"`                //
	SendNotificationEmail bool      `zoho:"Send Notification Email"` //
	RemindAt              bool      `zoho:"Remind At"`               //
	RecurringActivity     bool      `zoho:"Recurring Activity"`      //
	Description           string    `zoho:"Description"`             // 32000 chars
	Currency              string    `zoho:"Currency"`                //
	ExchangeRate          float64   `zoho:"Exchange Rate"`           //

	//Fields not documented but returned by some methods
	//Some of these fields may need to be stripped due to zoho requirements
	ActivityID   int       `zoho:"ACTIVITYID"`    //
	UID          string    `zoho:"UID"`           //
	SmOwnerID    int       `zoho:"SMOWNERID"`     //
	SmCreatorID  int       `zoho:"Created By"`    //
	CreatedBy    string    `zoho:"SMCREATORID"`   //
	ModifiedByID int       `zoho:"MODIFIEDBY"`    //
	ModifiedBy   string    `zoho:"Modified By"`   //
	CreatedTime  time.Time `zoho:"Created Time"`  //
	ModifiedTime time.Time `zoho:"Modified Time"` //
	ContactID    int       `zoho:"CONTACTID"`     //
	ContactName  string    `zoho:"Contact Name"`  //
	RelatedToID  int       `zoho:"RELATEDTOID"`   //
	SeModule     string    `zoho:"SEMODULE"`      //
	RelatedTo    string    `zoho:"Related To"`    //

	ExtraFields ExtraFields
}

func (e Event) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Events"}}
	x.addRow(e, 1)
	return x.encode()
}

func (e Event) String() string {
	return e.writeXML()
}

type Events []Event

func (E Events) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Events"}}
	for i := 0; i < len(E); i++ {
		x.addRow(E[i], i+1)
	}
	return x.encode()
}

func (E Events) String() string {
	return E.writeXML()
}
