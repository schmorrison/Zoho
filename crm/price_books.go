package crm

import (
	"encoding/xml"
	"time"
)

type PriceBooksModule CrmModule

func (a *API) PriceBooks() *PriceBooksModule {
	return &PriceBooksModule{
		id:     a.id,
		api:    a,
		module: priceBooksModule,
	}
}

func (P *PriceBooksModule) GetMyRecords(o GetRecordsOptions) (PriceBooks, error) {
	v, err := P.api.getMyRecords(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) GetRecords(o GetRecordsOptions) (PriceBooks, error) {
	v, err := P.api.getRecords(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) GetRecordsByID(o GetRecordsByIdOptions) (PriceBooks, error) {
	v, err := P.api.getRecordById(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (PriceBooks, error) {
	v, err := P.api.getDeletedRecordIds(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) InsertRecords(o InsertRecordsOptions) (PriceBooks, error) {
	v, err := P.api.insertRecords(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) UpdateRecords(o UpdateRecordsOptions) (PriceBooks, error) {
	v, err := P.api.updateRecords(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) DeleteRecords(id string) (PriceBooks, error) {
	v, err := P.api.deleteRecords(P.module, id)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (PriceBooks, error) {
	v, err := P.api.getSearchRecordsByPDC(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) GetRelatedRecords(o GetRelatedRecordsOptions) (PriceBooks, error) {
	v, err := P.api.getRelatedRecords(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (PriceBooks, error) {
	v, err := P.api.updateRelatedRecord(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) GetFields(kind int) (PriceBooks, error) {
	v, err := P.api.getFields(P.module, kind)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) UploadFile(o UploadFileOptions) (PriceBooks, error) {
	v, err := P.api.uploadFile(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) DownloadFile(id string) (PriceBooks, error) {
	v, err := P.api.downloadFile(P.module, id)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) DeleteFile(id string) (PriceBooks, error) {
	v, err := P.api.deleteFile(P.module, id)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) Delink(o DelinkOptions) (PriceBooks, error) {
	v, err := P.api.delink(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) UploadPhoto(o UploadPhotoOptions) (PriceBooks, error) {
	v, err := P.api.uploadPhoto(P.module, o)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) DownloadPhoto(id string) (PriceBooks, error) {
	v, err := P.api.downloadPhoto(P.module, id)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) DeletePhoto(id string) (PriceBooks, error) {
	v, err := P.api.deletePhoto(P.module, id)
	return v.(PriceBooks), err
}

func (P *PriceBooksModule) SearchRecords(o SearchRecordsOptions) (PriceBooks, error) {
	v, err := P.api.searchRecords(P.module, o)
	return v.(PriceBooks), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#PriceBooks
type PriceBook struct {
	PriceBookOwner string `zoho:"Price Book Owner"` //
	PriceBookName  string `zoho:"Price Book Name"`  // 50 chars
	Active         bool   `zoho:"Active"`           //
	PricingModel   string `zoho:"Pricing Model"`    //
	Description    string `zoho:"Description"`      // 32000 chars
	PricingDetails []struct {
		ModelID   string `zoho:"Model Id"`   //
		FromRange string `zoho:"From Range"` //
		ToRange   string `zoho:"To Range"`   //
		Discount  string `zoho:"Discount"`   //
	} `zoho:"Pricing Details>discount"` //

	//Fields not documented but returned by some methods
	//Some of these fields may need to be stripped due to zoho requirements
	BookID       int       `zoho:"BOOKID"`
	SmOwnerID    int       `zoho:"SMOWNERID"`
	SmCreatorID  int       `zoho:"SMCREATORID"`
	CreatedBy    string    `zoho:"Created By"`
	ModifiedByID int       `zoho:"MODIFIEDBY"`
	ModifiedBy   string    `zoho:"Modified By"`
	CreatedTime  time.Time `zoho:"Created Time"`
	ModifiedTime time.Time `zoho:"Modified Time"`

	ExtraFields ExtraFields
}

func (p PriceBook) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "PriceBooks"}}
	x.addRow(p, 1)
	return x.encode()
}

func (p PriceBook) String() string {
	return p.writeXML()
}

type PriceBooks []PriceBook

func (P PriceBooks) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "PriceBooks"}}
	for i := 0; i < len(P); i++ {
		x.addRow(P[i], i+1)
	}
	return x.encode()
}

func (P PriceBooks) String() string {
	return P.writeXML()
}
