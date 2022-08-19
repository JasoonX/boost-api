# UsersIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 
**Data** | [**User**](User.md) |  | 
**Included** | Pointer to [**UsersIdGet200ResponseIncluded**](UsersIdGet200ResponseIncluded.md) |  | [optional] 

## Methods

### NewUsersIdGet200Response

`func NewUsersIdGet200Response(data User, ) *UsersIdGet200Response`

NewUsersIdGet200Response instantiates a new UsersIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersIdGet200ResponseWithDefaults

`func NewUsersIdGet200ResponseWithDefaults() *UsersIdGet200Response`

NewUsersIdGet200ResponseWithDefaults instantiates a new UsersIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UsersIdGet200Response) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsersIdGet200Response) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsersIdGet200Response) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsersIdGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMeta

`func (o *UsersIdGet200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UsersIdGet200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UsersIdGet200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UsersIdGet200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *UsersIdGet200Response) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UsersIdGet200Response) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UsersIdGet200Response) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UsersIdGet200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *UsersIdGet200Response) GetData() User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsersIdGet200Response) GetDataOk() (*User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsersIdGet200Response) SetData(v User)`

SetData sets Data field to given value.


### GetIncluded

`func (o *UsersIdGet200Response) GetIncluded() UsersIdGet200ResponseIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *UsersIdGet200Response) GetIncludedOk() (*UsersIdGet200ResponseIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *UsersIdGet200Response) SetIncluded(v UsersIdGet200ResponseIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *UsersIdGet200Response) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


