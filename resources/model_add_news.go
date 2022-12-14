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

// AddNews struct for AddNews
type AddNews struct {
	Type          EntityType           `json:"type"`
	Attributes    AddNewsAttributes    `json:"attributes"`
	Relationships AddNewsRelationships `json:"relationships"`
}

// NewAddNews instantiates a new AddNews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddNews(type_ EntityType, attributes AddNewsAttributes, relationships AddNewsRelationships) *AddNews {
	this := AddNews{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAddNewsWithDefaults instantiates a new AddNews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddNewsWithDefaults() *AddNews {
	this := AddNews{}
	return &this
}

// GetType returns the Type field value
func (o *AddNews) GetType() EntityType {
	if o == nil {
		var ret EntityType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddNews) GetTypeOk() (*EntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddNews) SetType(v EntityType) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AddNews) GetAttributes() AddNewsAttributes {
	if o == nil {
		var ret AddNewsAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AddNews) GetAttributesOk() (*AddNewsAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AddNews) SetAttributes(v AddNewsAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AddNews) GetRelationships() AddNewsRelationships {
	if o == nil {
		var ret AddNewsRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AddNews) GetRelationshipsOk() (*AddNewsRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AddNews) SetRelationships(v AddNewsRelationships) {
	o.Relationships = v
}

func (o AddNews) MarshalJSON() ([]byte, error) {
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

type NullableAddNews struct {
	value *AddNews
	isSet bool
}

func (v NullableAddNews) Get() *AddNews {
	return v.value
}

func (v *NullableAddNews) Set(val *AddNews) {
	v.value = val
	v.isSet = true
}

func (v NullableAddNews) IsSet() bool {
	return v.isSet
}

func (v *NullableAddNews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddNews(val *AddNews) *NullableAddNews {
	return &NullableAddNews{value: val, isSet: true}
}

func (v NullableAddNews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddNews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
