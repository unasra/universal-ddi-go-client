# TopologySource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Required. Display name of __TopologySource__. | 
**Source** | **string** | Type of source.  Allowed values: - subnet - tag_rule  Required. | 
**Subnets** | Pointer to **[]string** | Optional. List of subnets in CIDR format.  Must be set if _source_ is set to _subnet_, otherwise must be empty. | [optional] 
**TagRules** | Pointer to [**[]TagRule**](TagRule.md) | Optional. List of tag rules to match against infrastructure source objects effective tags.  Must be set if _source_ is set to _tag_rule_, otherwise must be empty. | [optional] 

## Methods

### NewTopologySource

`func NewTopologySource(name string, source string, ) *TopologySource`

NewTopologySource instantiates a new TopologySource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologySourceWithDefaults

`func NewTopologySourceWithDefaults() *TopologySource`

NewTopologySourceWithDefaults instantiates a new TopologySource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TopologySource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologySource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologySource) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *TopologySource) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TopologySource) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TopologySource) SetSource(v string)`

SetSource sets Source field to given value.


### GetSubnets

`func (o *TopologySource) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *TopologySource) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *TopologySource) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *TopologySource) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetTagRules

`func (o *TopologySource) GetTagRules() []TagRule`

GetTagRules returns the TagRules field if non-nil, zero value otherwise.

### GetTagRulesOk

`func (o *TopologySource) GetTagRulesOk() (*[]TagRule, bool)`

GetTagRulesOk returns a tuple with the TagRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagRules

`func (o *TopologySource) SetTagRules(v []TagRule)`

SetTagRules sets TagRules field to given value.

### HasTagRules

`func (o *TopologySource) HasTagRules() bool`

HasTagRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


