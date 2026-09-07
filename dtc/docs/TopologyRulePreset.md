# TopologyRulePreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Required. Display name of __TopologyRulePreset__. Must be unique within __Topology__. | 
**Source** | Pointer to **string** | Type of source.  Allowed values: - subnet - tags - default  Defaults to _default_. | [optional] 
**Subnets** | Pointer to **[]string** | Optional. List of subnets in CIDR format.  Must be set if _source_ is _subnet_, otherwise must be empty. | [optional] 
**Tags** | Pointer to [**[]TagRule**](TagRule.md) | Optional. List of tag rules to match against a source object&#39;s effective tags. Effective tags &#x3D; direct tags plus tags inherited from the IPAM parent chain (IPSpace → Address Block → Subnet); the closer level wins on key conflicts. All rules use AND semantics: an object must satisfy every __TagRule__ to match.  Must be set if _source_ is _tags_, otherwise must be empty. | [optional] 

## Methods

### NewTopologyRulePreset

`func NewTopologyRulePreset(name string, ) *TopologyRulePreset`

NewTopologyRulePreset instantiates a new TopologyRulePreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyRulePresetWithDefaults

`func NewTopologyRulePresetWithDefaults() *TopologyRulePreset`

NewTopologyRulePresetWithDefaults instantiates a new TopologyRulePreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TopologyRulePreset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyRulePreset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyRulePreset) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *TopologyRulePreset) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TopologyRulePreset) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TopologyRulePreset) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TopologyRulePreset) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubnets

`func (o *TopologyRulePreset) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *TopologyRulePreset) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *TopologyRulePreset) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *TopologyRulePreset) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetTags

`func (o *TopologyRulePreset) GetTags() []TagRule`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TopologyRulePreset) GetTagsOk() (*[]TagRule, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TopologyRulePreset) SetTags(v []TagRule)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TopologyRulePreset) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


