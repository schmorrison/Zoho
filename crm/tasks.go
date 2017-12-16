package crm

import (
	"encoding/xml"
	"time"
)

type TasksModule CrmModule

func (a *API) Tasks() *TasksModule {
	return &TasksModule{
		id:     a.id,
		api:    a,
		module: tasksModule,
	}
}

func (T *TasksModule) GetMyRecords(o GetRecordsOptions) (Tasks, error) {
	v, err := T.api.getMyRecords(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) GetRecords(o GetRecordsOptions) (Tasks, error) {
	v, err := T.api.getRecords(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) GetRecordsByID(o GetRecordsByIdOptions) (Tasks, error) {
	v, err := T.api.getRecordById(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Tasks, error) {
	v, err := T.api.getDeletedRecordIds(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) InsertRecords(o InsertRecordsOptions) (Tasks, error) {
	v, err := T.api.insertRecords(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) UpdateRecords(o UpdateRecordsOptions) (Tasks, error) {
	v, err := T.api.updateRecords(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) DeleteRecords(id string) (Tasks, error) {
	v, err := T.api.deleteRecords(T.module, id)
	return v.(Tasks), err
}

func (T *TasksModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Tasks, error) {
	v, err := T.api.getSearchRecordsByPDC(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Tasks, error) {
	v, err := T.api.getRelatedRecords(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Tasks, error) {
	v, err := T.api.updateRelatedRecord(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) GetFields(kind int) (Tasks, error) {
	v, err := T.api.getFields(T.module, kind)
	return v.(Tasks), err
}

func (T *TasksModule) UploadFile(o UploadFileOptions) (Tasks, error) {
	v, err := T.api.uploadFile(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) DownloadFile(id string) (Tasks, error) {
	v, err := T.api.downloadFile(T.module, id)
	return v.(Tasks), err
}

func (T *TasksModule) DeleteFile(id string) (Tasks, error) {
	v, err := T.api.deleteFile(T.module, id)
	return v.(Tasks), err
}

func (T *TasksModule) Delink(o DelinkOptions) (Tasks, error) {
	v, err := T.api.delink(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) UploadPhoto(o UploadPhotoOptions) (Tasks, error) {
	v, err := T.api.uploadPhoto(T.module, o)
	return v.(Tasks), err
}

func (T *TasksModule) DownloadPhoto(id string) (Tasks, error) {
	v, err := T.api.downloadPhoto(T.module, id)
	return v.(Tasks), err
}

func (T *TasksModule) DeletePhoto(id string) (Tasks, error) {
	v, err := T.api.deletePhoto(T.module, id)
	return v.(Tasks), err
}

func (T *TasksModule) SearchRecords(o SearchRecordsOptions) (Tasks, error) {
	v, err := T.api.searchRecords(T.module, o)
	return v.(Tasks), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Tasks
type Task struct {
	TaskOwner             string    `zoho:"Task Owner"`              //
	Subject               string    `zoho:"Subject"`                 // 50 chars
	DueDate               time.Time `zoho:"Due Date"`                //
	ContactLead           string    `zoho:"Contact Lead"`            //
	Accounts              string    `zoho:"Accounts"`                //
	Status                string    `zoho:"Status"`                  //
	Priority              string    `zoho:"Priority"`                //
	SendNotificationEmail bool      `zoho:"Send Notification Email"` //
	RemindAt              bool      `zoho:"Remind At"`               //
	RecurringActivity     bool      `zoho:"Recurring Activity"`      //
	Description           string    `zoho:"Description"`             // 32000 chars
	Currency              string    `zoho:"Currency"`                //
	ExchangeRate          float64   `zoho:"Exchange Rate"`           //

	//Fields not documented but returned by some methods
	//Some of these fields may need to be stripped due to zoho requirements
	ActivityID   int       `zoho:"ACTIVITYID"`
	SmOwnerID    int       `zoho:"SMOWNERID"`
	RelatedToID  int       `zoho:"RELATEDTOID"`
	ContactID    int       `zoho:"CONTACTID"`
	ContactName  string    `zoho:"Contact Name"`
	SeModule     string    `zoho:"SEMODULE"`
	RelatedTo    string    `zoho:"Related To"`
	SmCreatorID  int       `zoho:"SMCREATORID"`
	CreatedBy    string    `zoho:"Created By"`
	ModifiedByID int       `zoho:"MODIFIEDBY"`
	ModifiedBy   string    `zoho:"Modified By"`
	CreatedTime  time.Time `zoho:"Created Time"`
	ModifiedTime time.Time `zoho:"Modified Time"`
	ClosedTime   time.Time `zoho:"Closed Time"`

	ExtraFields ExtraFields
}

func (t Task) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Tasks"}}
	x.addRow(t, 1)
	return x.encode()
}

func (t Task) String() string {
	return t.writeXML()
}

type Tasks []Task

func (T Tasks) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Tasks"}}
	for i := 0; i < len(T); i++ {
		x.addRow(T[i], i+1)
	}
	return x.encode()
}

func (T Tasks) String() string {
	return T.writeXML()
}
