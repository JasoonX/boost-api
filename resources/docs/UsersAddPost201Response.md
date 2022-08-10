# UsersAddPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 
**Data** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewUsersAddPost201Response

`func NewUsersAddPost201Response() *UsersAddPost201Response`

NewUsersAddPost201Response instantiates a new UsersAddPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersAddPost201ResponseWithDefaults

`func NewUsersAddPost201ResponseWithDefaults() *UsersAddPost201Response`

NewUsersAddPost201ResponseWithDefaults instantiates a new UsersAddPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UsersAddPost201Response) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsersAddPost201Response) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsersAddPost201Response) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsersAddPost201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *UsersAddPost201Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UsersAddPost201Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UsersAddPost201Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UsersAddPost201Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *UsersAddPost201Response) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UsersAddPost201Response) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UsersAddPost201Response) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UsersAddPost201Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *UsersAddPost201Response) GetData() User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsersAddPost201Response) GetDataOk() (*User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsersAddPost201Response) SetData(v User)`

SetData sets Data field to given value.

### HasData

`func (o *UsersAddPost201Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


