package crm

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
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
