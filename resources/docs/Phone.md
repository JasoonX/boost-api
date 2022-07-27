# Phone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**SubscriberNumber** | **string** |  |
**CountryCode** | Pointer to **string** |  | [optional]
**Extension** | Pointer to **string** |  | [optional]
**Verified** | Pointer to **bool** |  | [optional]
**PrimaryFlag** | Pointer to **bool** |  | [optional]
**UserId** | **string** |  |

## Methods

### NewPhone

`func NewPhone(id string, subscriberNumber string, userId string, ) *Phone`

NewPhone instantiates a new Phone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneWithDefaults

`func NewPhoneWithDefaults() *Phone`

NewPhoneWithDefaults instantiates a new Phone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Phone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Phone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Phone) SetId(v string)`

SetId sets Id field to given value.

### GetSubscriberNumber

`func (o *Phone) GetSubscriberNumber() string`

GetSubscriberNumber returns the SubscriberNumber field if non-nil, zero value otherwise.

### GetSubscriberNumberOk

`func (o *Phone) GetSubscriberNumberOk() (*string, bool)`

GetSubscriberNumberOk returns a tuple with the SubscriberNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberNumber

`func (o *Phone) SetSubscriberNumber(v string)`

SetSubscriberNumber sets SubscriberNumber field to given value.

### GetCountryCode

`func (o *Phone) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Phone) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Phone) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Phone) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetExtension

`func (o *Phone) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Phone) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Phone) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Phone) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetVerified

`func (o *Phone) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Phone) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Phone) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *Phone) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetPrimaryFlag

`func (o *Phone) GetPrimaryFlag() bool`

GetPrimaryFlag returns the PrimaryFlag field if non-nil, zero value otherwise.

### GetPrimaryFlagOk

`func (o *Phone) GetPrimaryFlagOk() (*bool, bool)`

GetPrimaryFlagOk returns a tuple with the PrimaryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFlag

`func (o *Phone) SetPrimaryFlag(v bool)`

SetPrimaryFlag sets PrimaryFlag field to given value.

### HasPrimaryFlag

`func (o *Phone) HasPrimaryFlag() bool`

HasPrimaryFlag returns a boolean if a field has been set.

### GetUserId

`func (o *Phone) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Phone) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Phone) SetUserId(v string)`

SetUserId sets UserId field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


