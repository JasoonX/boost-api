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

// NewsMediaResource struct for NewsMediaResource
type NewsMediaResource struct {
	Type *string `json:"type,omitempty"`
	Url  *string `json:"url,omitempty"`
	// Any meta information
	Meta map[string]interface{} `json:"meta,omitempty"`
}

// NewNewsMediaResource instantiates a new NewsMediaResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewsMediaResource() *NewsMediaResource {
	this := NewsMediaResource{}
	return &this
}

// NewNewsMediaResourceWithDefaults instantiates a new NewsMediaResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewsMediaResourceWithDefaults() *NewsMediaResource {
	this := NewsMediaResource{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NewsMediaResource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsMediaResource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NewsMediaResource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NewsMediaResource) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NewsMediaResource) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsMediaResource) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NewsMediaResource) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NewsMediaResource) SetUrl(v string) {
	o.Url = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NewsMediaResource) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsMediaResource) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NewsMediaResource) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *NewsMediaResource) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o NewsMediaResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableNewsMediaResource struct {
	value *NewsMediaResource
	isSet bool
}

func (v NullableNewsMediaResource) Get() *NewsMediaResource {
	return v.value
}

func (v *NullableNewsMediaResource) Set(val *NewsMediaResource) {
	v.value = val
	v.isSet = true
}

func (v NullableNewsMediaResource) IsSet() bool {
	return v.isSet
}

func (v *NullableNewsMediaResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewsMediaResource(val *NewsMediaResource) *NullableNewsMediaResource {
	return &NullableNewsMediaResource{value: val, isSet: true}
}

func (v NullableNewsMediaResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewsMediaResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
