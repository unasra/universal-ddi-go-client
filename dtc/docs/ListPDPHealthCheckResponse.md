# ListPDPHealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]PDPHealthCheck**](PDPHealthCheck.md) | List of __PDPHealthCheck__ objects. | [optional] 

## Methods

### NewListPDPHealthCheckResponse

`func NewListPDPHealthCheckResponse() *ListPDPHealthCheckResponse`

NewListPDPHealthCheckResponse instantiates a new ListPDPHealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPDPHealthCheckResponseWithDefaults

`func NewListPDPHealthCheckResponseWithDefaults() *ListPDPHealthCheckResponse`

NewListPDPHealthCheckResponseWithDefaults instantiates a new ListPDPHealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListPDPHealthCheckResponse) GetResults() []PDPHealthCheck`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListPDPHealthCheckResponse) GetResultsOk() (*[]PDPHealthCheck, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListPDPHealthCheckResponse) SetResults(v []PDPHealthCheck)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListPDPHealthCheckResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


