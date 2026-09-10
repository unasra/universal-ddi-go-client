# ReadSIPHealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**SIPHealthCheck**](SIPHealthCheck.md) | The __SIPHealthCheck__ object. | [optional] 

## Methods

### NewReadSIPHealthCheckResponse

`func NewReadSIPHealthCheckResponse() *ReadSIPHealthCheckResponse`

NewReadSIPHealthCheckResponse instantiates a new ReadSIPHealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadSIPHealthCheckResponseWithDefaults

`func NewReadSIPHealthCheckResponseWithDefaults() *ReadSIPHealthCheckResponse`

NewReadSIPHealthCheckResponseWithDefaults instantiates a new ReadSIPHealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ReadSIPHealthCheckResponse) GetResult() SIPHealthCheck`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReadSIPHealthCheckResponse) GetResultOk() (*SIPHealthCheck, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReadSIPHealthCheckResponse) SetResult(v SIPHealthCheck)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReadSIPHealthCheckResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


