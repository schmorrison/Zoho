package crm

import (
	"encoding/xml"
	"time"
)

type InvoicesModule CrmModule

func (a *API) Invoices() *InvoicesModule {
	return &InvoicesModule{
		id:     a.id,
		api:    a,
		module: invoicesModule,
	}
}

func (I *InvoicesModule) GetMyRecords(o GetRecordsOptions) (Invoices, error) {
	v, err := I.api.getMyRecords(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) GetRecords(o GetRecordsOptions) (Invoices, error) {
	v, err := I.api.getRecords(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) GetRecordsByID(o GetRecordsByIdOptions) (Invoices, error) {
	v, err := I.api.getRecordById(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Invoices, error) {
	v, err := I.api.getDeletedRecordIds(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) InsertRecords(o InsertRecordsOptions) (Invoices, error) {
	v, err := I.api.insertRecords(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) UpdateRecords(o UpdateRecordsOptions) (Invoices, error) {
	v, err := I.api.updateRecords(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) DeleteRecords(id string) (Invoices, error) {
	v, err := I.api.deleteRecords(I.module, id)
	return v.(Invoices), err
}

func (I *InvoicesModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Invoices, error) {
	v, err := I.api.getSearchRecordsByPDC(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Invoices, error) {
	v, err := I.api.getRelatedRecords(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Invoices, error) {
	v, err := I.api.updateRelatedRecord(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) GetFields(kind int) (Invoices, error) {
	v, err := I.api.getFields(I.module, kind)
	return v.(Invoices), err
}

func (I *InvoicesModule) UploadFile(o UploadFileOptions) (Invoices, error) {
	v, err := I.api.uploadFile(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) DownloadFile(id string) (Invoices, error) {
	v, err := I.api.downloadFile(I.module, id)
	return v.(Invoices), err
}

func (I *InvoicesModule) DeleteFile(id string) (Invoices, error) {
	v, err := I.api.deleteFile(I.module, id)
	return v.(Invoices), err
}

func (I *InvoicesModule) Delink(o DelinkOptions) (Invoices, error) {
	v, err := I.api.delink(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) UploadPhoto(o UploadPhotoOptions) (Invoices, error) {
	v, err := I.api.uploadPhoto(I.module, o)
	return v.(Invoices), err
}

func (I *InvoicesModule) DownloadPhoto(id string) (Invoices, error) {
	v, err := I.api.downloadPhoto(I.module, id)
	return v.(Invoices), err
}

func (I *InvoicesModule) DeletePhoto(id string) (Invoices, error) {
	v, err := I.api.deletePhoto(I.module, id)
	return v.(Invoices), err
}

func (I *InvoicesModule) SearchRecords(o SearchRecordsOptions) (Invoices, error) {
	v, err := I.api.searchRecords(I.module, o)
	return v.(Invoices), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Invoices
type Invoice struct {
	InvoiceOwner    string    `zoho:"Invoice Owner"`    //
	InvoiceNumber   int       `zoho:"Invoice Number"`   //
	Subject         string    `zoho:"Subject"`          // 50 chars
	SalesOrder      string    `zoho:"Sales Order"`      //
	PurchaseOrder   string    `zoho:"Purchase Order"`   //
	CustomerNumber  string    `zoho:"Customer Number"`  // 50 chars
	ExciseDuty      float64   `zoho:"Excise Duty"`      //
	InvoiceDate     time.Time `zoho:"Invoice Date"`     //
	DueDate         time.Time `zoho:"Due Date"`         //
	SalesCommission float64   `zoho:"Sales Commission"` //
	AccountName     string    `zoho:"Account Name"`     //
	ContactName     string    `zoho:"Contact Name"`     //
	Status          bool      `zoho:"Status"`           //
	AssignedTo      string    `zoho:"Assigned To"`      //
	BillingStreet   string    `zoho:"Billing Street"`   // 250 chars
	BillingCity     string    `zoho:"Billing City"`     // 30 chars
	BillingState    string    `zoho:"Billing State"`    // 30 chars
	BillingZip      string    `zoho:"Billing Zip"`      // 30 chars
	BillingCountry  string    `zoho:"Billing Country"`  // 30 chars
	ShippingStreet  string    `zoho:"Shipping Street"`  // 250 chars
	ShippingCity    string    `zoho:"Shipping City"`    // 30 chars
	ShippingState   string    `zoho:"Shipping State"`   // 30 chars
	ShippingZip     string    `zoho:"Shipping Zip"`     // 30 chars
	ShippingCountry string    `zoho:"Shipping Country"` // 30 chars
	ProductDetails  []struct {
		ProductID          int     `zoho:"Product ID"`           //
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
	Adjustment   float64 `zoho:"Adjustment"`    //
	GrandTotal   float64 `zoho:"Grand Total"`   //
	SubTotal     float64 `zoho:"Sub Total"`     //
	Discount     float64 `zoho:"Discount"`      //
	Currency     string  `zoho:"Currency"`      //
	ExchangeRate float64 `zoho:"Exchange Rate"` //

	//Fields not documented but returned by some methods
	//Some of these fields may need to be stripped due to zoho requirements
	InvoiceID    int       `zoho:"INVOICEID"`
	SalesOrderID int       `zoho:"SALESORDERID"`
	ContactID    int       `zoho:"CONTACTID"`
	SmOwnerID    int       `zoho:"SMOWNERID"`
	SmCreatorID  int       `zoho:"SMCREATORID"`
	CreatedBy    string    `zoho:"Created By"`
	ModifiedByID int       `zoho:"MODIFIEDBY"`
	ModifiedBy   string    `zoho:"Modified By"`
	CreatedTime  time.Time `zoho:"Created Time"`
	ModifiedTime time.Time `zoho:"Modified Time"`
	Tax          int       `zoho:"Tax"`

	ExtraFields ExtraFields
}

func (i Invoice) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Invoices"}}
	x.addRow(i, 1)
	return x.encode()
}

func (i Invoice) String() string {
	return i.writeXML()
}

type Invoices []Invoice

func (I Invoices) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Invoices"}}
	for i := 0; i < len(I)+1; i++ {
		x.addRow(I[i], i+1)
	}
	return x.encode()
}

func (I Invoices) String() string {
	return I.writeXML()
}
