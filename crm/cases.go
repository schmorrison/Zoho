package crm

import "encoding/xml"

type CasesModule CrmModule

func (a *API) Cases() *CasesModule {
	return &CasesModule{
		id:     a.id,
		api:    a,
		module: casesModule,
	}
}

func (C *CasesModule) GetMyRecords(o GetRecordsOptions) (Cases, error) {
	v, err := C.api.getMyRecords(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) GetRecords(o GetRecordsOptions) (Cases, error) {
	v, err := C.api.getRecords(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) GetRecordsByID(o GetRecordsByIdOptions) (Cases, error) {
	v, err := C.api.getRecordById(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Cases, error) {
	v, err := C.api.getDeletedRecordIds(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) InsertRecords(o InsertRecordsOptions) (Cases, error) {
	v, err := C.api.insertRecords(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) UpdateRecords(o UpdateRecordsOptions) (Cases, error) {
	v, err := C.api.updateRecords(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) DeleteRecords(id string) (Cases, error) {
	v, err := C.api.deleteRecords(C.module, id)
	return v.(Cases), err
}

func (C *CasesModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Cases, error) {
	v, err := C.api.getSearchRecordsByPDC(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Cases, error) {
	v, err := C.api.getRelatedRecords(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Cases, error) {
	v, err := C.api.updateRelatedRecord(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) GetFields(kind int) (Cases, error) {
	v, err := C.api.getFields(C.module, kind)
	return v.(Cases), err
}

func (C *CasesModule) UploadFile(o UploadFileOptions) (Cases, error) {
	v, err := C.api.uploadFile(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) DownloadFile(id string) (Cases, error) {
	v, err := C.api.downloadFile(C.module, id)
	return v.(Cases), err
}

func (C *CasesModule) DeleteFile(id string) (Cases, error) {
	v, err := C.api.deleteFile(C.module, id)
	return v.(Cases), err
}

func (C *CasesModule) Delink(o DelinkOptions) (Cases, error) {
	v, err := C.api.delink(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) UploadPhoto(o UploadPhotoOptions) (Cases, error) {
	v, err := C.api.uploadPhoto(C.module, o)
	return v.(Cases), err
}

func (C *CasesModule) DownloadPhoto(id string) (Cases, error) {
	v, err := C.api.downloadPhoto(C.module, id)
	return v.(Cases), err
}

func (C *CasesModule) DeletePhoto(id string) (Cases, error) {
	v, err := C.api.deletePhoto(C.module, id)
	return v.(Cases), err
}

func (C *CasesModule) SearchRecords(o SearchRecordsOptions) (Cases, error) {
	v, err := C.api.searchRecords(C.module, o)
	return v.(Cases), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Cases
type Case struct {
	CaseOwner        string `zoho:"Case Owner"`        //
	Subject          string `zoho:"Subject"`           // 255 chars
	ProductName      string `zoho:"Product Name"`      //
	Priority         string `zoho:"Priority"`          //
	CaseNumber       int    `zoho:"Case Number"`       //
	Status           string `zoho:"Status"`            //
	ReportedBy       string `zoho:"Reported By"`       //
	RelatedTo        string `zoho:"Related To"`        //
	Type             string `zoho:"Type"`              //
	Case             string `zoho:"Case"`              //
	Email            string `zoho:"Email"`             //
	AccountName      string `zoho:"Account Name"`      //
	PotentialName    string `zoho:"Potential Name"`    //
	Phone            string `zoho:"Phone"`             // 50 chars
	CaseReason       string `zoho:"Case Reason"`       //
	Description      string `zoho:"Description"`       // 32000 chars
	InternalComments string `zoho:"Internal Comments"` // 3000 chars
	Solution         string `zoho:"Solution"`          // 32000 chars
	AddComment       string `zoho:"Add Comment"`       // 32000 chars

	ExtraFields ExtraFields
}

func (c Case) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Cases"}}
	x.addRow(c, 1)
	return x.encode()
}

func (c Case) String() string {
	return c.writeXML()
}

type Cases []Case

func (C Cases) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Cases"}}
	for i := 0; i < len(C); i++ {
		x.addRow(C[i], i+1)
	}
	return x.encode()
}

func (C Cases) String() string {
	return C.writeXML()
}
