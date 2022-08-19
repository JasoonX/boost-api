# NewsGet201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**News**](News.md) |  | [optional] 

## Methods

### NewNewsGet201Response

`func NewNewsGet201Response() *NewsGet201Response`

NewNewsGet201Response instantiates a new NewsGet201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsGet201ResponseWithDefaults

`func NewNewsGet201ResponseWithDefaults() *NewsGet201Response`

NewNewsGet201ResponseWithDefaults instantiates a new NewsGet201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NewsGet201Response) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewsGet201Response) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewsGet201Response) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewsGet201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *NewsGet201Response) GetData() News`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NewsGet201Response) GetDataOk() (*News, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NewsGet201Response) SetData(v News)`

SetData sets Data field to given value.

### HasData

`func (o *NewsGet201Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


