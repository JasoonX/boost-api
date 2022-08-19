# RelatedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**EntityType**](EntityType.md) |  | [optional] 

## Methods

### NewRelatedEntity

`func NewRelatedEntity() *RelatedEntity`

NewRelatedEntity instantiates a new RelatedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedEntityWithDefaults

`func NewRelatedEntityWithDefaults() *RelatedEntity`

NewRelatedEntityWithDefaults instantiates a new RelatedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelatedEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelatedEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RelatedEntity) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelatedEntity) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelatedEntity) SetType(v EntityType)`

SetType sets Type field to given value.

### HasType

`func (o *RelatedEntity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


