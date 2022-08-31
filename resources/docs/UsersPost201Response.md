# UsersPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 
**Data** | Pointer to [**UsersPost201ResponseData**](UsersPost201ResponseData.md) |  | [optional] 

## Methods

### NewUsersPost201Response

`func NewUsersPost201Response() *UsersPost201Response`

NewUsersPost201Response instantiates a new UsersPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersPost201ResponseWithDefaults

`func NewUsersPost201ResponseWithDefaults() *UsersPost201Response`

NewUsersPost201ResponseWithDefaults instantiates a new UsersPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UsersPost201Response) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsersPost201Response) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsersPost201Response) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsersPost201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *UsersPost201Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UsersPost201Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UsersPost201Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UsersPost201Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *UsersPost201Response) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UsersPost201Response) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UsersPost201Response) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UsersPost201Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *UsersPost201Response) GetData() UsersPost201ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsersPost201Response) GetDataOk() (*UsersPost201ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsersPost201Response) SetData(v UsersPost201ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UsersPost201Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


