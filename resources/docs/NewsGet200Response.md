# NewsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Page** | Pointer to [**Page**](Page.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 
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

### GetPage

`func (o *NewsGet200Response) GetPage() Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *NewsGet200Response) GetPageOk() (*Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *NewsGet200Response) SetPage(v Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *NewsGet200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *NewsGet200Response) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NewsGet200Response) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NewsGet200Response) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NewsGet200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *NewsGet200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NewsGet200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NewsGet200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NewsGet200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *NewsGet200Response) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *NewsGet200Response) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *NewsGet200Response) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *NewsGet200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

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


