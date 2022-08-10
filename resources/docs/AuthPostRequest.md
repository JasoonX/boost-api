# AuthPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Credentials**](Credentials.md) |  | 

## Methods

### NewAuthPostRequest

`func NewAuthPostRequest(data Credentials, ) *AuthPostRequest`

NewAuthPostRequest instantiates a new AuthPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthPostRequestWithDefaults

`func NewAuthPostRequestWithDefaults() *AuthPostRequest`

NewAuthPostRequestWithDefaults instantiates a new AuthPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AuthPostRequest) GetData() Credentials`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthPostRequest) GetDataOk() (*Credentials, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthPostRequest) SetData(v Credentials)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


