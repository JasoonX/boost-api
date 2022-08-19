# AddUserAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Phone** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | **string** | Password must be at least 8 characters long, contain at least 2 numbers and 2 uppercase letters, and 1 special symbol. | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 

## Methods

### NewAddUserAttributes

`func NewAddUserAttributes(email string, password string, ) *AddUserAttributes`

NewAddUserAttributes instantiates a new AddUserAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserAttributesWithDefaults

`func NewAddUserAttributesWithDefaults() *AddUserAttributes`

NewAddUserAttributesWithDefaults instantiates a new AddUserAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AddUserAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddUserAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddUserAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *AddUserAttributes) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddUserAttributes) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddUserAttributes) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddUserAttributes) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetUsername

`func (o *AddUserAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddUserAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddUserAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddUserAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AddUserAttributes) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddUserAttributes) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddUserAttributes) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetFirstName

`func (o *AddUserAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddUserAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddUserAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AddUserAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AddUserAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddUserAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddUserAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AddUserAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


