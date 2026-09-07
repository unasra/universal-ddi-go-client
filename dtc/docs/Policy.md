# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for __Policy__. | [optional] 
**Disabled** | Pointer to **bool** | Optional. Flag which enables/disables __Policy__.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**TTLInheritance**](TTLInheritance.md) | Optional. The inheritance configuration. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | Output only. __Policy__ metadata. Defaults to empty object and should be explicitly requested using field selection. | [optional] 
**Method** | **string** | Load balancing method used for selecting __Pool__ assigned to __Policy__.  Valid values are: * _round_robin_ If the _round_robin_ load balancing method is selected, Universal DDI adjusts the response to a query in a sequential and circular manner, directing clients to pools.  * _ratio_ If _ratio_ load balancing method is selected, Universal DDI adjusts the response to a query so that clients are directed to pool using weighted round robin, a load-balancing pattern in which requests are distributed among several resources based on weight assigned to each resource. The distribution of responses over time will be equal for all available pools but the sequence of the responses won&#39;t be guaranteed. When equal weights are assigned for resources (pools) it effectively leads to basic round robin configuration which directs clients to pools in a sequential and circular manner.  * _topology_ If _topology_ load balancing method is selected the pools configured for the policy are ignored and topology rules are used instead.  * _global_availability_ If _global_availability_ load balancing method is selected clients are directed to the first pool that is up in the _pools_ list.  Defaults to _round_robin_. | 
**Name** | **string** | Display name of __Policy__. | 
**Pools** | Pointer to [**[]PolicyPool**](PolicyPool.md) | Optional. List of __Pool__ objects assigned to __Policy__.  Defaults to _empty_. | [optional] 
**Rules** | Pointer to [**[]TopologyRule**](TopologyRule.md) | Optional. List of inline __TopologyRule__ objects defining the resolving strategy for __Policy__.  Defaults to a list of single, default __TopologyRule__. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __Policy__ in JSON format. | [optional] 
**Ttl** | Pointer to **int64** | Optional. Time to live value (in seconds) to be used for records in DTC response. Unsigned integer, min: 0, max 2147483647 (31-bits per RFC-2181). | [optional] 

## Methods

### NewPolicy

`func NewPolicy(method string, name string, ) *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *Policy) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Policy) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Policy) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Policy) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *Policy) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Policy) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Policy) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Policy) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *Policy) GetInheritanceSources() TTLInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *Policy) GetInheritanceSourcesOk() (*TTLInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *Policy) SetInheritanceSources(v TTLInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *Policy) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetMetadata

`func (o *Policy) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Policy) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Policy) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Policy) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMethod

`func (o *Policy) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Policy) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Policy) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.


### GetPools

`func (o *Policy) GetPools() []PolicyPool`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *Policy) GetPoolsOk() (*[]PolicyPool, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *Policy) SetPools(v []PolicyPool)`

SetPools sets Pools field to given value.

### HasPools

`func (o *Policy) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetRules

`func (o *Policy) GetRules() []TopologyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Policy) GetRulesOk() (*[]TopologyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Policy) SetRules(v []TopologyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Policy) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTags

`func (o *Policy) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Policy) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Policy) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Policy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTtl

`func (o *Policy) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Policy) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Policy) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Policy) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


