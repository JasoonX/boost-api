# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**Verified** | Pointer to **bool** |  | [optional] 
**PrimaryFlag** | Pointer to **bool** |  | [optional] 
**UserId** | **string** |  | 

## Methods

### NewEmail

`func NewEmail(id string, email string, userId string, ) *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Email) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Email) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Email) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *Email) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Email) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Email) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetVerified

`func (o *Email) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Email) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Email) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *Email) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetPrimaryFlag

`func (o *Email) GetPrimaryFlag() bool`

GetPrimaryFlag returns the PrimaryFlag field if non-nil, zero value otherwise.

### GetPrimaryFlagOk

`func (o *Email) GetPrimaryFlagOk() (*bool, bool)`

GetPrimaryFlagOk returns a tuple with the PrimaryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFlag

`func (o *Email) SetPrimaryFlag(v bool)`

SetPrimaryFlag sets PrimaryFlag field to given value.

### HasPrimaryFlag

`func (o *Email) HasPrimaryFlag() bool`

HasPrimaryFlag returns a boolean if a field has been set.

### GetUserId

`func (o *Email) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Email) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Email) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


