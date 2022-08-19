# AddNews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EntityType**](EntityType.md) |  | 
**Attributes** | [**AddNewsAttributes**](AddNewsAttributes.md) |  | 
**Relationships** | [**AddNewsRelationships**](AddNewsRelationships.md) |  | 

## Methods

### NewAddNews

`func NewAddNews(type_ EntityType, attributes AddNewsAttributes, relationships AddNewsRelationships, ) *AddNews`

NewAddNews instantiates a new AddNews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNewsWithDefaults

`func NewAddNewsWithDefaults() *AddNews`

NewAddNewsWithDefaults instantiates a new AddNews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddNews) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddNews) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddNews) SetType(v EntityType)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AddNews) GetAttributes() AddNewsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AddNews) GetAttributesOk() (*AddNewsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AddNews) SetAttributes(v AddNewsAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AddNews) GetRelationships() AddNewsRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AddNews) GetRelationshipsOk() (*AddNewsRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AddNews) SetRelationships(v AddNewsRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


