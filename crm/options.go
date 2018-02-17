package crm

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type BlankOptions map[string]string

// encodeURL must be better considered
// there are quite a number of rules that zoho values must follow
func (o BlankOptions) encodeURL(u *url.URL) error {
	noencode := ""
	val := url.Values{}
	for k, v := range o {
		//This is the ',noencode' option
		// the result of url.QueryEncode changing ' ' to '+'
		// when we need percent encoded spaces we pass a url.PathEncoded value here
		if strings.Contains(k, ",noencode") {
			key := strings.Replace(k, ",noencode", "", -1)
			//unfortunately pathEncode misses the '=' signs so we do that manually
			v = strings.Replace(v, "=", "%3D",-1)
			noencode += key + "=" + v + "&"
		} else {
			val.Add(k, v)
		}
	}
	u.RawQuery = val.Encode()
	u.RawQuery += "&" + noencode
	return nil
}

type optionEncoder interface {
	encodeURL(*url.URL) error
}

func encodeOptionsToURL(o optionEncoder, u *url.URL) error {
	vals := url.Values{}
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}

	//PRIME:
	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		value := v.Field(i).Interface()
		tag := v.Type().Field(i).Tag.Get("zoho")

		//split the tag on comma
		tags := strings.Split(tag, ",")
		encode := true
		//	TAGS:
		for _, a := range tags {
			if a == "required" {
				//if the underlying value of the field is empty
				if reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
					return fmt.Errorf("Error field '%s' is required for this request", name)
				}
			} else if strings.HasPrefix(a, "default>") {
				//if the underlying value of the field is empty
				if reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
					//Set the 'default' value for the URL item
					value = strings.TrimPrefix(a, "default>")
				}
			} else if a == "noencode" {
				encode = false
			} else {
				//if the underlying value of the field is NOT empty
				if !reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
					//Set the 'default' value for the URL item
					name = a
				}
			}
		}

		//If the field is set, if field is false check default tag
		switch v := value.(type) {
		case bool:
			//encode to the URL
			if encode {
				vals.Set(name, fmt.Sprintf("%t", v))
			}else {
				u.RawQuery += fmt.Sprintf("%s=%t&",name, v)
			}
		case int:
			if encode {
				vals.Set(name, fmt.Sprintf("%d", v))
			}else {
				u.RawQuery += fmt.Sprintf("%s=%d&",name, v)
			}
		case float64:
			if encode {
				vals.Set(name, fmt.Sprintf("%f", v))
			}else {
				u.RawQuery += fmt.Sprintf("%s=%f&",name, v)
			}
		case string:
			if encode {
				vals.Set(name, fmt.Sprintf("%s", v))
			}else {
				u.RawQuery += fmt.Sprintf("%s=%s&",name, v)
			}
		case time.Time:
			tm := time.Time(v)
			if !tm.IsZero() {
				if encode {
					vals.Set(name, fmt.Sprintf("%d-%d-%d %d:%d:%d",
						v.Year(), v.Month(), v.Day(),
						v.Hour(), v.Minute(), v.Second(),
					))
				}else {
					u.RawQuery += fmt.Sprintf("%s=%v&",name, v)
				}
			}
		case crmData:
			//get the items XML
			if encode {
				vals.Set(name, v.writeXML())
			} else {
				u.RawQuery += fmt.Sprintf("%s=%v&",name, v.writeXML())
			}
		}
	}

	vals.Set("version", "2")
	vals.Set("newFormat", "1")

	u.RawQuery += vals.Encode()
	return nil
}
