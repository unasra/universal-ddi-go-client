# CreateTopologyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**Topology**](Topology.md) | The created __Topology__ object. | [optional] 

## Methods

### NewCreateTopologyResponse

`func NewCreateTopologyResponse() *CreateTopologyResponse`

NewCreateTopologyResponse instantiates a new CreateTopologyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTopologyResponseWithDefaults

`func NewCreateTopologyResponseWithDefaults() *CreateTopologyResponse`

NewCreateTopologyResponseWithDefaults instantiates a new CreateTopologyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CreateTopologyResponse) GetResult() Topology`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateTopologyResponse) GetResultOk() (*Topology, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateTopologyResponse) SetResult(v Topology)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateTopologyResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


