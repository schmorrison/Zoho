package crm

import (
	"encoding/xml"
	"time"
)

type SalesOrdersModule CrmModule

func (a *API) SalesOrders() *SalesOrdersModule {
	return &SalesOrdersModule{
		id:     a.id,
		api:    a,
		module: salesOrdersModule,
	}
}

func (S *SalesOrdersModule) GetMyRecords(o GetRecordsOptions) (SalesOrders, error) {
	v, err := S.api.getMyRecords(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) GetRecords(o GetRecordsOptions) (SalesOrders, error) {
	v, err := S.api.getRecords(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) GetRecordsByID(o GetRecordsByIdOptions) (SalesOrders, error) {
	v, err := S.api.getRecordById(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (SalesOrders, error) {
	v, err := S.api.getDeletedRecordIds(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) InsertRecords(o InsertRecordsOptions) (SalesOrders, error) {
	v, err := S.api.insertRecords(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) UpdateRecords(o UpdateRecordsOptions) (SalesOrders, error) {
	v, err := S.api.updateRecords(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) DeleteRecords(id string) (SalesOrders, error) {
	v, err := S.api.deleteRecords(S.module, id)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (SalesOrders, error) {
	v, err := S.api.getSearchRecordsByPDC(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) GetRelatedRecords(o GetRelatedRecordsOptions) (SalesOrders, error) {
	v, err := S.api.getRelatedRecords(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (SalesOrders, error) {
	v, err := S.api.updateRelatedRecord(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) GetFields(kind int) (SalesOrders, error) {
	v, err := S.api.getFields(S.module, kind)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) UploadFile(o UploadFileOptions) (SalesOrders, error) {
	v, err := S.api.uploadFile(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) DownloadFile(id string) (SalesOrders, error) {
	v, err := S.api.downloadFile(S.module, id)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) DeleteFile(id string) (SalesOrders, error) {
	v, err := S.api.deleteFile(S.module, id)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) Delink(o DelinkOptions) (SalesOrders, error) {
	v, err := S.api.delink(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) UploadPhoto(o UploadPhotoOptions) (SalesOrders, error) {
	v, err := S.api.uploadPhoto(S.module, o)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) DownloadPhoto(id string) (SalesOrders, error) {
	v, err := S.api.downloadPhoto(S.module, id)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) DeletePhoto(id string) (SalesOrders, error) {
	v, err := S.api.deletePhoto(S.module, id)
	return v.(SalesOrders), err
}

func (S *SalesOrdersModule) SearchRecords(o SearchRecordsOptions) (SalesOrders, error) {
	v, err := S.api.searchRecords(S.module, o)
	return v.(SalesOrders), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#SalesOrders
type SalesOrder struct {
	SalesOrderOwner string    `zoho:"Sales Order Owner"` //
	SONumber        int       `zoho:"SO Number"`         //
	Subject         string    `zoho:"Subject"`           // 50 chars
	PotentialName   string    `zoho:"Potential Name"`    //
	CustomerNumber  string    `zoho:"Customer Number"`   // 50 chars
	PurchaseOrder   string    `zoho:"Purchase Order"`    //
	QuoteName       string    `zoho:"Quote Name"`        //
	ContactName     string    `zoho:"Contact Name"`      //
	DueDate         time.Time `zoho:"Due Date"`          //
	Carrier         string    `zoho:"Carrier"`           //
	Pending         string    `zoho:"Pending"`           // 50 chars
	Status          string    `zoho:"Status"`            //
	SalesCommission float64   `zoho:"Sales Commission"`  //
	ExciseDuty      float64   `zoho:"Excise Duty"`       //
	AccountName     string    `zoho:"Account Name"`      //
	AssignedTo      string    `zoho:"Assigned To"`       //
	BillingStreet   string    `zoho:"Billing Street"`    // 250 chars
	BillingCity     string    `zoho:"Billing City"`      // 30 chars
	BillingState    string    `zoho:"Billing State"`     // 30 chars
	BillingZip      string    `zoho:"Billing Zip"`       // 30 chars
	BillingCountry  string    `zoho:"Billing Country"`   // 30 chars
	ShippingStreet  string    `zoho:"Shipping Street"`   // 250 chars
	ShippingCity    string    `zoho:"Shipping City"`     // 30 chars
	ShippingState   string    `zoho:"Shipping State"`    // 30 chars
	ShippingZip     string    `zoho:"Shipping Zip"`      // 30 chars
	ShippingCountry string    `zoho:"Shipping Country"`  // 30 chars
	ProductDetails  []struct {
		ProductID          int     `zoho:"Product Id"`           //
		ProductName        string  `zoho:"Product Name"`         //
		UnitPrice          float64 `zoho:"Unit Price"`           //
		Quantity           int     `zoho:"Quantity"`             //
		QuantityInStock    int     `zoho:"Quantity in Stock"`    //
		Total              float64 `zoho:"Total"`                //
		Discount           float64 `zoho:"Discount"`             //
		TotalAfterDiscount float64 `zoho:"Total After Discount"` //
		ListPrice          float64 `zoho:"List Price"`           //
		NetTotal           float64 `zoho:"Net Total"`            //
		Tax                float64 `zoho:"Tax"`                  //
		LineTax            string  `zoho:"Line Tax"`             //
		ProductDescription string  `zoho:"Product Description"`  //
	} `zoho:"Product Details>product"` //
	Discount     float64 `zoho:"Discount"`      //
	Currency     string  `zoho:"Currency"`      //
	ExchangeRate float64 `zoho:"Exchange Rate"` //
	Terms        string  `zoho:"Terms"`         //

	//Fields not documented but returned by some methods
	//Some of these fields may need to be stripped on 'insert' due to zoho requirements
	SalesOrderID int       `zoho:"SALESORDERID"`  //
	ContactID    int       `zoho:"CONTACTID"`     //
	SmOwnerID    int       `zoho:"SMOWNERID"`     //
	SmCreatorID  int       `zoho:"SMCREATORID"`   //
	CreatedBy    string    `zoho:"Created By"`    //
	ModifiedByID int       `zoho:"MODIFIEDBY"`    //
	ModifiedBy   string    `zoho:"Modified By"`   //
	CreatedTime  time.Time `zoho:"Created Time"`  //
	ModifiedTime time.Time `zoho:"Modified Time"` //
	SubTotal     float64   `zoho:"Sub Total"`     //
	Tax          float64   `zoho:"Tax"`
	BillingCode  string    `zoho:"Billing Code"`
	Adjustment   float64   `zoho:"Adjustment"`  //
	GrandTotal   float64   `zoho:"Grand Total"` //

	ExtraFields ExtraFields
}

func (s SalesOrder) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "SalesOrders"}}
	x.addRow(s, 1)
	return x.encode()
}

func (s SalesOrder) String() string {
	return s.writeXML()
}

type SalesOrders []SalesOrder

func (S SalesOrders) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "SalesOrders"}}
	for i := 0; i < len(S); i++ {
		x.addRow(S[i], i+1)
	}
	return x.encode()
}

func (S SalesOrders) String() string {
	return S.writeXML()
}
