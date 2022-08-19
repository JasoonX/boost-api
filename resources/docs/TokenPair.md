# TokenPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EntityType**](EntityType.md) |  | 
**Attributes** | [**TokenPairAttributes**](TokenPairAttributes.md) |  | 

## Methods

### NewTokenPair

`func NewTokenPair(type_ EntityType, attributes TokenPairAttributes, ) *TokenPair`

NewTokenPair instantiates a new TokenPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPairWithDefaults

`func NewTokenPairWithDefaults() *TokenPair`

NewTokenPairWithDefaults instantiates a new TokenPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TokenPair) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenPair) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenPair) SetType(v EntityType)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TokenPair) GetAttributes() TokenPairAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TokenPair) GetAttributesOk() (*TokenPairAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TokenPair) SetAttributes(v TokenPairAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


