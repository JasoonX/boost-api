# NewsGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AddNews**](AddNews.md) |  | 

## Methods

### NewNewsGetRequest

`func NewNewsGetRequest(data AddNews, ) *NewsGetRequest`

NewNewsGetRequest instantiates a new NewsGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsGetRequestWithDefaults

`func NewNewsGetRequestWithDefaults() *NewsGetRequest`

NewNewsGetRequestWithDefaults instantiates a new NewsGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NewsGetRequest) GetData() AddNews`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NewsGetRequest) GetDataOk() (*AddNews, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NewsGetRequest) SetData(v AddNews)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


