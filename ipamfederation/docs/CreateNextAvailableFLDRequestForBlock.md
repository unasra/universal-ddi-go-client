# CreateNextAvailableFLDRequestForBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **int64** | The CIDR of the __ForwardLookingDelegation__. This is required. | 
**Comment** | Pointer to **string** | The description for the __ForwardLookingDelegation__. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**Count** | Pointer to **int64** | The count of __ForwardLookingDelegation__ required. If not provided, it will default to 1. | [optional] 
**Id** | **string** | The resource identifier. | [readonly] 
**Name** | Pointer to **string** | The name to be provided. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the __ForwardLookingDelegation__ in JSON format. | [optional] 

## Methods

### NewCreateNextAvailableFLDRequestForBlock

`func NewCreateNextAvailableFLDRequestForBlock(cidr int64, id string, ) *CreateNextAvailableFLDRequestForBlock`

NewCreateNextAvailableFLDRequestForBlock instantiates a new CreateNextAvailableFLDRequestForBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNextAvailableFLDRequestForBlockWithDefaults

`func NewCreateNextAvailableFLDRequestForBlockWithDefaults() *CreateNextAvailableFLDRequestForBlock`

NewCreateNextAvailableFLDRequestForBlockWithDefaults instantiates a new CreateNextAvailableFLDRequestForBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *CreateNextAvailableFLDRequestForBlock) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateNextAvailableFLDRequestForBlock) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateNextAvailableFLDRequestForBlock) SetCidr(v int64)`

SetCidr sets Cidr field to given value.


### GetComment

`func (o *CreateNextAvailableFLDRequestForBlock) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateNextAvailableFLDRequestForBlock) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateNextAvailableFLDRequestForBlock) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateNextAvailableFLDRequestForBlock) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCount

`func (o *CreateNextAvailableFLDRequestForBlock) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateNextAvailableFLDRequestForBlock) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateNextAvailableFLDRequestForBlock) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CreateNextAvailableFLDRequestForBlock) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetId

`func (o *CreateNextAvailableFLDRequestForBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNextAvailableFLDRequestForBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNextAvailableFLDRequestForBlock) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateNextAvailableFLDRequestForBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNextAvailableFLDRequestForBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNextAvailableFLDRequestForBlock) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNextAvailableFLDRequestForBlock) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *CreateNextAvailableFLDRequestForBlock) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateNextAvailableFLDRequestForBlock) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateNextAvailableFLDRequestForBlock) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateNextAvailableFLDRequestForBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


