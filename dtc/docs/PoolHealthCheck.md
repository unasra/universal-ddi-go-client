# PoolHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsolidatedHealthCheck** | Pointer to [**ConsolidatedHealthCheck**](ConsolidatedHealthCheck.md) | Optional. Consolidated health check configuration. When set, the DNS server running the named designator __DNS Service__ performs the health check on behalf of the __Pool__, and the result is shared with other DNS servers linked to the __Pool__ via __LBDN__ association. When unset, each DNS server performs the health check independently and no health status is shared. | [optional] 
**HealthCheckId** | **string** | The resource identifier. | 
**Name** | Pointer to **string** | Display name of __HealthCheck__. | [optional] [readonly] 

## Methods

### NewPoolHealthCheck

`func NewPoolHealthCheck(healthCheckId string, ) *PoolHealthCheck`

NewPoolHealthCheck instantiates a new PoolHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolHealthCheckWithDefaults

`func NewPoolHealthCheckWithDefaults() *PoolHealthCheck`

NewPoolHealthCheckWithDefaults instantiates a new PoolHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsolidatedHealthCheck

`func (o *PoolHealthCheck) GetConsolidatedHealthCheck() ConsolidatedHealthCheck`

GetConsolidatedHealthCheck returns the ConsolidatedHealthCheck field if non-nil, zero value otherwise.

### GetConsolidatedHealthCheckOk

`func (o *PoolHealthCheck) GetConsolidatedHealthCheckOk() (*ConsolidatedHealthCheck, bool)`

GetConsolidatedHealthCheckOk returns a tuple with the ConsolidatedHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedHealthCheck

`func (o *PoolHealthCheck) SetConsolidatedHealthCheck(v ConsolidatedHealthCheck)`

SetConsolidatedHealthCheck sets ConsolidatedHealthCheck field to given value.

### HasConsolidatedHealthCheck

`func (o *PoolHealthCheck) HasConsolidatedHealthCheck() bool`

HasConsolidatedHealthCheck returns a boolean if a field has been set.

### GetHealthCheckId

`func (o *PoolHealthCheck) GetHealthCheckId() string`

GetHealthCheckId returns the HealthCheckId field if non-nil, zero value otherwise.

### GetHealthCheckIdOk

`func (o *PoolHealthCheck) GetHealthCheckIdOk() (*string, bool)`

GetHealthCheckIdOk returns a tuple with the HealthCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckId

`func (o *PoolHealthCheck) SetHealthCheckId(v string)`

SetHealthCheckId sets HealthCheckId field to given value.


### GetName

`func (o *PoolHealthCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolHealthCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolHealthCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolHealthCheck) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


