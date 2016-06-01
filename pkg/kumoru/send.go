/*
Copyright 2016 Kumoru.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kumoru

import (
	"bytes"
	"encoding/json"
	"net/url"
	"reflect"
	"strings"
)

// Send a string or a struct as parameters
func (k *Client) Send(content interface{}) *Client {
	switch v := reflect.ValueOf(content); v.Kind() {
	case reflect.String:
		k.SendString(v.String())
	case reflect.Struct:
		k.SendStruct(v.Interface())
	default:
	}
	return k
}

// SendString sends the information as a raw string
func (k *Client) SendString(content string) *Client {
	if !k.BounceToRawString {
		var val interface{}
		d := json.NewDecoder(strings.NewReader(content))
		d.UseNumber()
		if err := d.Decode(&val); err == nil {
			switch v := reflect.ValueOf(val); v.Kind() {
			case reflect.Map:
				for key, v := range val.(map[string]interface{}) {
					k.Data[key] = v
				}
			default:
				k.BounceToRawString = true
			}
		} else if formVal, err := url.ParseQuery(content); err == nil {
			for key := range formVal {
				// make it array if already have key
				if val, ok := k.Data[key]; ok {
					var strArray []string
					strArray = append(strArray, formVal.Get(key))
					// check if previous data is one string or array
					switch oldValue := val.(type) {
					case []string:
						strArray = append(strArray, oldValue...)
					case string:
						strArray = append(strArray, oldValue)
					}
					k.Data[key] = strArray
				} else {
					// make it just string if does not already have same key
					k.Data[key] = formVal.Get(key)
				}
			}
			k.TargetType = "form"
		} else {
			k.BounceToRawString = true
		}
	}
	// Dump all contents to RawString in case in the end user doesn't want json or form.
	k.RawString += content
	return k
}

// SendStruct converts a struct to parameters
func (k *Client) SendStruct(content interface{}) *Client {
	if marshalContent, err := json.Marshal(content); err != nil {
		k.Errors = append(k.Errors, err)
	} else {
		var val map[string]interface{}
		d := json.NewDecoder(bytes.NewBuffer(marshalContent))
		d.UseNumber()
		if err := d.Decode(&val); err != nil {
			k.Errors = append(k.Errors, err)
		} else {
			for key, v := range val {
				k.Data[key] = v
			}
		}
	}
	return k
}

// SendSlice appends an array into k.SliceData
func (k *Client) SendSlice(content []interface{}) *Client {
	k.SliceData = append(k.SliceData, content...)
	return k
}
