package crm

import (
	"encoding/xml"
	"time"
)

type PurchaseOrdersModule CrmModule

func (a *API) PurchaseOrders() *PurchaseOrdersModule {
	return &PurchaseOrdersModule{
		id:     a.id,
		api:    a,
		module: purchaseOrdersModule,
	}
}

func (P *PurchaseOrdersModule) GetMyRecords(o GetRecordsOptions) (PurchaseOrders, error) {
	v, err := P.api.getMyRecords(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) GetRecords(o GetRecordsOptions) (PurchaseOrders, error) {
	v, err := P.api.getRecords(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) GetRecordsByID(o GetRecordsByIdOptions) (PurchaseOrders, error) {
	v, err := P.api.getRecordById(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (PurchaseOrders, error) {
	v, err := P.api.getDeletedRecordIds(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) InsertRecords(o InsertRecordsOptions) (PurchaseOrders, error) {
	v, err := P.api.insertRecords(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) UpdateRecords(o UpdateRecordsOptions) (PurchaseOrders, error) {
	v, err := P.api.updateRecords(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) DeleteRecords(id string) (PurchaseOrders, error) {
	v, err := P.api.deleteRecords(P.module, id)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (PurchaseOrders, error) {
	v, err := P.api.getSearchRecordsByPDC(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) GetRelatedRecords(o GetRelatedRecordsOptions) (PurchaseOrders, error) {
	v, err := P.api.getRelatedRecords(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (PurchaseOrders, error) {
	v, err := P.api.updateRelatedRecord(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) GetFields(kind int) (PurchaseOrders, error) {
	v, err := P.api.getFields(P.module, kind)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) UploadFile(o UploadFileOptions) (PurchaseOrders, error) {
	v, err := P.api.uploadFile(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) DownloadFile(id string) (PurchaseOrders, error) {
	v, err := P.api.downloadFile(P.module, id)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) DeleteFile(id string) (PurchaseOrders, error) {
	v, err := P.api.deleteFile(P.module, id)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) Delink(o DelinkOptions) (PurchaseOrders, error) {
	v, err := P.api.delink(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) UploadPhoto(o UploadPhotoOptions) (PurchaseOrders, error) {
	v, err := P.api.uploadPhoto(P.module, o)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) DownloadPhoto(id string) (PurchaseOrders, error) {
	v, err := P.api.downloadPhoto(P.module, id)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) DeletePhoto(id string) (PurchaseOrders, error) {
	v, err := P.api.deletePhoto(P.module, id)
	return v.(PurchaseOrders), err
}

func (P *PurchaseOrdersModule) SearchRecords(o SearchRecordsOptions) (PurchaseOrders, error) {
	v, err := P.api.searchRecords(P.module, o)
	return v.(PurchaseOrders), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#PurchaseOrders
type PurchaseOrder struct {
	PurchaseOrderOwner string    `zoho:"Purchase Order Owner"`  //
	PONumber           int       `zoho:"Purchase Order Number"` //
	Subject            string    `zoho:"Subject"`               // 50 chars
	VendorName         string    `zoho:"Vendor Name"`           // 20 chars
	RequisitionNumber  string    `zoho:"Requisition Number"`    // 20 chars
	TrackingNumber     string    `zoho:"Tracking Number"`       //
	ContactName        string    `zoho:"Contact Name"`          //
	Carrier            string    `zoho:"Carrier"`               //
	PurchaseOrderDate  time.Time `zoho:"Purchase Order Date"`   //
	DueDate            time.Time `zoho:"Due Date"`              //
	ExciseDuty         float64   `zoho:"Excise Duty"`           //
	SalesCommission    float64   `zoho:"Sales Commission"`      //
	Status             bool      `zoho:"Status"`                //
	AssignedTo         string    `zoho:"Assigned To"`           //
	BillingStreet      string    `zoho:"Billing Street"`        // 250 chars
	BillingCity        string    `zoho:"Billing City"`          // 30 chars
	BillingState       string    `zoho:"Billing State"`         // 30 chars
	BillingZip         string    `zoho:"Billing Zip"`           // 30 chars
	BillingCountry     string    `zoho:"Billing Country"`       // 30 chars
	ShippingStreet     string    `zoho:"Shipping Street"`       // 250 chars
	ShippingCity       string    `zoho:"Shipping City"`         // 30 chars
	ShippingState      string    `zoho:"Shipping State"`        // 30 chars
	ShippingZip        string    `zoho:"Shipping Zip"`          // 30 chars
	ShippingCountry    string    `zoho:"Shipping Country"`      // 30 chars
	ProductDetails     []struct {
		ProductID          int     `zoho:"Product ID"`           //
		ProductName        string  `zoho:"Product Name"`         //
		UnitPrice          float64 `zoho:"Unit Price"`           //
		Quantity           int     `zoho:"Quantity"`             //
		QuantityInStock    int     `zoho:"Quantity In Stock"`    //
		Total              float64 `zoho:"Total"`                //
		Discount           float64 `zoho:"Discount"`             //
		TotalAfterDiscount float64 `zoho:"Total After Discount"` //
		ListPrice          float64 `zoho:"List Price"`           //
		NetTotal           float64 `zoho:"Net Total"`            //
		Tax                float64 `zoho:"Tax"`                  //
		ProductDescription string  `zoho:"Product Description"`  //
	} `zoho:"Product Details>product"` //
	Discount     float64 `zoho:"Discount"`      //
	Currency     string  `zoho:"Currency"`      //
	ExchangeRate float64 `zoho:"Exchange Rate"` //
	Terms        string  `zoho:"Terms"`         //

	ExtraFields ExtraFields
}

func (p PurchaseOrder) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "PurchaseOrders"}}
	x.addRow(p, 1)
	return x.encode()
}

func (p PurchaseOrder) String() string {
	return p.writeXML()
}

type PurchaseOrders []PurchaseOrder

func (P PurchaseOrders) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "PurchaseOrders"}}
	for i := 0; i < len(P); i++ {
		x.addRow(P[i], i+1)
	}
	return x.encode()
}

func (P PurchaseOrders) String() string {
	return P.writeXML()
}
