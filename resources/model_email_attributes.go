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

// EmailAttributes struct for EmailAttributes
type EmailAttributes struct {
	Email      string `json:"email"`
	IsVerified *bool  `json:"is_verified,omitempty"`
	IsPrimary  *bool  `json:"is_primary,omitempty"`
}

// NewEmailAttributes instantiates a new EmailAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailAttributes(email string) *EmailAttributes {
	this := EmailAttributes{}
	this.Email = email
	return &this
}

// NewEmailAttributesWithDefaults instantiates a new EmailAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailAttributesWithDefaults() *EmailAttributes {
	this := EmailAttributes{}
	return &this
}

// GetEmail returns the Email field value
func (o *EmailAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmailAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmailAttributes) SetEmail(v string) {
	o.Email = v
}

// GetIsVerified returns the IsVerified field value if set, zero value otherwise.
func (o *EmailAttributes) GetIsVerified() bool {
	if o == nil || o.IsVerified == nil {
		var ret bool
		return ret
	}
	return *o.IsVerified
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAttributes) GetIsVerifiedOk() (*bool, bool) {
	if o == nil || o.IsVerified == nil {
		return nil, false
	}
	return o.IsVerified, true
}

// HasIsVerified returns a boolean if a field has been set.
func (o *EmailAttributes) HasIsVerified() bool {
	if o != nil && o.IsVerified != nil {
		return true
	}

	return false
}

// SetIsVerified gets a reference to the given bool and assigns it to the IsVerified field.
func (o *EmailAttributes) SetIsVerified(v bool) {
	o.IsVerified = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *EmailAttributes) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAttributes) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *EmailAttributes) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *EmailAttributes) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

func (o EmailAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.IsVerified != nil {
		toSerialize["is_verified"] = o.IsVerified
	}
	if o.IsPrimary != nil {
		toSerialize["is_primary"] = o.IsPrimary
	}
	return json.Marshal(toSerialize)
}

type NullableEmailAttributes struct {
	value *EmailAttributes
	isSet bool
}

func (v NullableEmailAttributes) Get() *EmailAttributes {
	return v.value
}

func (v *NullableEmailAttributes) Set(val *EmailAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailAttributes(val *EmailAttributes) *NullableEmailAttributes {
	return &NullableEmailAttributes{value: val, isSet: true}
}

func (v NullableEmailAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}