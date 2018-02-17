package crm

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//This is the base request XMLData
// Name is set by the 'CrmModule' provided to the function
type XMLData struct {
	XMLName xml.Name
	Rows    []Row
	Error   XMLError
}

// addRow will get the fields from the provided data and add the fields to the new row.
func (x *XMLData) addRow(c crmData, i int) {
	//Create a Row
	r := Row{Number: i}
	//get fields from data
	f, err := getValuesFromStruct(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	//add fields to row
	r.Fields = f
	//add new row to row list
	x.Rows = append(x.Rows, r)
}

// encode will iterate over the fields in each row
func (x *XMLData) encode() string {
	x2 := x
	// iterate over the rows
	for i := range x2.Rows {
		fields := x2.Rows[i].Fields
		//iterate over the rows fields
	FIELDS:
		for j := range x2.Rows[i].Fields {
			//iterate over the fields Tags
			for _, tag := range x2.Rows[i].Fields[j].Tags {
				//if the 'strip' tag exists
				if tag == "strip" {
					//iterate over the fields reference
					for k := range fields {
						//if the fields label is same as reference fields label
						if fields[k].Label == x2.Rows[i].Fields[j].Label {
							//remove field with label from fields
							fields[len(fields)-1], fields[k] = fields[k], fields[len(fields)-1]
							fields = fields[:len(fields)-1]
							continue FIELDS
						}
					}
				}
			}
		}
		//replace the rows fields with the reference fields
		x2.Rows[i].Fields = fields
	}

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
	Tags    []string `xml:"-"`
}

type InternalGroup struct {
	XMLName xml.Name
	Number  int `xml:"no,attr"`
	Fields  []FieldLabel
}

// decode is used to get the data from inside a "FL" XML tag
func (f *FieldLabel) decode(decoder *xml.Decoder, v xml.Token) {
	e := v.(xml.StartElement)
	err := decoder.DecodeElement(f, &e)
	if err != nil {
		fmt.Println("Failed to decode field label: ", err.Error())
		fmt.Println("Got ", e.Name.Local)
		fmt.Println(v)
	}
}

// getFieldLabelByName will return a named field label
func getFieldLabelByName(f []FieldLabel, name string) FieldLabel {
	if strings.Contains(name, ">") {
		tags := strings.Split(name, ">")
		name = tags[0]
	}
	//iterate over the provided FieldLabels
	for _, a := range f {
		//if the fields label is the same as the provided name, or the iterpreted name from the tag
		if a.Label == name {
			return a
		}
	}
	return FieldLabel{}
}

// removeLabelByName will remove a named field from a list of fields
func removeLabelByName(f []FieldLabel, name string) []FieldLabel {
	// iterate over the provided fields
	for i, a := range f {
		// if the fields label is the name to be removed
		if a.Label == name {
			f = append(f[:i], f[i+1:]...)
			return f
		}
	}
	return f
}

// XMLError is returned from Zoho when an Error occurs internally
type XMLError struct {
	XMLName xml.Name
	Code    int    `xml:"code"`
	Message string `xml:"message"`
}

// getvaluesFromXML will create a valid data type from the XML returned by Zoho
func getValuesFromXML(b []byte) (XMLData, error) {
	data := XMLData{} //all data
	values := []Row{} // list of rows
	currentRow := []FieldLabel{} //fields in row
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
			fmt.Println("Got error on prime 'tokener'")
			log.Fatal(err)
		}
		// no more tokens
		if t1 == nil {
			break PRIME
		}
		// Inspect the type of the token just read.
		switch e1 := t1.(type) {
		case xml.StartElement:
			switch e1.Name.Local {
			case "response", "result":
				//these are just wrapper words and not particularly useful
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
				// set the module value
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
			case "response", "result":
				break PRIME
			}
		}
	}
	//set the datas 'Rows'
	data.Rows = values
	//return the data
	return data, nil
}

//checkForInternalGroud will return true when the provided XML startelement
// is followed by a set of elements deeper than the row>FL
// an example would be row>FL>products when retreiving a SalesOrder
func checkForInternalGroup(decoder *xml.Decoder, v xml.StartElement) bool {
	found := false
	for {
		t, err := decoder.Token()
		if err != nil && err != io.EOF {
			fmt.Println("TOKEN ERROR: check internal group")
			return false
		}
		//no more tokens
		if t == nil {
			return false
		}

		switch e := t.(type) {
		case xml.EndElement:
			//because we have an endelement after we found our token
			if found {
				return false
			}
		case xml.StartElement:
			if found {
				name := e.Name.Local
				switch name {
				case "FL":
					//got a 'FL' element inside our found element
					return false
				default:
					//got some other field element inside our found element
					return true
				}
			}
			// provided element name is same as new start element
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
				// the two tokens and all attributes are the same
				if same {
					found = true
				}
			}
		}
	}
}

