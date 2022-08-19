# PhoneAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberNumber** | **string** |  | 
**CountryCode** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**PrimaryFlag** | Pointer to **bool** |  | [optional] 

## Methods

### NewPhoneAttributes

`func NewPhoneAttributes(subscriberNumber string, ) *PhoneAttributes`

NewPhoneAttributes instantiates a new PhoneAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneAttributesWithDefaults

`func NewPhoneAttributesWithDefaults() *PhoneAttributes`

NewPhoneAttributesWithDefaults instantiates a new PhoneAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberNumber

`func (o *PhoneAttributes) GetSubscriberNumber() string`

GetSubscriberNumber returns the SubscriberNumber field if non-nil, zero value otherwise.

### GetSubscriberNumberOk

`func (o *PhoneAttributes) GetSubscriberNumberOk() (*string, bool)`

GetSubscriberNumberOk returns a tuple with the SubscriberNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberNumber

`func (o *PhoneAttributes) SetSubscriberNumber(v string)`

SetSubscriberNumber sets SubscriberNumber field to given value.


### GetCountryCode

`func (o *PhoneAttributes) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PhoneAttributes) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PhoneAttributes) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PhoneAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetExtension

`func (o *PhoneAttributes) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *PhoneAttributes) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *PhoneAttributes) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *PhoneAttributes) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetVerified

`func (o *PhoneAttributes) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PhoneAttributes) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PhoneAttributes) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *PhoneAttributes) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetPrimaryFlag

`func (o *PhoneAttributes) GetPrimaryFlag() bool`

GetPrimaryFlag returns the PrimaryFlag field if non-nil, zero value otherwise.

### GetPrimaryFlagOk

`func (o *PhoneAttributes) GetPrimaryFlagOk() (*bool, bool)`

GetPrimaryFlagOk returns a tuple with the PrimaryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFlag

`func (o *PhoneAttributes) SetPrimaryFlag(v bool)`

SetPrimaryFlag sets PrimaryFlag field to given value.

### HasPrimaryFlag

`func (o *PhoneAttributes) HasPrimaryFlag() bool`

HasPrimaryFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


