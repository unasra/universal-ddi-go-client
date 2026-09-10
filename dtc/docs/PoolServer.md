# PoolServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Display name of __Server__. | [optional] [readonly] 
**ServerId** | **string** | The resource identifier. | 
**Weight** | Pointer to **int64** | Weight of __Server__ to be used for load balancing. Unsigned integer, min 1; max 65535. | [optional] 

## Methods

### NewPoolServer

`func NewPoolServer(serverId string, ) *PoolServer`

NewPoolServer instantiates a new PoolServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolServerWithDefaults

`func NewPoolServerWithDefaults() *PoolServer`

NewPoolServerWithDefaults instantiates a new PoolServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PoolServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerId

`func (o *PoolServer) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *PoolServer) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *PoolServer) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetWeight

`func (o *PoolServer) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PoolServer) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PoolServer) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PoolServer) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


