# ConsolidatedHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Designators** | Pointer to [**[]DesignatorService**](DesignatorService.md) | Designator __DNS Service__ references where the corresponding health checks will be associated to. Must contain at least one entry when set.  On request: only _dns_service_id_ is honoured. On response: _dns_service_name_ is echoed alongside, resolved from inventory. | [optional] 

## Methods

### NewConsolidatedHealthCheck

`func NewConsolidatedHealthCheck() *ConsolidatedHealthCheck`

NewConsolidatedHealthCheck instantiates a new ConsolidatedHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsolidatedHealthCheckWithDefaults

`func NewConsolidatedHealthCheckWithDefaults() *ConsolidatedHealthCheck`

NewConsolidatedHealthCheckWithDefaults instantiates a new ConsolidatedHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesignators

`func (o *ConsolidatedHealthCheck) GetDesignators() []DesignatorService`

GetDesignators returns the Designators field if non-nil, zero value otherwise.

### GetDesignatorsOk

`func (o *ConsolidatedHealthCheck) GetDesignatorsOk() (*[]DesignatorService, bool)`

GetDesignatorsOk returns a tuple with the Designators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignators

`func (o *ConsolidatedHealthCheck) SetDesignators(v []DesignatorService)`

SetDesignators sets Designators field to given value.

### HasDesignators

`func (o *ConsolidatedHealthCheck) HasDesignators() bool`

HasDesignators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


