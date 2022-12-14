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

// AddEmail struct for AddEmail
type AddEmail struct {
	Type          EntityType            `json:"type"`
	Attributes    AddEmailAttributes    `json:"attributes"`
	Relationships AddEmailRelationships `json:"relationships"`
}

// NewAddEmail instantiates a new AddEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEmail(type_ EntityType, attributes AddEmailAttributes, relationships AddEmailRelationships) *AddEmail {
	this := AddEmail{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAddEmailWithDefaults instantiates a new AddEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEmailWithDefaults() *AddEmail {
	this := AddEmail{}
	return &this
}

// GetType returns the Type field value
func (o *AddEmail) GetType() EntityType {
	if o == nil {
		var ret EntityType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddEmail) GetTypeOk() (*EntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddEmail) SetType(v EntityType) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AddEmail) GetAttributes() AddEmailAttributes {
	if o == nil {
		var ret AddEmailAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AddEmail) GetAttributesOk() (*AddEmailAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AddEmail) SetAttributes(v AddEmailAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AddEmail) GetRelationships() AddEmailRelationships {
	if o == nil {
		var ret AddEmailRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AddEmail) GetRelationshipsOk() (*AddEmailRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AddEmail) SetRelationships(v AddEmailRelationships) {
	o.Relationships = v
}

func (o AddEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableAddEmail struct {
	value *AddEmail
	isSet bool
}

func (v NullableAddEmail) Get() *AddEmail {
	return v.value
}

func (v *NullableAddEmail) Set(val *AddEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEmail(val *AddEmail) *NullableAddEmail {
	return &NullableAddEmail{value: val, isSet: true}
}

func (v NullableAddEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
