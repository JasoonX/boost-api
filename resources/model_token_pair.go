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

// TokenPair struct for TokenPair
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewTokenPair instantiates a new TokenPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPair(accessToken string, refreshToken string) *TokenPair {
	this := TokenPair{}
	this.AccessToken = accessToken
	this.RefreshToken = refreshToken
	return &this
}

// NewTokenPairWithDefaults instantiates a new TokenPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPairWithDefaults() *TokenPair {
	this := TokenPair{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *TokenPair) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TokenPair) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TokenPair) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *TokenPair) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *TokenPair) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *TokenPair) SetRefreshToken(v string) {
	o.RefreshToken = v
}

func (o TokenPair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	return json.Marshal(toSerialize)
}

type NullableTokenPair struct {
	value *TokenPair
	isSet bool
}

func (v NullableTokenPair) Get() *TokenPair {
	return v.value
}

func (v *NullableTokenPair) Set(val *TokenPair) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPair) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPair(val *TokenPair) *NullableTokenPair {
	return &NullableTokenPair{value: val, isSet: true}
}

func (v NullableTokenPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}