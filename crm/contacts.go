package crm

import (
	"encoding/xml"
	"time"
)

type ContactsModule CrmModule

func (a *API) Contacts() *ContactsModule {
	return &ContactsModule{
		id:     a.id,
		api:    a,
		module: contactsModule,
	}
}

func (C *ContactsModule) GetMyRecords(o GetRecordsOptions) (Contacts, error) {
	v, err := C.api.getMyRecords(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) GetRecords(o GetRecordsOptions) (Contacts, error) {
	v, err := C.api.getRecords(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) GetRecordsByID(o GetRecordsByIdOptions) (Contacts, error) {
	v, err := C.api.getRecordById(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Contacts, error) {
	v, err := C.api.getDeletedRecordIds(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) InsertRecords(o InsertRecordsOptions) (Contacts, error) {
	v, err := C.api.insertRecords(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) UpdateRecords(o UpdateRecordsOptions) (Contacts, error) {
	v, err := C.api.updateRecords(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) DeleteRecords(id string) (Contacts, error) {
	v, err := C.api.deleteRecords(C.module, id)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Contacts, error) {
	v, err := C.api.getSearchRecordsByPDC(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Contacts, error) {
	v, err := C.api.getRelatedRecords(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Contacts, error) {
	v, err := C.api.updateRelatedRecord(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) GetFields(kind int) (Contacts, error) {
	v, err := C.api.getFields(C.module, kind)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) UploadFile(o UploadFileOptions) (Contacts, error) {
	v, err := C.api.uploadFile(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) DownloadFile(id string) (Contacts, error) {
	v, err := C.api.downloadFile(C.module, id)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) DeleteFile(id string) (Contacts, error) {
	v, err := C.api.deleteFile(C.module, id)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) Delink(o DelinkOptions) (Contacts, error) {
	v, err := C.api.delink(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) UploadPhoto(o UploadPhotoOptions) (Contacts, error) {
	v, err := C.api.uploadPhoto(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) DownloadPhoto(id string) (Contacts, error) {
	v, err := C.api.downloadPhoto(C.module, id)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) DeletePhoto(id string) (Contacts, error) {
	v, err := C.api.deletePhoto(C.module, id)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

func (C *ContactsModule) SearchRecords(o SearchRecordsOptions) (Contacts, error) {
	v, err := C.api.searchRecords(C.module, o)
	if v == nil {
		return Contacts{}, err
	}
	return v.(Contacts), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Contacts
type Contact struct {
	ContactOwner   string    `zoho:"Contact Owner"`    //
	Salutation     string    `zoho:"Salutation"`       //
	FirstName      string    `zoho:"First Name"`       // 40 chars
	LastName       string    `zoho:"Last Name"`        // 40 chars
	AccountName    string    `zoho:"Account Name"`     // 100 chars
	VendorName     string    `zoho:"Vendor Name"`      //
	CampaignSource string    `zoho:"Campaign Source"`  //
	LeadSource     string    `zoho:"Lead Source"`      //
	Title          string    `zoho:"Title"`            // 50 chars
	Department     string    `zoho:"Department"`       // 30 chars
	DateOfBirth    time.Time `zoho:"Date of Birth"`    //
	ReportsTo      string    `zoho:"Reports To"`       // 255 chars
	EmailOptOut    bool      `zoho:"Email Opt Out"`    //
	SkypeID        string    `zoho:"Skype ID"`         // 50 chars
	Phone          string    `zoho:"Phone"`            // 50 chars
	Mobile         string    `zoho:"Mobile"`           // 50 chars
	HomePhone      string    `zoho:"Home Phone"`       // 50 chars
	OtherPhone     string    `zoho:"Other Phone"`      // 50 chars
	Fax            string    `zoho:"Fax"`              // 50 chars
	Email          string    `zoho:"Email"`            // 100 chars
	SecondaryEmail string    `zoho:"Secondary Email"`  // 100 chars
	Assistant      string    `zoho:"Assistant"`        //
	AssistantPhone string    `zoho:"Assistant Phone"`  // 100 chars
	MaillingStreet string    `zoho:"Mailing Street"`   // 250 chars
	MailingCity    string    `zoho:"Mailing City"`     // 30 chars
	MailingState   string    `zoho:"Mailing State"`    // 30 chars
	MailingZipCode string    `zoho:"Mailing Zip Code"` // 30 chars
	MailingCountry string    `zoho:"Mailing Country"`  // 30 chars
	OtherStreet    string    `zoho:"Other Street"`     // 250 chars
	OtherCity      string    `zoho:"Other City"`       // 30 chars
	OtherState     string    `zoho:"Other State"`      // 30 chars
	OtherZipCode   string    `zoho:"Other Zip Code"`   // 30 chars
	OtherCountry   string    `zoho:"Other Country"`    // 30 chars
	Description    string    `zoho:"Description"`      // 32000 chars

	ContactID    int       `zoho:"CONTACTID"`     //
	SmOwnerID    int       `zoho:"SMOWNERID"`     //
	SmCreatorID  int       `zoho:"SMCREATORID"`   //
	CreatedBy    string    `zoho:"Created By"`    //
	ModifiedByID int       `zoho:"MODIFIEDBY"`    //
	ModifiedBy   string    `zoho:"Modified By"`   //
	CreatedTime  time.Time `zoho:"Created Time"`  //
	ModifiedTime time.Time `zoho:"Modified Time"` //
	FullName     string    `zoho:"Full Name"`     //
	Currency     string    `zoho:"Currency"`      //
	ExchangeRate float64   `zoho:"Exchange Rate"` //

	ExtraFields ExtraFields
}

func (c Contact) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Contacts"}}
	x.addRow(c, 1)
	return x.encode()
}

func (c Contact) String() string {
	return c.writeXML()
}

type Contacts []Contact

func (C Contacts) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Contacts"}}
	for i := 0; i < len(C); i++ {
		x.addRow(C[i], i+1)
	}
	return x.encode()
}

func (C Contacts) String() string {
	return C.writeXML()
}
