# UpdateSIPHealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**SIPHealthCheck**](SIPHealthCheck.md) | The __SIPHealthCheck__ object. | [optional] 

## Methods

### NewUpdateSIPHealthCheckResponse

`func NewUpdateSIPHealthCheckResponse() *UpdateSIPHealthCheckResponse`

NewUpdateSIPHealthCheckResponse instantiates a new UpdateSIPHealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSIPHealthCheckResponseWithDefaults

`func NewUpdateSIPHealthCheckResponseWithDefaults() *UpdateSIPHealthCheckResponse`

NewUpdateSIPHealthCheckResponseWithDefaults instantiates a new UpdateSIPHealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdateSIPHealthCheckResponse) GetResult() SIPHealthCheck`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateSIPHealthCheckResponse) GetResultOk() (*SIPHealthCheck, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateSIPHealthCheckResponse) SetResult(v SIPHealthCheck)`

SetResult sets Result field to given value.

### HasResult

`func (o *UpdateSIPHealthCheckResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


