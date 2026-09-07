# Pool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for __Pool__. | [optional] 
**ConsolidatedHealthEnabled** | Pointer to **bool** | Optional. Pool-level switch that enables consolidated health probing for this __Pool__.  Defaults to _false_ (consolidated health disabled). Set to _true_ to enable consolidated probing on this __Pool__. When _false_, any per-__PoolHealthCheck__ consolidation configuration is preserved in storage but suppressed at runtime. | [optional] 
**Disabled** | Pointer to **bool** | Optional. Flag which enables/disables __Pool__.  Defaults to _false_. | [optional] 
**HealthChecks** | Pointer to [**[]PoolHealthCheck**](PoolHealthCheck.md) | Optional. List of __HealthCheck__ objects IDs assigned to __Pool__.  Defaults to _empty_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**TTLInheritance**](TTLInheritance.md) | Optional. The inheritance configuration. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | Output only. __Pool__ metadata. Defaults to empty object and should be explicitly requested using field selection. | [optional] 
**Method** | **string** | Load balancing method used for selecting __Server__ assigned to __Pool__.  Valid values are: * _round_robin_ If the _round_robin_ load balancing method is selected, Universal DDI adjusts the response to a query in a sequential and circular manner, directing clients to pools.  * _ratio_ If _ratio_ load balancing method is selected, Universal DDI adjusts the response to a query so that clients are directed to pool using weighted round robin, a load-balancing pattern in which requests are distributed among several resources based on weight assigned to each resource. The distribution of responses over time will be equal for all available pools but the sequence of the responses won&#39;t be guaranteed. When equal weights are assigned for resources (pools) it effectively leads to basic round robin which directs clients to pools in sequential and circular manner.  * _global_availability_ If _global_availability_ load balancing method is selected clients are directed to the first server that is up in the _servers_ list.  Defaults to _round_robin_. | 
**Name** | **string** | Display name of __Pool__. | 
**PoolAvailability** | Pointer to **string** | Optional. Pool Availability setting defines how __Pool__ health is calculated.  Valid values are: * _all_ If _all_ availability selected then __Pool__ is treated healthy when all pool&#39;s servers are healthy. * _quorum_ If _quorum_ availability selected then __Pool__ is treated healthy when at least N pool&#39;s servers are healthy. N is configurable via the value from _pool_servers_quorum_ setting. * _any_ If _any_ availability selected then __Pool__ is treated healthy when at least one pool&#39;s server is healthy.  Defaults to _any_. | [optional] 
**PoolServersQuorum** | Pointer to **int64** | Pool Servers Quorum defines a minimal number of pool&#39;s healthy servers required for treating __Pool__ as healthy when Pool Availability is set to _quorum_. | [optional] 
**ServerAvailability** | Pointer to **string** | Optional. Server Availability setting defines how __Server__ health is calculated.  Valid values are: * _all_ If _all_ availability selected then __Server__ is treated healthy when all pool&#39;s health checks are positive. * _quorum_ If _quorum_ availability selected then __Server__ is treated healthy when at least N pool&#39;s health checks are positive. N is configurable via the value from _server_health_checks_quorum_ setting. * _any_ If _any_ availability selected then __Server__ is treated healthy when at least one pool&#39;s health check is positive  Defaults to _all_. | [optional] 
**ServerHealthChecksQuorum** | Pointer to **int64** | Server Health Checks Quorum defines a minimal number of pool&#39;s positive health checks required for treating __Server__ as healthy when Server Availability is set to _quorum_. | [optional] 
**Servers** | Pointer to [**[]PoolServer**](PoolServer.md) | Optional. List of __Server__ objects assigned to __Pool__.  Defaults to _empty_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __Pool__ in JSON format. | [optional] 
**Ttl** | Pointer to **int64** | Optional. Time to live value (in seconds) to be used for records in DTC response. Unsigned integer, min: 0, max 2147483647 (31-bits per RFC-2181). | [optional] 

## Methods

### NewPool

`func NewPool(method string, name string, ) *Pool`

NewPool instantiates a new Pool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolWithDefaults

`func NewPoolWithDefaults() *Pool`

NewPoolWithDefaults instantiates a new Pool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *Pool) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Pool) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Pool) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Pool) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConsolidatedHealthEnabled

`func (o *Pool) GetConsolidatedHealthEnabled() bool`

GetConsolidatedHealthEnabled returns the ConsolidatedHealthEnabled field if non-nil, zero value otherwise.

### GetConsolidatedHealthEnabledOk

`func (o *Pool) GetConsolidatedHealthEnabledOk() (*bool, bool)`

GetConsolidatedHealthEnabledOk returns a tuple with the ConsolidatedHealthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedHealthEnabled

