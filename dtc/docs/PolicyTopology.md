# PolicyTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]PolicyTopologyRuleBinding**](PolicyTopologyRuleBinding.md) | Ordered list of destination bindings, one per __TopologyRulePreset__ in the referenced __Topology__. Each binding maps a preset _name_ to the destination to use when that preset matches. | [optional] 
**TopologyId** | **string** | The resource identifier. | 

## Methods

### NewPolicyTopology

`func NewPolicyTopology(topologyId string, ) *PolicyTopology`

NewPolicyTopology instantiates a new PolicyTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTopologyWithDefaults

`func NewPolicyTopologyWithDefaults() *PolicyTopology`

NewPolicyTopologyWithDefaults instantiates a new PolicyTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *PolicyTopology) GetRules() []PolicyTopologyRuleBinding`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PolicyTopology) GetRulesOk() (*[]PolicyTopologyRuleBinding, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PolicyTopology) SetRules(v []PolicyTopologyRuleBinding)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PolicyTopology) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTopologyId

`func (o *PolicyTopology) GetTopologyId() string`

GetTopologyId returns the TopologyId field if non-nil, zero value otherwise.

### GetTopologyIdOk

`func (o *PolicyTopology) GetTopologyIdOk() (*string, bool)`

GetTopologyIdOk returns a tuple with the TopologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyId

`func (o *PolicyTopology) SetTopologyId(v string)`

SetTopologyId sets TopologyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


