# NewsMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional]
**Text** | Pointer to **string** |  | [optional]
**Resources** | Pointer to [**[]NewsMediaResource**](NewsMediaResource.md) |  | [optional]

## Methods

### NewNewsMedia

`func NewNewsMedia() *NewsMedia`

NewNewsMedia instantiates a new NewsMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsMediaWithDefaults

`func NewNewsMediaWithDefaults() *NewsMedia`

NewNewsMediaWithDefaults instantiates a new NewsMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewsMedia) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewsMedia) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewsMedia) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NewsMedia) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetText

`func (o *NewsMedia) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NewsMedia) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NewsMedia) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *NewsMedia) HasText() bool`

HasText returns a boolean if a field has been set.

### GetResources

`func (o *NewsMedia) GetResources() []NewsMediaResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NewsMedia) GetResourcesOk() (*[]NewsMediaResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NewsMedia) SetResources(v []NewsMediaResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *NewsMedia) HasResources() bool`

HasResources returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


