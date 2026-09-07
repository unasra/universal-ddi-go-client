# NextAvailableFLDPoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **int64** | The CIDR prefix length of the __ForwardLookingDelegation__ to allocate. Required. | 
**Comment** | Pointer to **string** | The description for the allocated __ForwardLookingDelegation__ objects. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**Count** | Pointer to **int64** | The number of __ForwardLookingDelegation__ objects to allocate. Defaults to 1 if not provided. Maximum is 20. | [optional] 
**Id** | **string** | The resource identifier. | [readonly] 
**Name** | Pointer to **string** | The name to assign to the allocated __ForwardLookingDelegation__ objects. | [optional] 
**Protocol** | Pointer to **string** | The IP version of the __ForwardLookingDelegation__ (_ip4_ or _ip6_). Defaults to _ip4_ if not provided. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the allocated __ForwardLookingDelegation__ objects in JSON format. | [optional] 

## Methods

### NewNextAvailableFLDPoolRequest

`func NewNextAvailableFLDPoolRequest(cidr int64, id string, ) *NextAvailableFLDPoolRequest`

NewNextAvailableFLDPoolRequest instantiates a new NextAvailableFLDPoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextAvailableFLDPoolRequestWithDefaults

`func NewNextAvailableFLDPoolRequestWithDefaults() *NextAvailableFLDPoolRequest`

NewNextAvailableFLDPoolRequestWithDefaults instantiates a new NextAvailableFLDPoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *NextAvailableFLDPoolRequest) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NextAvailableFLDPoolRequest) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NextAvailableFLDPoolRequest) SetCidr(v int64)`

SetCidr sets Cidr field to given value.


### GetComment

`func (o *NextAvailableFLDPoolRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NextAvailableFLDPoolRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NextAvailableFLDPoolRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NextAvailableFLDPoolRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCount

`func (o *NextAvailableFLDPoolRequest) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NextAvailableFLDPoolRequest) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NextAvailableFLDPoolRequest) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *NextAvailableFLDPoolRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetId

`func (o *NextAvailableFLDPoolRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NextAvailableFLDPoolRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NextAvailableFLDPoolRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NextAvailableFLDPoolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NextAvailableFLDPoolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NextAvailableFLDPoolRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NextAvailableFLDPoolRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *NextAvailableFLDPoolRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NextAvailableFLDPoolRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NextAvailableFLDPoolRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NextAvailableFLDPoolRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTags

`func (o *NextAvailableFLDPoolRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NextAvailableFLDPoolRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NextAvailableFLDPoolRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *NextAvailableFLDPoolRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


