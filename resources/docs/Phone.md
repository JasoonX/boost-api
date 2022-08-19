# Phone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**EntityType**](EntityType.md) |  | 
**Attributes** | [**PhoneAttributes**](PhoneAttributes.md) |  | 
**Relationships** | [**PhoneRelationships**](PhoneRelationships.md) |  | 

## Methods

### NewPhone

`func NewPhone(id string, type_ EntityType, attributes PhoneAttributes, relationships PhoneRelationships, ) *Phone`

NewPhone instantiates a new Phone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneWithDefaults

`func NewPhoneWithDefaults() *Phone`

NewPhoneWithDefaults instantiates a new Phone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Phone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Phone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Phone) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Phone) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Phone) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Phone) SetType(v EntityType)`

SetType sets Type field to given value.


### GetAttributes

`func (o *Phone) GetAttributes() PhoneAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Phone) GetAttributesOk() (*PhoneAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Phone) SetAttributes(v PhoneAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *Phone) GetRelationships() PhoneRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Phone) GetRelationshipsOk() (*PhoneRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Phone) SetRelationships(v PhoneRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


