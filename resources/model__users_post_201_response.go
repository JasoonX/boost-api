/*
Boost API

This is a REST API of the Boost App.

API version: 0.0.1-beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// UsersPost201Response struct for UsersPost201Response
type UsersPost201Response struct {
	Status *Status                   `json:"status,omitempty"`
	Meta   map[string]interface{}    `json:"meta,omitempty"`
	Errors []Error                   `json:"errors,omitempty"`
	Data   *UsersPost201ResponseData `json:"data,omitempty"`
}

// NewUsersPost201Response instantiates a new UsersPost201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersPost201Response() *UsersPost201Response {
	this := UsersPost201Response{}
	return &this
}

// NewUsersPost201ResponseWithDefaults instantiates a new UsersPost201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersPost201ResponseWithDefaults() *UsersPost201Response {
	this := UsersPost201Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UsersPost201Response) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersPost201Response) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UsersPost201Response) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *UsersPost201Response) SetStatus(v Status) {
	o.Status = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UsersPost201Response) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersPost201Response) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UsersPost201Response) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *UsersPost201Response) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UsersPost201Response) GetErrors() []Error {
	if o == nil || o.Errors == nil {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersPost201Response) GetErrorsOk() ([]Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UsersPost201Response) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *UsersPost201Response) SetErrors(v []Error) {
	o.Errors = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UsersPost201Response) GetData() UsersPost201ResponseData {
	if o == nil || o.Data == nil {
		var ret UsersPost201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersPost201Response) GetDataOk() (*UsersPost201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UsersPost201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UsersPost201ResponseData and assigns it to the Data field.
func (o *UsersPost201Response) SetData(v UsersPost201ResponseData) {
	o.Data = &v
}

func (o UsersPost201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUsersPost201Response struct {
	value *UsersPost201Response
	isSet bool
}

func (v NullableUsersPost201Response) Get() *UsersPost201Response {
	return v.value
}

func (v *NullableUsersPost201Response) Set(val *UsersPost201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersPost201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersPost201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersPost201Response(val *UsersPost201Response) *NullableUsersPost201Response {
	return &NullableUsersPost201Response{value: val, isSet: true}
}

func (v NullableUsersPost201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersPost201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
