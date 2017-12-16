package crm

import (
	"encoding/xml"
	"time"
)

type CallsModule CrmModule

func (a *API) Calls() *CallsModule {
	return &CallsModule{
		id:     a.id,
		api:    a,
		module: callsModule,
	}
}

func (C *CallsModule) GetMyRecords(o GetRecordsOptions) (Calls, error) {
	v, err := C.api.getMyRecords(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) GetRecords(o GetRecordsOptions) (Calls, error) {
	v, err := C.api.getRecords(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) GetRecordsByID(o GetRecordsByIdOptions) (Calls, error) {
	v, err := C.api.getRecordById(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Calls, error) {
	v, err := C.api.getDeletedRecordIds(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) InsertRecords(o InsertRecordsOptions) (Calls, error) {
	v, err := C.api.insertRecords(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) UpdateRecords(o UpdateRecordsOptions) (Calls, error) {
	v, err := C.api.updateRecords(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) DeleteRecords(id string) (Calls, error) {
	v, err := C.api.deleteRecords(C.module, id)
	return v.(Calls), err
}

func (C *CallsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Calls, error) {
	v, err := C.api.getSearchRecordsByPDC(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Calls, error) {
	v, err := C.api.getRelatedRecords(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Calls, error) {
	v, err := C.api.updateRelatedRecord(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) GetFields(kind int) (Calls, error) {
	v, err := C.api.getFields(C.module, kind)
	return v.(Calls), err
}

func (C *CallsModule) UploadFile(o UploadFileOptions) (Calls, error) {
	v, err := C.api.uploadFile(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) DownloadFile(id string) (Calls, error) {
	v, err := C.api.downloadFile(C.module, id)
	return v.(Calls), err
}

func (C *CallsModule) DeleteFile(id string) (Calls, error) {
	v, err := C.api.deleteFile(C.module, id)
	return v.(Calls), err
}

func (C *CallsModule) Delink(o DelinkOptions) (Calls, error) {
	v, err := C.api.delink(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) UploadPhoto(o UploadPhotoOptions) (Calls, error) {
	v, err := C.api.uploadPhoto(C.module, o)
	return v.(Calls), err
}

func (C *CallsModule) DownloadPhoto(id string) (Calls, error) {
	v, err := C.api.downloadPhoto(C.module, id)
	return v.(Calls), err
}

func (C *CallsModule) DeletePhoto(id string) (Calls, error) {
	v, err := C.api.deletePhoto(C.module, id)
	return v.(Calls), err
}

func (C *CallsModule) SearchRecords(o SearchRecordsOptions) (Calls, error) {
	v, err := C.api.searchRecords(C.module, o)
	return v.(Calls), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Calls
type Call struct {
	Subject             string    `zoho:"Subject"`         // 50 chars
	CallType            string    `zoho:"Call Type"`       // inbound or outbound
	CallPurpose         string    `zoho:"Call Purpose"`    //
	CallFromTo          string    `zoho:"Call From To"`    //
	RelatedTo           string    `zoho:"Related To"`      //
	CallDetails         string    `zoho:"Call Details"`    // currentCall or completedCall
	CallStartTime       time.Time `zoho:"Call Start Time"` //
	CallDuration        string    `zoho:"Call Duration"`
	CallDurationSeconds int       `zoho:"Call Duration (in seconds)"` //
	Description         string    `zoho:"Description"`                // 32000 chars
	Billable            bool      `zoho:"Billable"`                   //
	CallResult          string    `zoho:"Call Result"`                //

	//Fields not documented but returned by some methods
	//Some of these fields may need to be stripped due to zoho requirements
	ActivityID   int       `zoho:"ACTIVITYID"`
	SmOwnerID    int       `zoho:"SMOWNERID"`
	CallOwner    string    `zoho:"Call Owner"`
	RelatedToID  int       `zoho:"RELATEDTOID"`
	SeModule     string    `zoho:"SEMODULE"`
	CreatedBy    string    `zoho:"SMCREATORID"`
	SmCreatorID  int       `zoho:"Created By"`
	ModifiedByID int       `zoho:"MODIFIEDBY"`
	ModifiedBy   string    `zoho:"Modified By"`
	CreatedTime  time.Time `zoho:"Created Time"`
	ModifiedTime time.Time `zoho:"Modified Time"`
	CallStatus   string    `zoho:"Call Status"`

	ExtraFields ExtraFields
}

func (c Call) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Calls"}}
	x.addRow(c, 1)
	return x.encode()
}

func (c Call) String() string {
	return c.writeXML()
}

type Calls []Call

func (C Calls) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Calls"}}
	for i := 0; i < len(C); i++ {
		x.addRow(C[i], i+1)
	}
	return x.encode()
}

func (C Calls) String() string {
	return C.writeXML()
}