//We expect to get a slice or pointer to slice here
//iterate over the items in the slice
func decodeXML(b []byte, data crmData) (crmData, error) {
	dV := reflect.ValueOf(data)
	dT := reflect.TypeOf(data)

	if dV.Kind() == reflect.Ptr {
		dV = reflect.Indirect(dV)
		dT = reflect.TypeOf(dV.Interface())
	}

	x, err := getValuesFromXML(b)
	if err != nil {
		return nil, err
	}

	//There was an error in the XML
	if x.Error != (XMLError{}) {
		if x.Error.XMLName.Local == "error" {
			//make a 'crmData' of errortype
			return CrmError{Type: x.Error.XMLName.Local, Code: x.Error.Code, Message: x.Error.Message},
				fmt.Errorf("Zoho CRM returned an Error: Code %d: %s", x.Error.Code, x.Error.Message)
		} else if x.Error.XMLName.Local == "nodata" {
			return dV.Interface().(crmData), nil
		}
	}

	for _, a := range x.Rows {
		switch dV.Kind() {
		case reflect.Slice:
			uT := dV.Type().Elem()
			item := reflect.New(uT)

			err = fillStructFromValues(a.Fields, item.Interface())
			if err != nil {
				return nil, err
			}
			dV = reflect.Append(dV, reflect.Indirect(item))
		case reflect.Struct:
			item := reflect.New(dT)

			err = fillStructFromValues(a.Fields, item.Interface())
			if err != nil {
				return nil, err
			}
			dV.Set(item)

			return dV.Interface().(crmData), nil
		}
	}

	return dV.Interface().(crmData), nil
}

//Provide the fields and a pointer to the structure and we will fill each structure field
// with the corresponding value given fields[i].Label == data.Field(i).Tag.Get("zoho")[0]
func fillStructFromValues(fields []FieldLabel, data interface{}) error {
	dV := reflect.ValueOf(data)
	dT := reflect.TypeOf(data)
	//check for pointer
	if dV.Kind() == reflect.Ptr {
		//replace dV with the value it points to
		dV = reflect.Indirect(dV)
		dT = reflect.TypeOf(dV.Interface())
	}

	switch dV.Kind() {
	case reflect.Struct:
		// Iterate over the struct fields
	FIELDS:
		for i := 0; i < dV.NumField(); i++ {
			field := dV.Field(i)
			//check for pointer
			if field.Kind() == reflect.Ptr {
				//replace dV with the value it points to
				field = reflect.Indirect(field)
			}

			if !field.IsValid() || !field.CanSet() {
				fmt.Println("Field invalid", dT.Field(i).Name)
				fmt.Println(dV.Type().Name(), dV.Type().PkgPath())
				continue FIELDS
			}
			//get the zoho tag
			fieldTag := dT.Field(i).Tag.Get("zoho")
			fieldTags := strings.Split(fieldTag, ",")
			//get the first portion of the zoho tag
			tag := fieldTags[0]
			f := getFieldLabelByName(fields, tag)
			if f.Label == "" {
				//				fmt.Println("Field had no label", tag)
				continue FIELDS
			}
			fields = removeLabelByName(fields, f.Label)

			//if fields type is string, int, float, etc.
			//getFieldLabelByName(fields, tag[0])
			//just parse the value from the fields[i].Value
			switch field.Kind() {
			case reflect.Int:
				d, err := strconv.ParseInt(string(f.Value), 10, 64)
				if err != nil {
					if len(f.Value) > 0 {
						field.SetInt(1)
					} else {
						field.SetInt(0)
					}
				}
				if !field.OverflowInt(d) {
					field.SetInt(d)
				}
			case reflect.Int64:
				d, err := strconv.ParseInt(string(f.Value), 10, 64)
				if err != nil {
					if len(f.Value) > 0 {
						field.SetInt(1)
					} else {
						field.SetInt(0)
					}
				}
				if !field.OverflowInt(d) {
					field.SetInt(d)
				}
			case reflect.Float64:
				fl, err := strconv.ParseFloat(string(f.Value), 64)
				if err != nil {
					if len(f.Value) > 0 {
						field.SetFloat(1)
					} else {
						field.SetFloat(0)
					}
				}
				if !field.OverflowFloat(fl) {
					field.SetFloat(fl)
				}
			case reflect.String:
				field.SetString(string(f.Value))
			case reflect.Bool:
				b, err := strconv.ParseBool(string(f.Value))
				if err != nil {
					if len(f.Value) > 0 {
						field.SetBool(true)
					} else {
						field.SetBool(false)
					}
				}
				field.SetBool(b)
			case reflect.Struct:
				//if the fields type is struct
				//if the type is time.Time or time.Duration
				if field.Type().Name() == "Time" && field.Type().PkgPath() == "time" {
					//just parse the time with layout and add the structure
					t, err := time.Parse("2006-01-02 15:04:05", string(f.Value))
					if err != nil {
						t, err := time.Parse("2006-01-02", string(f.Value))
						if err != nil {
							return fmt.Errorf("Got error parsing Time format: %s", err.Error())
						}
						field.Set(reflect.ValueOf(t))
					} else {
						field.Set(reflect.ValueOf(t))
					}
				} else {
					//pass fields[i].Data[0].Fields and a pointer to the field to this function
					if len(f.Data) > 0 {
						err := fillStructFromValues(f.Data[0].Fields, field.Addr().Interface())
						if err != nil {
							return err
						}
					}
				}

			case reflect.Slice:
				//if field type is slice
				//iterate over the fields[i].Data
				for _, a := range f.Data {
					// get a pointer to type underlying slice
					uT := field.Type().Elem()
					item := reflect.New(uT)

					//pass fields[i].Data[j].Fields and pointer to item to this function
					err := fillStructFromValues(a.Fields, item.Interface())
					if err != nil {
						return err
					}
					//append the item to the slice
					field = reflect.Append(field, reflect.Indirect(item))
					dV.Field(i).Set(field)
				}
			}

		} //FIELDS END
	}

	//put all remaining fields in the data struct
	e := ExtraFields{}
	for _, a := range fields {
		e[a.Label] = string(a.Value)
	}
	ev := reflect.ValueOf(e)
	dV.FieldByName("ExtraFields").Set(ev)

	return nil
}
