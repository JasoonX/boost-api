# UsersIdGet200ResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to [**[]Email**](Email.md) | Emails that belong to this user. | [optional] 
**Phones** | Pointer to [**[]Phone**](Phone.md) | Phones that belong to this user. | [optional] 

## Methods

### NewUsersIdGet200ResponseIncluded

`func NewUsersIdGet200ResponseIncluded() *UsersIdGet200ResponseIncluded`

NewUsersIdGet200ResponseIncluded instantiates a new UsersIdGet200ResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersIdGet200ResponseIncludedWithDefaults

`func NewUsersIdGet200ResponseIncludedWithDefaults() *UsersIdGet200ResponseIncluded`

NewUsersIdGet200ResponseIncludedWithDefaults instantiates a new UsersIdGet200ResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UsersIdGet200ResponseIncluded) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UsersIdGet200ResponseIncluded) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UsersIdGet200ResponseIncluded) SetEmails(v []Email)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UsersIdGet200ResponseIncluded) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPhones

`func (o *UsersIdGet200ResponseIncluded) GetPhones() []Phone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *UsersIdGet200ResponseIncluded) GetPhonesOk() (*[]Phone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *UsersIdGet200ResponseIncluded) SetPhones(v []Phone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *UsersIdGet200ResponseIncluded) HasPhones() bool`

HasPhones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


