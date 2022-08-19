# UserAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Status** | [**UserStatus**](UserStatus.md) |  | 
**Role** | [**[]UserRole**](UserRole.md) |  | 

## Methods

### NewUserAttributes

`func NewUserAttributes(createdAt time.Time, status UserStatus, role []UserRole, ) *UserAttributes`

NewUserAttributes instantiates a new UserAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAttributesWithDefaults

`func NewUserAttributesWithDefaults() *UserAttributes`

NewUserAttributesWithDefaults instantiates a new UserAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UserAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *UserAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *UserAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetStatus

`func (o *UserAttributes) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserAttributes) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserAttributes) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *UserAttributes) GetRole() []UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserAttributes) GetRoleOk() (*[]UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserAttributes) SetRole(v []UserRole)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


