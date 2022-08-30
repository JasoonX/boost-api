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

// NewsGet403ResponseError struct for NewsGet403ResponseError
type NewsGet403ResponseError struct {
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewNewsGet403ResponseError instantiates a new NewsGet403ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewsGet403ResponseError() *NewsGet403ResponseError {
	this := NewsGet403ResponseError{}
	return &this
}

// NewNewsGet403ResponseErrorWithDefaults instantiates a new NewsGet403ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewsGet403ResponseErrorWithDefaults() *NewsGet403ResponseError {
	this := NewsGet403ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *NewsGet403ResponseError) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsGet403ResponseError) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *NewsGet403ResponseError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *NewsGet403ResponseError) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NewsGet403ResponseError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsGet403ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NewsGet403ResponseError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NewsGet403ResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o NewsGet403ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableNewsGet403ResponseError struct {
	value *NewsGet403ResponseError
	isSet bool
}

func (v NullableNewsGet403ResponseError) Get() *NewsGet403ResponseError {
	return v.value
}

func (v *NullableNewsGet403ResponseError) Set(val *NewsGet403ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableNewsGet403ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableNewsGet403ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewsGet403ResponseError(val *NewsGet403ResponseError) *NullableNewsGet403ResponseError {
	return &NullableNewsGet403ResponseError{value: val, isSet: true}
}

func (v NullableNewsGet403ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewsGet403ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}