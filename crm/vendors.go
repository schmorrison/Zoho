package crm

import (
	"encoding/xml"
	"time"
)

type VendorsModule CrmModule

func (a *API) Vendors() *VendorsModule {
	return &VendorsModule{
		id:     a.id,
		api:    a,
		module: vendorsModule,
	}
}

func (V *VendorsModule) GetMyRecords(o GetRecordsOptions) (Vendors, error) {
	v, err := V.api.getMyRecords(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) GetRecords(o GetRecordsOptions) (Vendors, error) {
	v, err := V.api.getRecords(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) GetRecordsByID(o GetRecordsByIdOptions) (Vendors, error) {
	v, err := V.api.getRecordById(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Vendors, error) {
	v, err := V.api.getDeletedRecordIds(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) InsertRecords(o InsertRecordsOptions) (Vendors, error) {
	v, err := V.api.insertRecords(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) UpdateRecords(o UpdateRecordsOptions) (Vendors, error) {
	v, err := V.api.updateRecords(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) DeleteRecords(id string) (Vendors, error) {
	v, err := V.api.deleteRecords(V.module, id)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Vendors, error) {
	v, err := V.api.getSearchRecordsByPDC(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Vendors, error) {
	v, err := V.api.getRelatedRecords(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Vendors, error) {
	v, err := V.api.updateRelatedRecord(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) GetFields(kind int) (Vendors, error) {
	v, err := V.api.getFields(V.module, kind)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) UploadFile(o UploadFileOptions) (Vendors, error) {
	v, err := V.api.uploadFile(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) DownloadFile(id string) (Vendors, error) {
	v, err := V.api.downloadFile(V.module, id)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) DeleteFile(id string) (Vendors, error) {
	v, err := V.api.deleteFile(V.module, id)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) Delink(o DelinkOptions) (Vendors, error) {
	v, err := V.api.delink(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) UploadPhoto(o UploadPhotoOptions) (Vendors, error) {
	v, err := V.api.uploadPhoto(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) DownloadPhoto(id string) (Vendors, error) {
	v, err := V.api.downloadPhoto(V.module, id)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) DeletePhoto(id string) (Vendors, error) {
	v, err := V.api.deletePhoto(V.module, id)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

func (V *VendorsModule) SearchRecords(o SearchRecordsOptions) (Vendors, error) {
	v, err := V.api.searchRecords(V.module, o)
	if v == nil {
		return Vendors{}, err
	}
	return v.(Vendors), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Vendors
type Vendor struct {
	VendorOwner  string  `zoho:"Vendor Owner"`  //
	VendorName   string  `zoho:"Vendor Name"`   // 50 chars
	Phone        string  `zoho:"Phone"`         // 40 chars
	Email        string  `zoho:"Email"`         // 100 chars
	Website      string  `zoho:"Website"`       // 50 chars
	GLAccount    string  `zoho:"GL Account"`    //
	Category     string  `zoho:"Category"`      // 40 chars
	Street       string  `zoho:"Street"`        // 250 chars
	City         string  `zoho:"City"`          // 30 chars
	State        string  `zoho:"State"`         // 30 chars
	ZipCode      string  `zoho:"Zip Code"`      // 30 chars
	Country      string  `zoho:"Country"`       // 30 chars
	Currency     string  `zoho:"Currency"`      //
	ExchangeRate float64 `zoho:"Exchange Rate"` //
	Description  string  `zoho:"Description"`   // 32000 chars

	//Fields not documented but returned by some methods
	//Some of these fields may need to be stripped due to zoho requirements
	VendorID     int       `zoho:"VENDORID"`
	SmOwnerID    int       `zoho:"SMOWNERID"`
	SmCreatorID  int       `zoho:"SMCREATORID"`
	CreatedBy    string    `zoho:"Created By"`
	ModifiedByID int       `zoho:"MODIFIEDBY"`
	ModifiedBy   string    `zoho:"Modified By"`
	CreatedTime  time.Time `zoho:"Created Time"`
	ModifiedTime time.Time `zoho:"Modified Time"`

	ExtraFields ExtraFields
}

func (v Vendor) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Vendors"}}
	x.addRow(v, 1)
	return x.encode()
}

func (v Vendor) String() string {
	return v.writeXML()
}

type Vendors []Vendor

func (V Vendors) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Vendors"}}
	for i := 0; i < len(V); i++ {
		x.addRow(V[i], i+1)
	}
	return x.encode()
}

func (V Vendors) String() string {
	return V.writeXML()
}
