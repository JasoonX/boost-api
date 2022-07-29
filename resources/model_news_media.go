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

// NewsMedia struct for NewsMedia
type NewsMedia struct {
	Title     *string             `json:"title,omitempty"`
	Text      *string             `json:"text,omitempty"`
	Resources []NewsMediaResource `json:"resources,omitempty"`
}

// NewNewsMedia instantiates a new NewsMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewsMedia() *NewsMedia {
	this := NewsMedia{}
	return &this
}

// NewNewsMediaWithDefaults instantiates a new NewsMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewsMediaWithDefaults() *NewsMedia {
	this := NewsMedia{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *NewsMedia) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsMedia) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *NewsMedia) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *NewsMedia) SetTitle(v string) {
	o.Title = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *NewsMedia) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsMedia) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *NewsMedia) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *NewsMedia) SetText(v string) {
	o.Text = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *NewsMedia) GetResources() []NewsMediaResource {
	if o == nil || o.Resources == nil {
		var ret []NewsMediaResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsMedia) GetResourcesOk() ([]NewsMediaResource, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *NewsMedia) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []NewsMediaResource and assigns it to the Resources field.
func (o *NewsMedia) SetResources(v []NewsMediaResource) {
	o.Resources = v
}

func (o NewsMedia) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableNewsMedia struct {
	value *NewsMedia
	isSet bool
}

func (v NullableNewsMedia) Get() *NewsMedia {
	return v.value
}

func (v *NullableNewsMedia) Set(val *NewsMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableNewsMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableNewsMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewsMedia(val *NewsMedia) *NullableNewsMedia {
	return &NullableNewsMedia{value: val, isSet: true}
}

func (v NullableNewsMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewsMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}