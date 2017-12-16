package crm

import (
	"encoding/xml"
	"time"
)

type LeadsModule CrmModule

func (a *API) Leads() *LeadsModule {
	return &LeadsModule{
		id:     a.id,
		api:    a,
		module: leadsModule,
	}
}

func (L *LeadsModule) GetMyRecords(o GetRecordsOptions) (Leads, error) {
	v, err := L.api.getMyRecords(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) GetRecords(o GetRecordsOptions) (Leads, error) {
	v, err := L.api.getRecords(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) GetRecordsByID(o GetRecordsByIdOptions) (Leads, error) {
	v, err := L.api.getRecordById(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Leads, error) {
	v, err := L.api.getDeletedRecordIds(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) InsertRecords(o InsertRecordsOptions) (Leads, error) {
	v, err := L.api.insertRecords(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) UpdateRecords(o UpdateRecordsOptions) (Leads, error) {
	v, err := L.api.updateRecords(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) DeleteRecords(id string) (Leads, error) {
	v, err := L.api.deleteRecords(L.module, id)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Leads, error) {
	v, err := L.api.getSearchRecordsByPDC(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Leads, error) {
	v, err := L.api.getRelatedRecords(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Leads, error) {
	v, err := L.api.updateRelatedRecord(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) GetFields(kind int) (Leads, error) {
	v, err := L.api.getFields(L.module, kind)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) UploadFile(o UploadFileOptions) (Leads, error) {
	v, err := L.api.uploadFile(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) DownloadFile(id string) (Leads, error) {
	v, err := L.api.downloadFile(L.module, id)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) DeleteFile(id string) (Leads, error) {
	v, err := L.api.deleteFile(L.module, id)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) Delink(o DelinkOptions) (Leads, error) {
	v, err := L.api.delink(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) UploadPhoto(o UploadPhotoOptions) (Leads, error) {
	v, err := L.api.uploadPhoto(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) DownloadPhoto(id string) (Leads, error) {
	v, err := L.api.downloadPhoto(L.module, id)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) DeletePhoto(id string) (Leads, error) {
	v, err := L.api.deletePhoto(L.module, id)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

func (L *LeadsModule) SearchRecords(o SearchRecordsOptions) (Leads, error) {
	v, err := L.api.searchRecords(L.module, o)
	if v == nil {
		return Leads{}, err
	}
	return v.(Leads), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Leads
type Lead struct {
	LeadOwner         string  `zoho:"Lead Owner"`          //
	Salutation        string  `zoho:"Salutation"`          //
	FirstName         string  `zoho:"First Name"`          // 40 chars
	Title             string  `zoho:"Title"`               // 100 chars
	LastName          string  `zoho:"Last Name"`           // 80 chars
	Company           string  `zoho:"Company"`             // 100 chars
	LeadSource        string  `zoho:"Lead Source"`         //
	Industry          string  `zoho:"Industry"`            //
	AnnualRevenue     float64 `zoho:"Annual Revenue"`      // 16 digits
	Phone             string  `zoho:"Phone"`               // 30 chars
	Mobile            string  `zoho:"Mobile"`              // 30 chars
	Fax               string  `zoho:"Fax"`                 // 30 chars
	Email             string  `zoho:"Email"`               // 100 chars
	SecondaryEmail    string  `zoho:"Secondary Email"`     // 100 chars
	SkypeID           string  `zoho:"Skype ID"`            // 50 chars
	Website           string  `zoho:"Website"`             // 120 chars
	LeadStatus        string  `zoho:"Lead Status"`         //
	Rating            string  `zoho:"Rating"`              //
	NumberOfEmployees int     `zoho:"Number Of Employees"` // 16 digits
	EmailOptOut       bool    `zoho:"Email Opt Out"`       //
	Street            string  `zoho:"Street"`              // 250 chars
	City              string  `zoho:"City"`                // 30 chars
	State             string  `zoho:"State"`               // 30 chars
	ZipCode           string  `zoho:"Zip Code"`            // 30 chars
	Country           string  `zoho:"Country"`             // 30 chars
	Description       string  `zoho:"Description"`         // 32000 chars

	//Fields not documented but returned by some methods
	//Some of these fields may need to be stripped due to zoho requirements
	LeadID           int       `zoho:"LEADID"`
	SmOwnerID        int       `zoho:"SMOWNERID"`
	SmCreatorID      int       `zoho:"SMCREATORID"`
	CreatedBy        string    `zoho:"Created By"`
	ModifiedByID     int       `zoho:"MODIFIEDBY"`
	ModifiedBy       string    `zoho:"Modified By"`
	CreatedTime      time.Time `zoho:"Created Time"`
	ModifiedTime     time.Time `zoho:"Modified Time"`
	FullName         string    `zoho:"Full Name"`
	Currency         string    `zoho:"Currency"`
	ExchangeRate     float64   `zoho:"Exchange Rate"`
	LastActivityTime time.Time `zoho:"Last Activity Time"`

	ExtraFields ExtraFields
}

func (l Lead) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Leads"}}
	x.addRow(l, 1)
	return x.encode()
}

func (l Lead) String() string {
	return l.writeXML()
}

type Leads []Lead

func (L Leads) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Leads"}}
	for i := 0; i < len(L); i++ {
		x.addRow(L[i], i+1)
	}
	return x.encode()
}

func (L Leads) String() string {
	return L.writeXML()
}
