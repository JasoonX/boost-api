# NewsAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Media** | Pointer to [**NewsMedia**](NewsMedia.md) |  | [optional] 

## Methods

### NewNewsAttributes

`func NewNewsAttributes(createdAt time.Time, ) *NewsAttributes`

NewNewsAttributes instantiates a new NewsAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsAttributesWithDefaults

`func NewNewsAttributesWithDefaults() *NewsAttributes`

NewNewsAttributesWithDefaults instantiates a new NewsAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NewsAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NewsAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NewsAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NewsAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NewsAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NewsAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NewsAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMedia

`func (o *NewsAttributes) GetMedia() NewsMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *NewsAttributes) GetMediaOk() (*NewsMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *NewsAttributes) SetMedia(v NewsMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *NewsAttributes) HasMedia() bool`

HasMedia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


