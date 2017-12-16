package crm

import (
	"encoding/xml"
	"time"
)

type ProductsModule CrmModule

func (a *API) Products() *ProductsModule {
	return &ProductsModule{
		id:     a.id,
		api:    a,
		module: productsModule,
	}
}

func (P *ProductsModule) GetMyRecords(o GetRecordsOptions) (Products, error) {
	v, err := P.api.getMyRecords(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) GetRecords(o GetRecordsOptions) (Products, error) {
	v, err := P.api.getRecords(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) GetRecordsByID(o GetRecordsByIdOptions) (Products, error) {
	v, err := P.api.getRecordById(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Products, error) {
	v, err := P.api.getDeletedRecordIds(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) InsertRecords(o InsertRecordsOptions) (Products, error) {
	v, err := P.api.insertRecords(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) UpdateRecords(o UpdateRecordsOptions) (Products, error) {
	v, err := P.api.updateRecords(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) DeleteRecords(id string) (Products, error) {
	v, err := P.api.deleteRecords(P.module, id)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Products, error) {
	v, err := P.api.getSearchRecordsByPDC(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Products, error) {
	v, err := P.api.getRelatedRecords(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Products, error) {
	v, err := P.api.updateRelatedRecord(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) GetFields(kind int) (Products, error) {
	v, err := P.api.getFields(P.module, kind)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) UploadFile(o UploadFileOptions) (Products, error) {
	v, err := P.api.uploadFile(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) DownloadFile(id string) (Products, error) {
	v, err := P.api.downloadFile(P.module, id)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) DeleteFile(id string) (Products, error) {
	v, err := P.api.deleteFile(P.module, id)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) Delink(o DelinkOptions) (Products, error) {
	v, err := P.api.delink(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) UploadPhoto(o UploadPhotoOptions) (Products, error) {
	v, err := P.api.uploadPhoto(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) DownloadPhoto(id string) (Products, error) {
	v, err := P.api.downloadPhoto(P.module, id)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) DeletePhoto(id string) (Products, error) {
	v, err := P.api.deletePhoto(P.module, id)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

func (P *ProductsModule) SearchRecords(o SearchRecordsOptions) (Products, error) {
	v, err := P.api.searchRecords(P.module, o)
	if v == nil {
		return Products{}, err
	}
	return v.(Products), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Products
type Product struct {
	ProductName       string    `zoho:"Product Name"`        // 50 chars
	ProductOwner      string    `zoho:"Product Owner"`       //
	ProductCode       string    `zoho:"Product Code"`        // 40 chars
	ProductActive     bool      `zoho:"Product Active"`      //
	VendorName        string    `zoho:"Vendor Name"`         //
	ProductCategory   string    `zoho:"Product Category"`    //
	SalesStartDate    time.Time `zoho:"Sales Start Date"`    //
	SalesEndDate      time.Time `zoho:"Sales End Date"`      //
	CommissionRate    float64   `zoho:"Commission Rate"`     //
	Manufacturer      string    `zoho:"Manufacturer"`        //
	UnitPrice         float64   `zoho:"Unit Price"`          //
	Taxable           bool      `zoho:"Taxable"`             //
	SupportStartDate  time.Time `zoho:"Support Start Date"`  //
	SupportExpiryDate time.Time `zoho:"Support Expiry Date"` //
	UsageUnit         string    `zoho:"Usage Unit"`          //
	QuantityOrdered   int       `zoho:"Qty Ordered"`         //
	QuantityStocked   int       `zoho:"Qty in Stock"`        //
	ReorderLevel      int       `zoho:"Reorder Level"`       //
	Handler           string    `zoho:"Handler"`             //
	QuantityDemand    int       `zoho:"Qty in Demand"`       //
	Description       string    `zoho:"Description"`         // 32000 chars

	ProductID    int       `zoho:"PRODUCTID"`
	SmOwnerID    int       `zoho:"SMOWNERID"`
	SmCreatorID  int       `zoho:"SMCREATORID"`
	CreatedBy    string    `zoho:"Created By"`
	ModifiedByID int       `zoho:"MODIFIEDBY"`
	ModifiedBy   string    `zoho:"Modified By"`
	CreatedTime  time.Time `zoho:"Created Time"`
	ModifiedTime time.Time `zoho:"Modified Time"`

	ExtraFields ExtraFields
}

func (p Product) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Products"}}
	x.addRow(p, 1)
	return x.encode()
}

func (p Product) String() string {
	return p.writeXML()
}

type Products []Product

func (P Products) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Products"}}
	for i := 0; i < len(P); i++ {
		x.addRow(P[i], i+1)
	}
	return x.encode()
}

func (P Products) String() string {
	return P.writeXML()
}
