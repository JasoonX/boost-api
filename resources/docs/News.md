# News

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**EntityType**](EntityType.md) |  | 
**Attributes** | [**NewsAttributes**](NewsAttributes.md) |  | 
**Relationships** | [**NewsRelationships**](NewsRelationships.md) |  | 

## Methods

### NewNews

`func NewNews(id string, type_ EntityType, attributes NewsAttributes, relationships NewsRelationships, ) *News`

NewNews instantiates a new News object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsWithDefaults

`func NewNewsWithDefaults() *News`

NewNewsWithDefaults instantiates a new News object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *News) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *News) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *News) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *News) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *News) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *News) SetType(v EntityType)`

SetType sets Type field to given value.


### GetAttributes

`func (o *News) GetAttributes() NewsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *News) GetAttributesOk() (*NewsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *News) SetAttributes(v NewsAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *News) GetRelationships() NewsRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *News) GetRelationshipsOk() (*NewsRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *News) SetRelationships(v NewsRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


