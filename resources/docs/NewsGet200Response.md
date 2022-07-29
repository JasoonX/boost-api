# NewsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**NewsGet200ResponseData**](NewsGet200ResponseData.md) |  | [optional] 

## Methods

### NewNewsGet200Response

`func NewNewsGet200Response() *NewsGet200Response`

NewNewsGet200Response instantiates a new NewsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsGet200ResponseWithDefaults

`func NewNewsGet200ResponseWithDefaults() *NewsGet200Response`

NewNewsGet200ResponseWithDefaults instantiates a new NewsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NewsGet200Response) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewsGet200Response) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewsGet200Response) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewsGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *NewsGet200Response) GetData() NewsGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NewsGet200Response) GetDataOk() (*NewsGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NewsGet200Response) SetData(v NewsGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *NewsGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


