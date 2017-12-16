package crm

import (
	"encoding/xml"
	"time"
)

type CampaignsModule CrmModule

func (a *API) Campaigns() *CampaignsModule {
	return &CampaignsModule{
		id:     a.id,
		api:    a,
		module: campaignsModule,
	}
}

func (C *CampaignsModule) GetMyRecords(o GetRecordsOptions) (Campaigns, error) {
	v, err := C.api.getMyRecords(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) GetRecords(o GetRecordsOptions) (Campaigns, error) {
	v, err := C.api.getRecords(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) GetRecordsByID(o GetRecordsByIdOptions) (Campaigns, error) {
	v, err := C.api.getRecordById(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Campaigns, error) {
	v, err := C.api.getDeletedRecordIds(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) InsertRecords(o InsertRecordsOptions) (Campaigns, error) {
	v, err := C.api.insertRecords(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) UpdateRecords(o UpdateRecordsOptions) (Campaigns, error) {
	v, err := C.api.updateRecords(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) DeleteRecords(id string) (Campaigns, error) {
	v, err := C.api.deleteRecords(C.module, id)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Campaigns, error) {
	v, err := C.api.getSearchRecordsByPDC(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Campaigns, error) {
	v, err := C.api.getRelatedRecords(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Campaigns, error) {
	v, err := C.api.updateRelatedRecord(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) GetFields(kind int) (Campaigns, error) {
	v, err := C.api.getFields(C.module, kind)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) UploadFile(o UploadFileOptions) (Campaigns, error) {
	v, err := C.api.uploadFile(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) DownloadFile(id string) (Campaigns, error) {
	v, err := C.api.downloadFile(C.module, id)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) DeleteFile(id string) (Campaigns, error) {
	v, err := C.api.deleteFile(C.module, id)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) Delink(o DelinkOptions) (Campaigns, error) {
	v, err := C.api.delink(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) UploadPhoto(o UploadPhotoOptions) (Campaigns, error) {
	v, err := C.api.uploadPhoto(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) DownloadPhoto(id string) (Campaigns, error) {
	v, err := C.api.downloadPhoto(C.module, id)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) DeletePhoto(id string) (Campaigns, error) {
	v, err := C.api.deletePhoto(C.module, id)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

func (C *CampaignsModule) SearchRecords(o SearchRecordsOptions) (Campaigns, error) {
	v, err := C.api.searchRecords(C.module, o)
	if v == nil {
		return Campaigns{}, err
	}
	return v.(Campaigns), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Campaigns
type Campaign struct {
	CampaignOwner    string    `zoho:"Campaign Owner"`    //
	CampaignName     string    `zoho:"Campaign Name"`     // 40 chars
	Type             string    `zoho:"Type"`              //
	Status           string    `zoho:"Status"`            //
	StartDate        time.Time `zoho:"Start Date"`        //
	EndDate          time.Time `zoho:"End Date"`          //
	ExpectedRevenue  float64   `zoho:"Expected Revenue"`  //
	ActualCost       float64   `zoho:"Actual Cost"`       //
	BudgetedCost     float64   `zoho:"Budgeted Cost"`     //
	ExpectedResponse int       `zoho:"Expected Response"` //
	NumberSent       int       `zoho:"Num sent"`          //
	Description      string    `zoho:"Description"`       // 32000 chars

	CampaignID   int       `zoho:"CAMPAIGNID"`    //
	SmOwnerID    int       `zoho:"SMOWNERID"`     //
	SmCreatorID  int       `zoho:"SMCREATORID"`   //
	CreatedBy    string    `zoho:"Created By"`    //
	ModifiedByID int       `zoho:"MODIFIEDBY"`    //
	ModifiedBy   string    `zoho:"Modified By"`   //
	CreatedTime  time.Time `zoho:"Created Time"`  //
	ModifiedTime time.Time `zoho:"Modified Time"` //
	Currency     string    `zoho:"Currency"`      //
	ExchangeRate float64   `zoho:"Exchange Rate"` //

	ExtraFields ExtraFields
}

func (c Campaign) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Campaigns"}}
	x.addRow(c, 1)
	return x.encode()
}

func (c Campaign) String() string {
	return c.writeXML()
}

type Campaigns []Campaign

func (C Campaigns) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Campaigns"}}
	for i := 0; i < len(C); i++ {
		x.addRow(C[i], i+1)
	}
	return x.encode()
}

func (C Campaigns) String() string {
	return C.writeXML()
}
