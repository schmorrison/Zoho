package crm

import (
	"encoding/xml"
	"time"
)

type QuotesModule CrmModule

func (a *API) Quotes() *QuotesModule {
	return &QuotesModule{
		id:     a.id,
		api:    a,
		module: quotesModule,
	}
}

func (Q *QuotesModule) GetMyRecords(o GetRecordsOptions) (Quotes, error) {
	v, err := Q.api.getMyRecords(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) GetRecords(o GetRecordsOptions) (Quotes, error) {
	v, err := Q.api.getRecords(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) GetRecordsByID(o GetRecordsByIdOptions) (Quotes, error) {
	v, err := Q.api.getRecordById(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Quotes, error) {
	v, err := Q.api.getDeletedRecordIds(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) InsertRecords(o InsertRecordsOptions) (Quotes, error) {
	v, err := Q.api.insertRecords(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) UpdateRecords(o UpdateRecordsOptions) (Quotes, error) {
	v, err := Q.api.updateRecords(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) DeleteRecords(id string) (Quotes, error) {
	v, err := Q.api.deleteRecords(Q.module, id)
	return v.(Quotes), err
}

func (Q *QuotesModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Quotes, error) {
	v, err := Q.api.getSearchRecordsByPDC(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Quotes, error) {
	v, err := Q.api.getRelatedRecords(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Quotes, error) {
	v, err := Q.api.updateRelatedRecord(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) GetFields(kind int) (Quotes, error) {
	v, err := Q.api.getFields(Q.module, kind)
	return v.(Quotes), err
}

func (Q *QuotesModule) UploadFile(o UploadFileOptions) (Quotes, error) {
	v, err := Q.api.uploadFile(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) DownloadFile(id string) (Quotes, error) {
	v, err := Q.api.downloadFile(Q.module, id)
	return v.(Quotes), err
}

func (Q *QuotesModule) DeleteFile(id string) (Quotes, error) {
	v, err := Q.api.deleteFile(Q.module, id)
	return v.(Quotes), err
}

func (Q *QuotesModule) Delink(o DelinkOptions) (Quotes, error) {
	v, err := Q.api.delink(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) UploadPhoto(o UploadPhotoOptions) (Quotes, error) {
	v, err := Q.api.uploadPhoto(Q.module, o)
	return v.(Quotes), err
}

func (Q *QuotesModule) DownloadPhoto(id string) (Quotes, error) {
	v, err := Q.api.downloadPhoto(Q.module, id)
	return v.(Quotes), err
}

func (Q *QuotesModule) DeletePhoto(id string) (Quotes, error) {
	v, err := Q.api.deletePhoto(Q.module, id)
	return v.(Quotes), err
}

func (Q *QuotesModule) SearchRecords(o SearchRecordsOptions) (Quotes, error) {
	v, err := Q.api.searchRecords(Q.module, o)
	return v.(Quotes), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Quotes
type Quote struct {
	QuoteOwner       string    `zoho:"Quote Owner"`       //
	Subject          string    `zoho:"Subject"`           // 50 chars
	PotentialName    string    `zoho:"Potential Name"`    // 40 chars
	QuoteStage       bool      `zoho:"Quote Stage"`       //
	ValidTill        time.Time `zoho:"Valid Till"`        //
	ContactName      string    `zoho:"Contact Name"`      //
	AccountName      string    `zoho:"Account Name"`      //
	Carrier          string    `zoho:"Carrier"`           //
	Shipping         string    `zoho:"Shipping"`          // 50 chars
	InventoryManager string    `zoho:"Inventory Manager"` // 50 chars
	BillingStreet    string    `zoho:"Billing Street"`    // 250 chars
	BillingCity      string    `zoho:"Billing City"`      // 30 chars
	BillingState     string    `zoho:"Billing State"`     // 30 chars
	BillingZip       string    `zoho:"Billing Zip"`       // 30 chars
	BillingCountry   string    `zoho:"Billing Country"`   // 30 chars
	ShippingStreet   string    `zoho:"Shipping Street"`   // 250 chars
	ShippingCity     string    `zoho:"Shipping City"`     // 30 chars
	ShippingState    string    `zoho:"Shipping State"`    // 30 chars
	ShippingZip      string    `zoho:"Shipping Zip"`      // 30 chars
	ShippingCountry  string    `zoho:"Shipping Country"`  // 30 chars
	ProductDetails   []struct {
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
		ProductDescription string  `zoho:"Product Description"`  //
	} `zoho:"Product Details>product"` //
	Discount           int     `zoho:"Discount"`      //
	Currency           string  `zoho:"Currency"`      //
	ExchangeRate       float64 `zoho:"Exchange Rate"` //
	TermsAndConditions string  `zoho:"Terms"`         //
	Description        string  `zoho:"Description"`   // 32000 chars

	QuoteID      int       `zoho:"QUOTEID"`
	QuoteNumber  int       `zoho:"Quote Number"`
	ContactID    int       `zoho:"CONTACTID"`
	SmOwnerID    int       `zoho:"SMOWNERID"`
	SmCreatorID  int       `zoho:"SMCREATORID"`
	CreatedBy    string    `zoho:"Created By"`
	ModifiedByID int       `zoho:"MODIFIEDBY"`
	ModifiedBy   string    `zoho:"Modified By"`
	CreatedTime  time.Time `zoho:"Created Time"`
	ModifiedTime time.Time `zoho:"Modified Time"`
	SubTotal     float64   `zoho:"Sub Total"`
	Tax          float64   `zoho:"Tax"`
	Adjustment   float64   `zoho:"Adjustment"`
	GrandTotal   float64   `zoho:"Grand Total"`

	ExtraFields ExtraFields
}

func (q Quote) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Quotes"}}
	x.addRow(q, 1)
	return x.encode()
}

func (q Quote) String() string {
	return q.writeXML()
}

type Quotes []Quote

func (Q Quotes) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Quotes"}}
	for i := 0; i < len(Q); i++ {
		x.addRow(Q[i], i+1)
	}
	return x.encode()
}

func (Q Quotes) String() string {
	return Q.writeXML()
}
