package crm

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
)

func getValuesFromXML(b []byte) (XMLData, error) {
	data := XMLData{}
	values := []Row{}
	currentRow := []FieldLabel{}
	nestedField := FieldLabel{}
	nestedGroup := InternalGroup{}
	module := ""
	nested := ""
	//Make an XML decoder from the response body
	decoder := xml.NewDecoder(strings.NewReader(string(b)))
PRIME:
	for {
		// iterate over XML documents tokens
		t1, err := decoder.Token()
		if err != nil && err != io.EOF {
			fmt.Println("Got error on primer 'tokener'")
			log.Fatal(err)
		}
		if t1 == nil {
			break PRIME
		}
		// Inspect the type of the token just read.
		switch e1 := t1.(type) {
		case xml.StartElement:
			switch e1.Name.Local {
			case "response", "result":
			case "nodata":
				//found 'nodata' error so decode element into 'XMLError'
				e := XMLError{XMLName: xml.Name{Local: "nodata"}}
				err = decoder.DecodeElement(&e, &e1)
				if err != nil {
					fmt.Println("Got 'error' but couldn't decode")
					log.Fatal(err)
				}
				data.Error = e
				return data, nil
			case "error":
				//found error so decode element into 'XMLError'
				e := XMLError{XMLName: xml.Name{Local: "error"}}
				err = decoder.DecodeElement(&e, &e1)
				if err != nil {
					fmt.Println("Got 'error' but couldn't decode")
					log.Fatal(err)
				}
				data.Error = e
				return data, nil
			case "row":
				//start a []FieldLabel in 'currentRow'
				currentRow = []FieldLabel{}
			case "FL":
				//check for a nested item
				if checkForInternalGroup(xml.NewDecoder(strings.NewReader(string(b))), e1) {
					//if another start element with name
				ATTR:
					for _, a := range e1.Attr {
						//save e1.Attr("val") as field label in nestField
						if a.Name.Local == "val" {
							nestedField = FieldLabel{Label: a.Value}
							break ATTR
						}
					}
				} else {
					//decode FL and append to 'currentRow'
					fl := FieldLabel{}
					fl.decode(decoder, t1)
					if nested != "" {
						//append to 'internal group'
						nestedGroup.Fields = append(nestedGroup.Fields, fl)
					} else {
						currentRow = append(currentRow, fl)
					}
				}
			default:
				if module == "" {
					module = e1.Name.Local
				} else {
					//should be expanded to support multiple nested items
					nested = e1.Name.Local
					nestedGroup.XMLName = xml.Name{Local: nested}
				}
			}
		case xml.EndElement:
			switch e1.Name.Local {
			case "row":
				//append 'currentRow' to 'values'
				values = append(values, Row{Number: len(values) + 1, Fields: currentRow})
			case nested:
				if nested != "" {
					//Got closing nested element
					//clear nested element
					nested = ""
					t := false
				cROW:
					for i := range currentRow {
						//if nestedField label is already in current row
						if currentRow[i].Label == nestedField.Label {
							nestedGroup.Number = len(currentRow[i].Data) + 1
							//append nestedGroup to currentRow[i]
							currentRow[i].Data = append(currentRow[i].Data, nestedGroup)
							nestedGroup = InternalGroup{}
							t = true
							break cROW
						}
					}
					//current row doesn't have the nestfields label in it
					if !t {
						//place nestedGroup in nestedField
						nestedGroup.Number = 1
						nestedField.Data = append(nestedField.Data, nestedGroup)
						nestedGroup = InternalGroup{}
						//append the 'nestedGroup' to the 'currentRow'
						currentRow = append(currentRow, nestedField)
						nestedField = FieldLabel{Label: nestedField.Label}
					}
				}
			case module:
				if module != "" {
					//got closing module element
					//should be safe to return the fields
					break PRIME
				}
			case "response":
				break PRIME
			case "result":
				break PRIME
			}
		}
	}
	data.Rows = values
	return data, nil
}

func checkForInternalGroup(decoder *xml.Decoder, v xml.StartElement) bool {
	found := false
	for {
		t, err := decoder.Token()
		if err != nil && err != io.EOF {
			fmt.Println("TOKEN ERROR: check internal group")
			return false
		}
		if t == nil {
			return false
		}
		switch e := t.(type) {
		case xml.EndElement:
			if found {
				return false
			}
		case xml.StartElement:
			if found {
				name := e.Name.Local
				switch name {
				case "FL":
					return false
				default:

					return true
				}
			}
			if e.Name.Local == v.Name.Local {
				same := true
				//iterate attributes on new token
			ATTR_1:
				for _, a := range e.Attr {
					//iterate attributes on old token
					//ATTR_2:
					for _, b := range v.Attr {
						//if both tokens have attr with same name
						if a.Name.Local == b.Name.Local {
							//and have same value
							if a.Value == b.Value {
								//test next attribute on new token
								continue ATTR_1
							}
						}
					}
					//tested all attributes on this token
					// and they are not identical
					same = false
					break ATTR_1
				}
				if same {
					found = true
				}
			}
		}
	}
}

//This is the base request XMLData
// Name is set by the 'CrmModule' provided to the function
type XMLData struct {
	XMLName xml.Name
	Rows    []Row
	Error   XMLError
}

func (x *XMLData) addRow(c crmData, i int) {
	//Create a Row
	r := Row{Number: i}
	f, err := getValuesFromStruct(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.Fields = f
	x.Rows = append(x.Rows, r)
}

func (x *XMLData) encode() string {
	//Encode the XML
	b, err := xml.Marshal(x)
	if err != nil {
		fmt.Println("Got error encoding XML")
		fmt.Println(err.Error())
		return ""
	}

	return string(b)
}

//Each record must be in a row, with a sequential number scheme
type Row struct {
	XMLName xml.Name `xml:"row"`
	Number  int      `xml:"no,attr"`
	ID      string   `xml:"id,attr,omitempty"`
	PL      string   `xml:"pl,attr,omitempty"`
	SL      string   `xml:"sl,attr,omitempty"`
	GT      string   `xml:"gt,attr,omitempty"`
	Value   string   `xml:",chardata"`
	Fields  []FieldLabel
}

// The fields of each record as such `<FL val="">{{CharData}}</FL>`
type FieldLabel struct {
	XMLName xml.Name `xml:"FL"`
	Label   string   `xml:"val,attr"`
	Value   []byte   `xml:",chardata"`
	Data    []InternalGroup
}

type InternalGroup struct {
	XMLName xml.Name
	Number  int `xml:"no,attr"`
	Fields  []FieldLabel
}

func (f *FieldLabel) decode(decoder *xml.Decoder, v xml.Token) {
	e := v.(xml.StartElement)
	err := decoder.DecodeElement(f, &e)
	if err != nil {
		fmt.Println("Failed to decode field label: ", err.Error())
		fmt.Println("Got ", e.Name.Local)
		fmt.Println(v)
	}
}

func getFieldLabelByName(f []FieldLabel, name string) FieldLabel {
	if strings.Contains(name, ">") {
		tags := strings.Split(name, ">")
		name = tags[0]
	}
	for _, a := range f {
		if a.Label == name {
			return a
		}
	}
	return FieldLabel{}
}

func removeLabelByName(f []FieldLabel, name string) []FieldLabel {
	for i, a := range f {
		if a.Label == name {
			f = append(f[:i], f[i+1:]...)
			return f
		}
	}
	return f
}

type XMLError struct {
	XMLName xml.Name
	Code    int    `xml:"code"`
	Message string `xml:"message"`
}
