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

// HealthcheckGet500ResponseError struct for HealthcheckGet500ResponseError
type HealthcheckGet500ResponseError struct {
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewHealthcheckGet500ResponseError instantiates a new HealthcheckGet500ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthcheckGet500ResponseError() *HealthcheckGet500ResponseError {
	this := HealthcheckGet500ResponseError{}
	return &this
}

// NewHealthcheckGet500ResponseErrorWithDefaults instantiates a new HealthcheckGet500ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthcheckGet500ResponseErrorWithDefaults() *HealthcheckGet500ResponseError {
	this := HealthcheckGet500ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *HealthcheckGet500ResponseError) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckGet500ResponseError) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *HealthcheckGet500ResponseError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *HealthcheckGet500ResponseError) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HealthcheckGet500ResponseError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthcheckGet500ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HealthcheckGet500ResponseError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HealthcheckGet500ResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o HealthcheckGet500ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableHealthcheckGet500ResponseError struct {
	value *HealthcheckGet500ResponseError
	isSet bool
}

func (v NullableHealthcheckGet500ResponseError) Get() *HealthcheckGet500ResponseError {
	return v.value
}

func (v *NullableHealthcheckGet500ResponseError) Set(val *HealthcheckGet500ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthcheckGet500ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthcheckGet500ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthcheckGet500ResponseError(val *HealthcheckGet500ResponseError) *NullableHealthcheckGet500ResponseError {
	return &NullableHealthcheckGet500ResponseError{value: val, isSet: true}
}

func (v NullableHealthcheckGet500ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthcheckGet500ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
