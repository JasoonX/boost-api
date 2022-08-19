# UserRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | [**UserRelationshipsEmails**](UserRelationshipsEmails.md) |  | 
**Phones** | Pointer to [**UserRelationshipsPhones**](UserRelationshipsPhones.md) |  | [optional] 

## Methods

### NewUserRelationships

`func NewUserRelationships(emails UserRelationshipsEmails, ) *UserRelationships`

NewUserRelationships instantiates a new UserRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRelationshipsWithDefaults

`func NewUserRelationshipsWithDefaults() *UserRelationships`

NewUserRelationshipsWithDefaults instantiates a new UserRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UserRelationships) GetEmails() UserRelationshipsEmails`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UserRelationships) GetEmailsOk() (*UserRelationshipsEmails, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UserRelationships) SetEmails(v UserRelationshipsEmails)`

SetEmails sets Emails field to given value.


### GetPhones

`func (o *UserRelationships) GetPhones() UserRelationshipsPhones`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *UserRelationships) GetPhonesOk() (*UserRelationshipsPhones, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *UserRelationships) SetPhones(v UserRelationshipsPhones)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *UserRelationships) HasPhones() bool`

HasPhones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


