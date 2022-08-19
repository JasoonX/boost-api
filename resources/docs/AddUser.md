# AddUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EntityType**](EntityType.md) |  | 
**Attributes** | [**AddUserAttributes**](AddUserAttributes.md) |  | 

## Methods

### NewAddUser

`func NewAddUser(type_ EntityType, attributes AddUserAttributes, ) *AddUser`

NewAddUser instantiates a new AddUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserWithDefaults

`func NewAddUserWithDefaults() *AddUser`

NewAddUserWithDefaults instantiates a new AddUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddUser) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddUser) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddUser) SetType(v EntityType)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AddUser) GetAttributes() AddUserAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AddUser) GetAttributesOk() (*AddUserAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AddUser) SetAttributes(v AddUserAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


