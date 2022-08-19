# Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**EntityType**](EntityType.md) |  | 
**Attributes** | [**CredentialsAttributes**](CredentialsAttributes.md) |  | 

## Methods

### NewCredentials

`func NewCredentials(id string, type_ EntityType, attributes CredentialsAttributes, ) *Credentials`

NewCredentials instantiates a new Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsWithDefaults

`func NewCredentialsWithDefaults() *Credentials`

NewCredentialsWithDefaults instantiates a new Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Credentials) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Credentials) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Credentials) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Credentials) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Credentials) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Credentials) SetType(v EntityType)`

SetType sets Type field to given value.


### GetAttributes

`func (o *Credentials) GetAttributes() CredentialsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Credentials) GetAttributesOk() (*CredentialsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Credentials) SetAttributes(v CredentialsAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


