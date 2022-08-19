# AddPhoneAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberNumber** | Pointer to **string** | Phone number of the subscriber. | [optional] 
**CountryCode** | Pointer to **string** | Country code of the subscriber. | [optional] 
**Extension** | Pointer to **string** | Extension of the subscriber. | [optional] 

## Methods

### NewAddPhoneAttributes

`func NewAddPhoneAttributes() *AddPhoneAttributes`

NewAddPhoneAttributes instantiates a new AddPhoneAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPhoneAttributesWithDefaults

`func NewAddPhoneAttributesWithDefaults() *AddPhoneAttributes`

NewAddPhoneAttributesWithDefaults instantiates a new AddPhoneAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberNumber

`func (o *AddPhoneAttributes) GetSubscriberNumber() string`

GetSubscriberNumber returns the SubscriberNumber field if non-nil, zero value otherwise.

### GetSubscriberNumberOk

`func (o *AddPhoneAttributes) GetSubscriberNumberOk() (*string, bool)`

GetSubscriberNumberOk returns a tuple with the SubscriberNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberNumber

`func (o *AddPhoneAttributes) SetSubscriberNumber(v string)`

SetSubscriberNumber sets SubscriberNumber field to given value.

### HasSubscriberNumber

`func (o *AddPhoneAttributes) HasSubscriberNumber() bool`

HasSubscriberNumber returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddPhoneAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddPhoneAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddPhoneAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AddPhoneAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetExtension

`func (o *AddPhoneAttributes) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *AddPhoneAttributes) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *AddPhoneAttributes) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *AddPhoneAttributes) HasExtension() bool`

HasExtension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


