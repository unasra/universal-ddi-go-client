# PolicyTopologyRuleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Optional. DNS code to return if the referenced preset matches. Must be set if _destination_ is _code_.  Allowed values: - nodata - nxdomain  Defaults to _nodata_. | [optional] 
**Destination** | Pointer to **string** | Destination of the matched __TopologyRulePreset__.  Allowed values: - code - pool  Defaults to _code_. | [optional] 
**Name** | **string** | Required. Name of the __TopologyRulePreset__ from the referenced __Topology__. | 
**PoolId** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewPolicyTopologyRuleBinding

`func NewPolicyTopologyRuleBinding(name string, ) *PolicyTopologyRuleBinding`

NewPolicyTopologyRuleBinding instantiates a new PolicyTopologyRuleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTopologyRuleBindingWithDefaults

`func NewPolicyTopologyRuleBindingWithDefaults() *PolicyTopologyRuleBinding`

NewPolicyTopologyRuleBindingWithDefaults instantiates a new PolicyTopologyRuleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PolicyTopologyRuleBinding) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PolicyTopologyRuleBinding) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PolicyTopologyRuleBinding) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PolicyTopologyRuleBinding) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDestination

`func (o *PolicyTopologyRuleBinding) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PolicyTopologyRuleBinding) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PolicyTopologyRuleBinding) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PolicyTopologyRuleBinding) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetName

`func (o *PolicyTopologyRuleBinding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyTopologyRuleBinding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyTopologyRuleBinding) SetName(v string)`

SetName sets Name field to given value.


### GetPoolId

`func (o *PolicyTopologyRuleBinding) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PolicyTopologyRuleBinding) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PolicyTopologyRuleBinding) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *PolicyTopologyRuleBinding) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


