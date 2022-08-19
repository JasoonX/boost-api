# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int64** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Orderby** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewPage

`func NewPage() *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *Page) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Page) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Page) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Page) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *Page) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Page) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Page) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Page) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderby

`func (o *Page) GetOrderby() string`

GetOrderby returns the Orderby field if non-nil, zero value otherwise.

### GetOrderbyOk

`func (o *Page) GetOrderbyOk() (*string, bool)`

GetOrderbyOk returns a tuple with the Orderby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderby

`func (o *Page) SetOrderby(v string)`

SetOrderby sets Orderby field to given value.

### HasOrderby

`func (o *Page) HasOrderby() bool`

HasOrderby returns a boolean if a field has been set.

### GetOrder

`func (o *Page) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Page) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Page) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Page) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTotal

`func (o *Page) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Page) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Page) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Page) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


