# NextAppropriateDelegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **int64** | The CIDR Prefix Length of the required __Delegation__ object. This is a required field. | [optional] 
**Count** | Pointer to **int64** | The count of the unique allocation with the provided CIDR. If not present then it will default to 1. | [optional] 
**DelegatedTo** | Pointer to **string** | The identifier, which denotes as to whom the allocation would belong to. This field would take the format of &lt;provider_name&gt;://&lt;resource_to_be_assigned_to_id&gt;. Like for AWS, it will be aws://&lt;scope_arn&gt;. | [optional] 
**Exclusions** | Pointer to **[]string** | The exclusions are the list of address prefixes (in the form of __IP_ADDRESS__/__CIDR__) which should be excluded while calculating for available allocations. This is an optional parameter. | [optional] 
**IdempotencyKey** | Pointer to **string** | If the idempotency key is provided, a previously successful request using an identical key will be returned rather than being processed as an additional request. | [optional] 
**Inclusions** | Pointer to **[]string** | The inclusions are the list of address prefixes (in the form of __IP_ADDRESS__/__CIDR__) which should be considered while calculating for available allocations. This is an optional parameter. | [optional] 
**Name** | Pointer to **string** | The name of the __Delegation__ object created. | [optional] 
**PoolId** | Pointer to **string** | The unique identifier of the foreign __IPAM__ __Pool__ for which the request is being made, so that addresses appropriate for the __Pool__ may be selected. | [optional] 
**Protocol** | Pointer to **string** | The version of the address (_ip4_ or _ip6_). If not present then it will default to _ip4_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the __Delegation__ object created. It will be a JSON object. | [optional] 

## Methods

### NewNextAppropriateDelegation

`func NewNextAppropriateDelegation() *NextAppropriateDelegation`

NewNextAppropriateDelegation instantiates a new NextAppropriateDelegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextAppropriateDelegationWithDefaults

`func NewNextAppropriateDelegationWithDefaults() *NextAppropriateDelegation`

NewNextAppropriateDelegationWithDefaults instantiates a new NextAppropriateDelegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *NextAppropriateDelegation) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NextAppropriateDelegation) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NextAppropriateDelegation) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NextAppropriateDelegation) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCount

`func (o *NextAppropriateDelegation) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NextAppropriateDelegation) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NextAppropriateDelegation) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *NextAppropriateDelegation) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDelegatedTo

`func (o *NextAppropriateDelegation) GetDelegatedTo() string`

GetDelegatedTo returns the DelegatedTo field if non-nil, zero value otherwise.

### GetDelegatedToOk

`func (o *NextAppropriateDelegation) GetDelegatedToOk() (*string, bool)`

GetDelegatedToOk returns a tuple with the DelegatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedTo

`func (o *NextAppropriateDelegation) SetDelegatedTo(v string)`

SetDelegatedTo sets DelegatedTo field to given value.

### HasDelegatedTo

`func (o *NextAppropriateDelegation) HasDelegatedTo() bool`

HasDelegatedTo returns a boolean if a field has been set.

### GetExclusions

`func (o *NextAppropriateDelegation) GetExclusions() []string`

GetExclusions returns the Exclusions field if non-nil, zero value otherwise.

### GetExclusionsOk

`func (o *NextAppropriateDelegation) GetExclusionsOk() (*[]string, bool)`

GetExclusionsOk returns a tuple with the Exclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusions

`func (o *NextAppropriateDelegation) SetExclusions(v []string)`

SetExclusions sets Exclusions field to given value.

### HasExclusions

`func (o *NextAppropriateDelegation) HasExclusions() bool`

HasExclusions returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *NextAppropriateDelegation) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *NextAppropriateDelegation) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *NextAppropriateDelegation) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *NextAppropriateDelegation) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetInclusions

`func (o *NextAppropriateDelegation) GetInclusions() []string`

GetInclusions returns the Inclusions field if non-nil, zero value otherwise.

### GetInclusionsOk

`func (o *NextAppropriateDelegation) GetInclusionsOk() (*[]string, bool)`

GetInclusionsOk returns a tuple with the Inclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusions

`func (o *NextAppropriateDelegation) SetInclusions(v []string)`

SetInclusions sets Inclusions field to given value.

### HasInclusions

`func (o *NextAppropriateDelegation) HasInclusions() bool`

HasInclusions returns a boolean if a field has been set.

### GetName

`func (o *NextAppropriateDelegation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NextAppropriateDelegation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NextAppropriateDelegation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NextAppropriateDelegation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoolId

`func (o *NextAppropriateDelegation) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *NextAppropriateDelegation) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *NextAppropriateDelegation) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *NextAppropriateDelegation) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetProtocol

`func (o *NextAppropriateDelegation) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NextAppropriateDelegation) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NextAppropriateDelegation) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NextAppropriateDelegation) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTags

`func (o *NextAppropriateDelegation) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NextAppropriateDelegation) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NextAppropriateDelegation) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *NextAppropriateDelegation) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


