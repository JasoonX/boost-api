# NewsAddPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**News**](News.md) |  | [optional] 

## Methods

### NewNewsAddPost201Response

`func NewNewsAddPost201Response() *NewsAddPost201Response`

NewNewsAddPost201Response instantiates a new NewsAddPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsAddPost201ResponseWithDefaults

`func NewNewsAddPost201ResponseWithDefaults() *NewsAddPost201Response`

NewNewsAddPost201ResponseWithDefaults instantiates a new NewsAddPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NewsAddPost201Response) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewsAddPost201Response) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewsAddPost201Response) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewsAddPost201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *NewsAddPost201Response) GetData() News`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NewsAddPost201Response) GetDataOk() (*News, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NewsAddPost201Response) SetData(v News)`

SetData sets Data field to given value.

### HasData

`func (o *NewsAddPost201Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


