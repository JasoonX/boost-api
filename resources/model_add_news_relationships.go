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

// AddNewsRelationships Relationships of the news to add.
type AddNewsRelationships struct {
	Author AddNewsRelationshipsAuthor `json:"author"`
}

// NewAddNewsRelationships instantiates a new AddNewsRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddNewsRelationships(author AddNewsRelationshipsAuthor) *AddNewsRelationships {
	this := AddNewsRelationships{}
	this.Author = author
	return &this
}

// NewAddNewsRelationshipsWithDefaults instantiates a new AddNewsRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddNewsRelationshipsWithDefaults() *AddNewsRelationships {
	this := AddNewsRelationships{}
	return &this
}

// GetAuthor returns the Author field value
func (o *AddNewsRelationships) GetAuthor() AddNewsRelationshipsAuthor {
	if o == nil {
		var ret AddNewsRelationshipsAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *AddNewsRelationships) GetAuthorOk() (*AddNewsRelationshipsAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *AddNewsRelationships) SetAuthor(v AddNewsRelationshipsAuthor) {
	o.Author = v
}

func (o AddNewsRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["author"] = o.Author
	}
	return json.Marshal(toSerialize)
}

type NullableAddNewsRelationships struct {
	value *AddNewsRelationships
	isSet bool
}

func (v NullableAddNewsRelationships) Get() *AddNewsRelationships {
	return v.value
}

func (v *NullableAddNewsRelationships) Set(val *AddNewsRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAddNewsRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAddNewsRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddNewsRelationships(val *AddNewsRelationships) *NullableAddNewsRelationships {
	return &NullableAddNewsRelationships{value: val, isSet: true}
}

func (v NullableAddNewsRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddNewsRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
