# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**EntityType**](EntityType.md) |  | 
**Attributes** | [**EmailAttributes**](EmailAttributes.md) |  | 
**Relationships** | [**EmailRelationships**](EmailRelationships.md) |  | 

## Methods

### NewEmail

`func NewEmail(id string, type_ EntityType, attributes EmailAttributes, relationships EmailRelationships, ) *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Email) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Email) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Email) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Email) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Email) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Email) SetType(v EntityType)`

SetType sets Type field to given value.


### GetAttributes

`func (o *Email) GetAttributes() EmailAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Email) GetAttributesOk() (*EmailAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Email) SetAttributes(v EmailAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *Email) GetRelationships() EmailRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Email) GetRelationshipsOk() (*EmailRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Email) SetRelationships(v EmailRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


