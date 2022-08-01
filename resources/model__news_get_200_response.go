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

// NewsGet200Response struct for NewsGet200Response
type NewsGet200Response struct {
	Status *Status                 `json:"status,omitempty"`
	Page   *Page                   `json:"page,omitempty"`
	Links  *Links                  `json:"links,omitempty"`
	Meta   map[string]interface{}  `json:"meta,omitempty"`
	Errors []Error                 `json:"errors,omitempty"`
	Data   *NewsGet200ResponseData `json:"data,omitempty"`
}

// NewNewsGet200Response instantiates a new NewsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewsGet200Response() *NewsGet200Response {
	this := NewsGet200Response{}
	return &this
}

// NewNewsGet200ResponseWithDefaults instantiates a new NewsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewsGet200ResponseWithDefaults() *NewsGet200Response {
	this := NewsGet200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NewsGet200Response) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsGet200Response) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NewsGet200Response) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *NewsGet200Response) SetStatus(v Status) {
	o.Status = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *NewsGet200Response) GetPage() Page {
	if o == nil || o.Page == nil {
		var ret Page
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsGet200Response) GetPageOk() (*Page, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *NewsGet200Response) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given Page and assigns it to the Page field.
func (o *NewsGet200Response) SetPage(v Page) {
	o.Page = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NewsGet200Response) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsGet200Response) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NewsGet200Response) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *NewsGet200Response) SetLinks(v Links) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NewsGet200Response) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsGet200Response) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NewsGet200Response) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *NewsGet200Response) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *NewsGet200Response) GetErrors() []Error {
	if o == nil || o.Errors == nil {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsGet200Response) GetErrorsOk() ([]Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *NewsGet200Response) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *NewsGet200Response) SetErrors(v []Error) {
	o.Errors = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NewsGet200Response) GetData() NewsGet200ResponseData {
	if o == nil || o.Data == nil {
		var ret NewsGet200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewsGet200Response) GetDataOk() (*NewsGet200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NewsGet200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given NewsGet200ResponseData and assigns it to the Data field.
func (o *NewsGet200Response) SetData(v NewsGet200ResponseData) {
	o.Data = &v
}

func (o NewsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableNewsGet200Response struct {
	value *NewsGet200Response
	isSet bool
}

func (v NullableNewsGet200Response) Get() *NewsGet200Response {
	return v.value
}

func (v *NullableNewsGet200Response) Set(val *NewsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableNewsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableNewsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewsGet200Response(val *NewsGet200Response) *NullableNewsGet200Response {
	return &NullableNewsGet200Response{value: val, isSet: true}
}

func (v NullableNewsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
