package crm

import (
	"encoding/xml"
	"time"
)

type PotentialsModule CrmModule

func (a *API) Potentials() *PotentialsModule {
	return &PotentialsModule{
		id:     a.id,
		api:    a,
		module: potentialsModule,
	}
}

func (P *PotentialsModule) GetMyRecords(o GetRecordsOptions) (Potentials, error) {
	v, err := P.api.getMyRecords(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) GetRecords(o GetRecordsOptions) (Potentials, error) {
	v, err := P.api.getRecords(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) GetRecordsByID(o GetRecordsByIdOptions) (Potentials, error) {
	v, err := P.api.getRecordById(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Potentials, error) {
	v, err := P.api.getDeletedRecordIds(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) InsertRecords(o InsertRecordsOptions) (Potentials, error) {
	v, err := P.api.insertRecords(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) UpdateRecords(o UpdateRecordsOptions) (Potentials, error) {
	v, err := P.api.updateRecords(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) DeleteRecords(id string) (Potentials, error) {
	v, err := P.api.deleteRecords(P.module, id)
	return v.(Potentials), err
}

func (P *PotentialsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Potentials, error) {
	v, err := P.api.getSearchRecordsByPDC(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Potentials, error) {
	v, err := P.api.getRelatedRecords(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Potentials, error) {
	v, err := P.api.updateRelatedRecord(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) GetFields(kind int) (Potentials, error) {
	v, err := P.api.getFields(P.module, kind)
	return v.(Potentials), err
}

func (P *PotentialsModule) UploadFile(o UploadFileOptions) (Potentials, error) {
	v, err := P.api.uploadFile(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) DownloadFile(id string) (Potentials, error) {
	v, err := P.api.downloadFile(P.module, id)
	return v.(Potentials), err
}

func (P *PotentialsModule) DeleteFile(id string) (Potentials, error) {
	v, err := P.api.deleteFile(P.module, id)
	return v.(Potentials), err
}

func (P *PotentialsModule) Delink(o DelinkOptions) (Potentials, error) {
	v, err := P.api.delink(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) UploadPhoto(o UploadPhotoOptions) (Potentials, error) {
	v, err := P.api.uploadPhoto(P.module, o)
	return v.(Potentials), err
}

func (P *PotentialsModule) DownloadPhoto(id string) (Potentials, error) {
	v, err := P.api.downloadPhoto(P.module, id)
	return v.(Potentials), err
}

func (P *PotentialsModule) DeletePhoto(id string) (Potentials, error) {
	v, err := P.api.deletePhoto(P.module, id)
	return v.(Potentials), err
}

func (P *PotentialsModule) SearchRecords(o SearchRecordsOptions) (Potentials, error) {
	v, err := P.api.searchRecords(P.module, o)
	return v.(Potentials), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Potentials
type Potential struct {
	PotentialOwner  string    `zoho:"Potential Owner"`  //
	PotentialName   string    `zoho:"Potential Name"`   // chars 100
	AccountName     string    `zoho:"Account Name"`     //
	Type            string    `zoho:"Type"`             //
	LeadSource      string    `zoho:"Lead Source"`      //
	CampaignSource  string    `zoho:"Campaign Source"`  //
	ContactName     string    `zoho:"Contact Name"`     //
	Amount          float64   `zoho:"Amount"`           //
	ClosingDate     time.Time `zoho:"Closing Date"`     //
	NextStep        string    `zoho:"Next Step"`        // chars 100
	Stage           string    `zoho:"Stage"`            //
	Probability     int       `zoho:"Probability"`      //
	ExpectedRevenue float64   `zoho:"Expected Revenue"` //
	Description     string    `zoho:"Description"`      // chars 32000

	PotentialID      int       `zoho:"POTENTIALID"`
	SmOwnerID        int       `zoho:"SMOWNERID"`
	SmCreatorID      int       `zoho:"SMCREATORID"`
	CreatedBy        string    `zoho:"Created By"`
	ModifiedByID     int       `zoho:"MODIFIEDBY"`
	ModifiedBy       string    `zoho:"Modified By"`
	CreatedTime      time.Time `zoho:"Created Time"`
	ModifiedTime     time.Time `zoho:"Modified Time"`
	ContactID        int       `zoho:"CONTACTID"`
	Currency         string    `zoho:"Currency"`
	ExchangeRate     float64   `zoho:"Exchange Rate"`
	LastActivityTime time.Time `zoho:"Last Activity Time"`

	ExtraFields ExtraFields
}

func (p Potential) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Potentials"}}
	x.addRow(p, 1)
	return x.encode()
}

func (p Potential) String() string {
	return p.writeXML()
}

type Potentials []Potential

func (P Potentials) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Potentials"}}
	for i := 0; i < len(P); i++ {
		x.addRow(P[i], i+1)
	}
	return x.encode()
}

func (P Potentials) String() string {
	return P.writeXML()
}
