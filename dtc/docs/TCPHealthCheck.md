# TCPHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for __TCPHealthCheck__. | [optional] 
**Disabled** | Pointer to **bool** | Optional. Flag which enables/disables __TCPHealthCheck__. Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Interval** | Pointer to **int64** | Optional. Interval value in seconds. The health check runs only for the specified interval and it is measured from the beginning of the previous check cycle. Defaults to _15_. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | Output only. __TCPHealthCheck__ metadata. Defaults to empty object and should be explicitly requested using field selection. | [optional] 
**Name** | **string** | Display name of __TCPHealthCheck__. | 
**Port** | **int64** | Destination TCP port of __TCPHealthCheck__. | 
**RetryDown** | Pointer to **int64** | Optional. Retry down count. The value determines how many bad health checks in a row must be received by the onprem host from the DTC Server for treating the health check as failed. Defaults to _1_. | [optional] 
**RetryUp** | Pointer to **int64** | Optional. Retry up count. The value determines how many good health checks in a row must be received by the onprem host from the DTC Server for treating the health check as successful. Defaults to _1_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __TCPHealthCheck__ in JSON format. | [optional] 
**Timeout** | Pointer to **int64** | Optional. Timeout value in seconds. The health check waits for the specified number of seconds after sending a request. If it does not receive a response within the number of seconds, then the health check is considered as failed. Defaults to _10_. | [optional] 

## Methods

### NewTCPHealthCheck

`func NewTCPHealthCheck(name string, port int64, ) *TCPHealthCheck`

NewTCPHealthCheck instantiates a new TCPHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTCPHealthCheckWithDefaults

`func NewTCPHealthCheckWithDefaults() *TCPHealthCheck`

NewTCPHealthCheckWithDefaults instantiates a new TCPHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *TCPHealthCheck) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TCPHealthCheck) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TCPHealthCheck) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TCPHealthCheck) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *TCPHealthCheck) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *TCPHealthCheck) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *TCPHealthCheck) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *TCPHealthCheck) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *TCPHealthCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TCPHealthCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TCPHealthCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TCPHealthCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *TCPHealthCheck) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TCPHealthCheck) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TCPHealthCheck) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *TCPHealthCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetadata

`func (o *TCPHealthCheck) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TCPHealthCheck) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TCPHealthCheck) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TCPHealthCheck) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *TCPHealthCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TCPHealthCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TCPHealthCheck) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *TCPHealthCheck) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TCPHealthCheck) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TCPHealthCheck) SetPort(v int64)`

SetPort sets Port field to given value.


### GetRetryDown

`func (o *TCPHealthCheck) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *TCPHealthCheck) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *TCPHealthCheck) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *TCPHealthCheck) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *TCPHealthCheck) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *TCPHealthCheck) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *TCPHealthCheck) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *TCPHealthCheck) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTags

`func (o *TCPHealthCheck) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TCPHealthCheck) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TCPHealthCheck) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *TCPHealthCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeout

`func (o *TCPHealthCheck) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TCPHealthCheck) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TCPHealthCheck) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TCPHealthCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


