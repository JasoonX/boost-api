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

// AddPhoneRelationshipsUser User who owns this phone
type AddPhoneRelationshipsUser struct {
	Data *RelatedEntity `json:"data,omitempty"`
}

// NewAddPhoneRelationshipsUser instantiates a new AddPhoneRelationshipsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPhoneRelationshipsUser() *AddPhoneRelationshipsUser {
	this := AddPhoneRelationshipsUser{}
	return &this
}

// NewAddPhoneRelationshipsUserWithDefaults instantiates a new AddPhoneRelationshipsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPhoneRelationshipsUserWithDefaults() *AddPhoneRelationshipsUser {
	this := AddPhoneRelationshipsUser{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddPhoneRelationshipsUser) GetData() RelatedEntity {
	if o == nil || o.Data == nil {
		var ret RelatedEntity
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPhoneRelationshipsUser) GetDataOk() (*RelatedEntity, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddPhoneRelationshipsUser) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RelatedEntity and assigns it to the Data field.
func (o *AddPhoneRelationshipsUser) SetData(v RelatedEntity) {
	o.Data = &v
}

func (o AddPhoneRelationshipsUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAddPhoneRelationshipsUser struct {
	value *AddPhoneRelationshipsUser
	isSet bool
}

func (v NullableAddPhoneRelationshipsUser) Get() *AddPhoneRelationshipsUser {
	return v.value
}

func (v *NullableAddPhoneRelationshipsUser) Set(val *AddPhoneRelationshipsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPhoneRelationshipsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPhoneRelationshipsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPhoneRelationshipsUser(val *AddPhoneRelationshipsUser) *NullableAddPhoneRelationshipsUser {
	return &NullableAddPhoneRelationshipsUser{value: val, isSet: true}
}

func (v NullableAddPhoneRelationshipsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPhoneRelationshipsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
