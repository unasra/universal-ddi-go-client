# UpdateHTTPHealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**HTTPHealthCheck**](HTTPHealthCheck.md) | The __HTTPHealthCheck__ object. | [optional] 

## Methods

### NewUpdateHTTPHealthCheckResponse

`func NewUpdateHTTPHealthCheckResponse() *UpdateHTTPHealthCheckResponse`

NewUpdateHTTPHealthCheckResponse instantiates a new UpdateHTTPHealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHTTPHealthCheckResponseWithDefaults

`func NewUpdateHTTPHealthCheckResponseWithDefaults() *UpdateHTTPHealthCheckResponse`

NewUpdateHTTPHealthCheckResponseWithDefaults instantiates a new UpdateHTTPHealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdateHTTPHealthCheckResponse) GetResult() HTTPHealthCheck`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateHTTPHealthCheckResponse) GetResultOk() (*HTTPHealthCheck, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateHTTPHealthCheckResponse) SetResult(v HTTPHealthCheck)`

SetResult sets Result field to given value.

### HasResult

`func (o *UpdateHTTPHealthCheckResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


