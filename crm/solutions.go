package crm

import "encoding/xml"

type SolutionsModule CrmModule

func (a *API) Solutions() *SolutionsModule {
	return &SolutionsModule{
		id:     a.id,
		api:    a,
		module: solutionsModule,
	}
}

func (S *SolutionsModule) GetMyRecords(o GetRecordsOptions) (Solutions, error) {
	v, err := S.api.getMyRecords(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) GetRecords(o GetRecordsOptions) (Solutions, error) {
	v, err := S.api.getRecords(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) GetRecordsByID(o GetRecordsByIdOptions) (Solutions, error) {
	v, err := S.api.getRecordById(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Solutions, error) {
	v, err := S.api.getDeletedRecordIds(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) InsertRecords(o InsertRecordsOptions) (Solutions, error) {
	v, err := S.api.insertRecords(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) UpdateRecords(o UpdateRecordsOptions) (Solutions, error) {
	v, err := S.api.updateRecords(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) DeleteRecords(id string) (Solutions, error) {
	v, err := S.api.deleteRecords(S.module, id)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Solutions, error) {
	v, err := S.api.getSearchRecordsByPDC(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Solutions, error) {
	v, err := S.api.getRelatedRecords(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Solutions, error) {
	v, err := S.api.updateRelatedRecord(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) GetFields(kind int) (Solutions, error) {
	v, err := S.api.getFields(S.module, kind)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) UploadFile(o UploadFileOptions) (Solutions, error) {
	v, err := S.api.uploadFile(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) DownloadFile(id string) (Solutions, error) {
	v, err := S.api.downloadFile(S.module, id)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) DeleteFile(id string) (Solutions, error) {
	v, err := S.api.deleteFile(S.module, id)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) Delink(o DelinkOptions) (Solutions, error) {
	v, err := S.api.delink(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) UploadPhoto(o UploadPhotoOptions) (Solutions, error) {
	v, err := S.api.uploadPhoto(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) DownloadPhoto(id string) (Solutions, error) {
	v, err := S.api.downloadPhoto(S.module, id)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) DeletePhoto(id string) (Solutions, error) {
	v, err := S.api.deletePhoto(S.module, id)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

func (S *SolutionsModule) SearchRecords(o SearchRecordsOptions) (Solutions, error) {
	v, err := S.api.searchRecords(S.module, o)
	if v == nil {
		return Solutions{}, err
	}
	return v.(Solutions), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Solutions
type Solution struct {
	SolutionNumber int    `zoho:"Solution Number"` // 16 digits
	SolutionName   string `zoho:"Solution Name"`   // 255 chars
	SolutionOwner  string `zoho:"Solution Owner"`  //
	ProductName    string `zoho:"Product Name"`    //
	Question       string `zoho:"Question"`        // 255 chars
	Answer         string `zoho:"Answer"`          //
	Status         string `zoho:"Status"`          //
	Description    string `zoho:"Description"`     // 32000 chars
	Comments       string `zoho:"Comments"`        // 32000 chars

	ExtraFields ExtraFields
}

func (s Solution) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Solutions"}}
	x.addRow(s, 1)
	return x.encode()
}

func (s Solution) String() string {
	return s.writeXML()
}

type Solutions []Solution

func (S Solutions) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Solutions"}}
	for i := 0; i < len(S); i++ {
		x.addRow(S[i], i+1)
	}
	return x.encode()
}

func (S Solutions) String() string {
	return S.writeXML()
}
