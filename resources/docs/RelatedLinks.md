# RelatedLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **string** | The self link of the resource | 
**Related** | **string** | The related link of the resource (might think of it as of parent link) | 

## Methods

### NewRelatedLinks

`func NewRelatedLinks(self string, related string, ) *RelatedLinks`

NewRelatedLinks instantiates a new RelatedLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedLinksWithDefaults

`func NewRelatedLinksWithDefaults() *RelatedLinks`

NewRelatedLinksWithDefaults instantiates a new RelatedLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *RelatedLinks) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *RelatedLinks) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *RelatedLinks) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetRelated

`func (o *RelatedLinks) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *RelatedLinks) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *RelatedLinks) SetRelated(v string)`

SetRelated sets Related field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


