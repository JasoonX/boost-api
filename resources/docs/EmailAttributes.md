# EmailAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**IsVerified** | Pointer to **bool** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 

## Methods

### NewEmailAttributes

`func NewEmailAttributes(email string, ) *EmailAttributes`

NewEmailAttributes instantiates a new EmailAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAttributesWithDefaults

`func NewEmailAttributesWithDefaults() *EmailAttributes`

NewEmailAttributesWithDefaults instantiates a new EmailAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsVerified

`func (o *EmailAttributes) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *EmailAttributes) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *EmailAttributes) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *EmailAttributes) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetIsPrimary

`func (o *EmailAttributes) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *EmailAttributes) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *EmailAttributes) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *EmailAttributes) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


