# PolicyPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Display name of __Pool__. | [optional] [readonly] 
**PoolId** | **string** | The resource identifier. | 
**Weight** | Pointer to **int64** | Weight of __Pool__ to be used for load balancing. Unsigned integer, min 1; max 65535. | [optional] 

## Methods

### NewPolicyPool

`func NewPolicyPool(poolId string, ) *PolicyPool`

NewPolicyPool instantiates a new PolicyPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyPoolWithDefaults

`func NewPolicyPoolWithDefaults() *PolicyPool`

NewPolicyPoolWithDefaults instantiates a new PolicyPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoolId

`func (o *PolicyPool) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PolicyPool) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PolicyPool) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWeight

`func (o *PolicyPool) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PolicyPool) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PolicyPool) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PolicyPool) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


