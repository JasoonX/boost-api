# AddEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EntityType**](EntityType.md) |  | 
**Attributes** | [**AddEmailAttributes**](AddEmailAttributes.md) |  | 
**Relationships** | [**AddEmailRelationships**](AddEmailRelationships.md) |  | 

## Methods

### NewAddEmail

`func NewAddEmail(type_ EntityType, attributes AddEmailAttributes, relationships AddEmailRelationships, ) *AddEmail`

NewAddEmail instantiates a new AddEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEmailWithDefaults

`func NewAddEmailWithDefaults() *AddEmail`

NewAddEmailWithDefaults instantiates a new AddEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddEmail) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddEmail) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddEmail) SetType(v EntityType)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AddEmail) GetAttributes() AddEmailAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AddEmail) GetAttributesOk() (*AddEmailAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AddEmail) SetAttributes(v AddEmailAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AddEmail) GetRelationships() AddEmailRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AddEmail) GetRelationshipsOk() (*AddEmailRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AddEmail) SetRelationships(v AddEmailRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