`func (o *Pool) SetConsolidatedHealthEnabled(v bool)`

SetConsolidatedHealthEnabled sets ConsolidatedHealthEnabled field to given value.

### HasConsolidatedHealthEnabled

`func (o *Pool) HasConsolidatedHealthEnabled() bool`

HasConsolidatedHealthEnabled returns a boolean if a field has been set.

### GetDisabled

`func (o *Pool) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Pool) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Pool) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Pool) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetHealthChecks

`func (o *Pool) GetHealthChecks() []PoolHealthCheck`

GetHealthChecks returns the HealthChecks field if non-nil, zero value otherwise.

### GetHealthChecksOk

`func (o *Pool) GetHealthChecksOk() (*[]PoolHealthCheck, bool)`

GetHealthChecksOk returns a tuple with the HealthChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthChecks

`func (o *Pool) SetHealthChecks(v []PoolHealthCheck)`

SetHealthChecks sets HealthChecks field to given value.

### HasHealthChecks

`func (o *Pool) HasHealthChecks() bool`

HasHealthChecks returns a boolean if a field has been set.

### GetId

`func (o *Pool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Pool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *Pool) GetInheritanceSources() TTLInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *Pool) GetInheritanceSourcesOk() (*TTLInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *Pool) SetInheritanceSources(v TTLInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *Pool) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetMetadata

`func (o *Pool) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Pool) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Pool) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Pool) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMethod

`func (o *Pool) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Pool) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Pool) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetName

`func (o *Pool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pool) SetName(v string)`

SetName sets Name field to given value.


### GetPoolAvailability

`func (o *Pool) GetPoolAvailability() string`

GetPoolAvailability returns the PoolAvailability field if non-nil, zero value otherwise.

### GetPoolAvailabilityOk

`func (o *Pool) GetPoolAvailabilityOk() (*string, bool)`

GetPoolAvailabilityOk returns a tuple with the PoolAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAvailability

`func (o *Pool) SetPoolAvailability(v string)`

SetPoolAvailability sets PoolAvailability field to given value.

### HasPoolAvailability

`func (o *Pool) HasPoolAvailability() bool`

HasPoolAvailability returns a boolean if a field has been set.

### GetPoolServersQuorum

`func (o *Pool) GetPoolServersQuorum() int64`

GetPoolServersQuorum returns the PoolServersQuorum field if non-nil, zero value otherwise.

### GetPoolServersQuorumOk

`func (o *Pool) GetPoolServersQuorumOk() (*int64, bool)`

GetPoolServersQuorumOk returns a tuple with the PoolServersQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolServersQuorum

`func (o *Pool) SetPoolServersQuorum(v int64)`

SetPoolServersQuorum sets PoolServersQuorum field to given value.

### HasPoolServersQuorum

`func (o *Pool) HasPoolServersQuorum() bool`

HasPoolServersQuorum returns a boolean if a field has been set.

### GetServerAvailability

`func (o *Pool) GetServerAvailability() string`

GetServerAvailability returns the ServerAvailability field if non-nil, zero value otherwise.

### GetServerAvailabilityOk

`func (o *Pool) GetServerAvailabilityOk() (*string, bool)`

GetServerAvailabilityOk returns a tuple with the ServerAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAvailability

`func (o *Pool) SetServerAvailability(v string)`

SetServerAvailability sets ServerAvailability field to given value.

### HasServerAvailability

`func (o *Pool) HasServerAvailability() bool`

HasServerAvailability returns a boolean if a field has been set.

### GetServerHealthChecksQuorum

`func (o *Pool) GetServerHealthChecksQuorum() int64`

GetServerHealthChecksQuorum returns the ServerHealthChecksQuorum field if non-nil, zero value otherwise.

### GetServerHealthChecksQuorumOk

`func (o *Pool) GetServerHealthChecksQuorumOk() (*int64, bool)`

GetServerHealthChecksQuorumOk returns a tuple with the ServerHealthChecksQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHealthChecksQuorum

`func (o *Pool) SetServerHealthChecksQuorum(v int64)`

SetServerHealthChecksQuorum sets ServerHealthChecksQuorum field to given value.

### HasServerHealthChecksQuorum

`func (o *Pool) HasServerHealthChecksQuorum() bool`

HasServerHealthChecksQuorum returns a boolean if a field has been set.

### GetServers

`func (o *Pool) GetServers() []PoolServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Pool) GetServersOk() (*[]PoolServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Pool) SetServers(v []PoolServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Pool) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTags

`func (o *Pool) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Pool) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Pool) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Pool) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTtl

`func (o *Pool) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Pool) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Pool) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Pool) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


