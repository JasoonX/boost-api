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

// UsersPostRequest struct for UsersPostRequest
type UsersPostRequest struct {
	Data AddUser `json:"data"`
}

// NewUsersPostRequest instantiates a new UsersPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersPostRequest(data AddUser) *UsersPostRequest {
	this := UsersPostRequest{}
	this.Data = data
	return &this
}

// NewUsersPostRequestWithDefaults instantiates a new UsersPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersPostRequestWithDefaults() *UsersPostRequest {
	this := UsersPostRequest{}
	return &this
}

// GetData returns the Data field value
func (o *UsersPostRequest) GetData() AddUser {
	if o == nil {
		var ret AddUser
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UsersPostRequest) GetDataOk() (*AddUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UsersPostRequest) SetData(v AddUser) {
	o.Data = v
}

func (o UsersPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUsersPostRequest struct {
	value *UsersPostRequest
	isSet bool
}

func (v NullableUsersPostRequest) Get() *UsersPostRequest {
	return v.value
}

func (v *NullableUsersPostRequest) Set(val *UsersPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersPostRequest(val *UsersPostRequest) *NullableUsersPostRequest {
	return &NullableUsersPostRequest{value: val, isSet: true}
}

func (v NullableUsersPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
