# NextAvailableFLDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **int64** | The CIDR of the __ForwardLookingDelegation__. This is required. | 
**Comment** | Pointer to **string** | The description for the __ForwardLookingDelegation__. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**Count** | Pointer to **int64** | The count of __ForwardLookingDelegation__ required. If not provided, it will default to 1. | [optional] 
**Name** | Pointer to **string** | The name to be provided. | [optional] 
**Protocol** | Pointer to **string** | The version of the address (_ip4_ or _ip6_). If not present then it will default to _ip4_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the __ForwardLookingDelegation__ in JSON format. The tag will be used to identify the federated blocks which should be considered for __ForwardLookingDelegation__. | [optional] 

## Methods

### NewNextAvailableFLDRequest

`func NewNextAvailableFLDRequest(cidr int64, ) *NextAvailableFLDRequest`

NewNextAvailableFLDRequest instantiates a new NextAvailableFLDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextAvailableFLDRequestWithDefaults

`func NewNextAvailableFLDRequestWithDefaults() *NextAvailableFLDRequest`

NewNextAvailableFLDRequestWithDefaults instantiates a new NextAvailableFLDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *NextAvailableFLDRequest) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NextAvailableFLDRequest) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NextAvailableFLDRequest) SetCidr(v int64)`

SetCidr sets Cidr field to given value.


### GetComment

`func (o *NextAvailableFLDRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NextAvailableFLDRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NextAvailableFLDRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NextAvailableFLDRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCount

`func (o *NextAvailableFLDRequest) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NextAvailableFLDRequest) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NextAvailableFLDRequest) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *NextAvailableFLDRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetName

`func (o *NextAvailableFLDRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NextAvailableFLDRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NextAvailableFLDRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NextAvailableFLDRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *NextAvailableFLDRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NextAvailableFLDRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NextAvailableFLDRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NextAvailableFLDRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTags

`func (o *NextAvailableFLDRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NextAvailableFLDRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NextAvailableFLDRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *NextAvailableFLDRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


