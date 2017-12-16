package crm

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//All CRM data types should implement this interface
type crmData interface {
	writeXML() string
	//	String() string
}

type ExtraFields map[string]interface{}

func (e ExtraFields) GetFieldByName(name string) interface{} {
	for k, v := range e {
		if k == name {
			return v
		}
	}
	return nil
}

//getValuesFromStruct expects a struct and using reflect
// makes a []FieldLabel for the struct
// only fields with label `zoho:"Name"` are supported
// nest structs and slices are supported by using the FieldLabel.Data
func getValuesFromStruct(data interface{}) ([]FieldLabel, error) {
	dV := reflect.ValueOf(data)
	dT := reflect.TypeOf(data)
	//check for pointer
	if dV.Kind() == reflect.Ptr {
		//replace dV with the value it points to
		dV = reflect.Indirect(dV)
		dT = dV.Type()
	}

	values := []FieldLabel{}
	switch dV.Kind() {
	case reflect.Struct:
		// Iterate over the struct fields
	FIELDS:
		for i := 0; i < dV.NumField(); i++ {
			field := dV.Field(i)
			fieldT := dT.Field(i)
			//check for pointer
			if field.Kind() == reflect.Ptr {
				//replace dV with the value it points to
				field = reflect.Indirect(field)
				fmt.Println(fieldT)
			}
			fieldTag := dT.Field(i).Tag.Get("zoho")
			fieldTags := strings.Split(fieldTag, ",")
			tag := fieldTags[0]
			//TAGS:
			f := FieldLabel{Label: tag}
			switch field.Kind() {
			case reflect.Int:
				f.Value = []byte(fmt.Sprintf("%d", field.Interface().(int)))
			case reflect.Int64:
				f.Value = []byte(fmt.Sprintf("%d", field.Interface().(int64)))
			case reflect.Float64:
				f.Value = []byte(fmt.Sprintf("%f", field.Interface().(float64)))
			case reflect.String:
				f.Value = []byte(field.Interface().(string))
			case reflect.Bool:
				f.Value = []byte(fmt.Sprintf("%t", field.Interface().(bool)))
			case reflect.Struct, reflect.Slice:
				//Get the wrapper tag for the interal fields
				iTag := ""
				if strings.Contains(tag, ">") {
					tagStr := strings.Split(tag, ">")
					tag = tagStr[0]
					f.Label = tag
					iTag = tagStr[1]
				}
				//Field is a nest struct
				if field.Kind() == reflect.Struct {
					//check if that struct is actually a time.Time or time.Duration
					if field.Type().Name() == "Time" && field.Type().PkgPath() == "time" {
						//This is a time type so we should parse it instead of walking the structure
						if field.Interface().(time.Time).IsZero() {
							f.Value = []byte(fmt.Sprint(time.Time{}))
							continue FIELDS
						}
						if t, ok := field.Interface().(time.Time); ok {
							f.Value = []byte(t.Format("2006-01-02 15:04:05"))
						} else {
							t, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprint(field.Interface()))
							if err != nil {
								t, err := time.Parse("2006-01-02", fmt.Sprint(field.Interface()))
								if err != nil {
									return nil, fmt.Errorf("Got error parsing time format: %s", err.Error())
								}
								f.Value = []byte(fmt.Sprint(t))
							} else {
								//Set the 'Field Labels' value
								f.Value = []byte(t.Format("2006-01-02 15:04:05"))
							}
						}
					} else {

						//This is a struct but not from pkg 'time'
						//submit this field to this function again
						fields, err := getValuesFromStruct(field.Interface())
						if err != nil {
							return nil, err
						}
						fmt.Println(fields)
						//apply the returned fields to an 'Internal Group' on the Field Label
						f.Data = append(f.Data, InternalGroup{
							XMLName: xml.Name{Local: iTag},
							Number:  1,
							Fields:  fields,
						})
					}
				}
				//the kind of the field is a slice
				if field.Kind() == reflect.Slice {
					//iterate over the slices items
					for i := 0; i < field.Len(); i++ {
						//submit each item in the slice to this function
						//MIGHT HAVE TO DO A TYPE CHECK FOR SLICES OF NON STRUC TYPE
						fields, err := getValuesFromStruct(field.Index(i).Interface())
						if err != nil {
							return nil, err
						}
						//apply the returned fields to an InternalGroup and append as data to the field Label
						f.Data = append(f.Data, InternalGroup{
							XMLName: xml.Name{Local: iTag},
							Number:  i,
							Fields:  fields,
						})
					}
				}
			}
			//append the FieldLabel for this field to the list of field labels
			values = append(values, f)
		} //NumFields END
	case reflect.Slice:
		//Slice should not be sent to this function. Only items in a slice
		return []FieldLabel{}, errors.New("Slices are not supported: provide the items of the slice instead")
	default:
		return []FieldLabel{}, errors.New("Unsupported kind: " + dV.Kind().String())
	}

	return values, nil
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

	if len(fields) > 0 {
		fmt.Println("FIELDS REMAINING")
	}
	for _, a := range fields {

		fmt.Printf("FIELD:\t'%s'\twith Value:\t'%s'\n", a.Label, string(a.Value))
	}

	return nil
}

type crmModule string

//Proper names for CRM modules
const (
	accountsModule       crmModule = "Accounts"
	callsModule          crmModule = "Calls"
	campaignsModule      crmModule = "Campaigns"
	casesModule          crmModule = "Cases"
	contactsModule       crmModule = "Contacts"
	eventsModule         crmModule = "Events"
	invoicesModule       crmModule = "Invoices"
	leadsModule          crmModule = "Leads"
	potentialsModule     crmModule = "Potentials"
	priceBooksModule     crmModule = "PriceBooks"
	productsModule       crmModule = "Products"
	purchaseOrdersModule crmModule = "PurchaseOrders"
	quotesModule         crmModule = "Quotes"
	salesOrdersModule    crmModule = "SalesOrders"
	solutionsModule      crmModule = "Solutions"
	tasksModule          crmModule = "Tasks"
	vendorsModule        crmModule = "Vendors"
)

func (c crmModule) getType() crmData {
	switch c {
	case accountsModule:
		return &Accounts{}
	case callsModule:
		return &Calls{}
	case campaignsModule:
		return &Campaigns{}
	case casesModule:
		return &Cases{}
	case contactsModule:
		return &Contacts{}
	case eventsModule:
		return &Events{}
	case invoicesModule:
		return &Invoices{}
	case leadsModule:
		return &Leads{}
	case potentialsModule:
		return &Potentials{}
	case priceBooksModule:
		return &PriceBooks{}
	case productsModule:
		return &Products{}
	case purchaseOrdersModule:
		return &PurchaseOrders{}
	case quotesModule:
		return &Quotes{}
	case salesOrdersModule:
		return &SalesOrders{}
	case solutionsModule:
		return &Solutions{}
	case tasksModule:
		return &Tasks{}
	case vendorsModule:
		return &Vendors{}
	default:
		return nil
	}
}
