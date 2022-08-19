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

// AddEmailAttributes struct for AddEmailAttributes
type AddEmailAttributes struct {
	Email string `json:"email"`
}

// NewAddEmailAttributes instantiates a new AddEmailAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEmailAttributes(email string) *AddEmailAttributes {
	this := AddEmailAttributes{}
	this.Email = email
	return &this
}

// NewAddEmailAttributesWithDefaults instantiates a new AddEmailAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEmailAttributesWithDefaults() *AddEmailAttributes {
	this := AddEmailAttributes{}
	return &this
}

// GetEmail returns the Email field value
func (o *AddEmailAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AddEmailAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AddEmailAttributes) SetEmail(v string) {
	o.Email = v
}

func (o AddEmailAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableAddEmailAttributes struct {
	value *AddEmailAttributes
	isSet bool
}

func (v NullableAddEmailAttributes) Get() *AddEmailAttributes {
	return v.value
}

func (v *NullableAddEmailAttributes) Set(val *AddEmailAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEmailAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEmailAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEmailAttributes(val *AddEmailAttributes) *NullableAddEmailAttributes {
	return &NullableAddEmailAttributes{value: val, isSet: true}
}

func (v NullableAddEmailAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEmailAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
