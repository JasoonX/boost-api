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

// AddEmailRelationshipsUser struct for AddEmailRelationshipsUser
type AddEmailRelationshipsUser struct {
	Data *RelatedEntity `json:"data,omitempty"`
}

// NewAddEmailRelationshipsUser instantiates a new AddEmailRelationshipsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEmailRelationshipsUser() *AddEmailRelationshipsUser {
	this := AddEmailRelationshipsUser{}
	return &this
}

// NewAddEmailRelationshipsUserWithDefaults instantiates a new AddEmailRelationshipsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEmailRelationshipsUserWithDefaults() *AddEmailRelationshipsUser {
	this := AddEmailRelationshipsUser{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddEmailRelationshipsUser) GetData() RelatedEntity {
	if o == nil || o.Data == nil {
		var ret RelatedEntity
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEmailRelationshipsUser) GetDataOk() (*RelatedEntity, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddEmailRelationshipsUser) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RelatedEntity and assigns it to the Data field.
func (o *AddEmailRelationshipsUser) SetData(v RelatedEntity) {
	o.Data = &v
}

func (o AddEmailRelationshipsUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAddEmailRelationshipsUser struct {
	value *AddEmailRelationshipsUser
	isSet bool
}

func (v NullableAddEmailRelationshipsUser) Get() *AddEmailRelationshipsUser {
	return v.value
}

func (v *NullableAddEmailRelationshipsUser) Set(val *AddEmailRelationshipsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEmailRelationshipsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEmailRelationshipsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEmailRelationshipsUser(val *AddEmailRelationshipsUser) *NullableAddEmailRelationshipsUser {
	return &NullableAddEmailRelationshipsUser{value: val, isSet: true}
}

func (v NullableAddEmailRelationshipsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEmailRelationshipsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}