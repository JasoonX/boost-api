# AuthPost400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**AuthPost400ResponseError**](AuthPost400ResponseError.md) |  | [optional] 

## Methods

### NewAuthPost400Response

`func NewAuthPost400Response() *AuthPost400Response`

NewAuthPost400Response instantiates a new AuthPost400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthPost400ResponseWithDefaults

`func NewAuthPost400ResponseWithDefaults() *AuthPost400Response`

NewAuthPost400ResponseWithDefaults instantiates a new AuthPost400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AuthPost400Response) GetError() AuthPost400ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthPost400Response) GetErrorOk() (*AuthPost400ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthPost400Response) SetError(v AuthPost400ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *AuthPost400Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


