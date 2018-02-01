package crm

import (
	"encoding/xml"
	"fmt"
)

// https://www.zoho.com/crm/help/api/error-messages.html
type CrmError struct {
	Type    string
	Code    int
	Message string
}

func (e CrmError) String() string {
	b, err := xml.Marshal(&XMLError{XMLName: xml.Name{Local: e.Type}, Code: e.Code, Message: e.Message})
	if err != nil {
		return ""
	}
	return string(b)
}

func (e CrmError) writeXML() string {
	return fmt.Sprint(e)
}

const (
	UseAuthTokenError              = 4000
	InternalServerError            = 4500
	APIKeyInactiveError            = 4501
	ModuleNotSupportedError        = 4502
	MandatoryFieldMissingError     = 4401
	IncorrectAPIParameterError     = 4600
	RateLimitExceededError         = 4820
	MissingParameterError          = 4831
	StringForIntegerError          = 4821
	InvalidTicketError             = 4824
	XMLParsingError                = 4835
	WrongAPIKeyError               = 4890
	ConvertLeadUnauthorizedError   = 4487
	AccessAPIUnauthorizedError     = 4001
	AccessModuleUnauthorizedError  = 401
	CreateRecordUnauthorizedError  = 401.1
	EditRecordUnauthorizedError    = 401.2
	DelteRecordUnauthorizedError   = 401.3
	ZohoCRMDisabledError           = 4101
	NoCRMAccountError              = 4102
	RecordIDUnavailableError       = 4103
	NoRecordsInModuleError         = 4422
	WrongSearchParameterValueError = 4420
	APICallsExceededError          = 4421
	RecordSearchLimitExceededError = 4423
	FileSizeLimitExceededError     = 4807
	InvalidFileTypeError           = 4424
	ExceededStorageSpaceError      = 4809
)
