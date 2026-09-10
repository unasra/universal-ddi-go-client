# SIPHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for __SIPHealthCheck__. | [optional] 
**Disabled** | Pointer to **bool** | Optional. Flag which enables/disables __SIPHealthCheck__. Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Interval** | Pointer to **int64** | Optional. Interval value in seconds. The health check runs only for the specified interval and it is measured from the beginning of the previous check cycle. Defaults to _15_. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | Output only. __SIPHealthCheck__ metadata. Defaults to empty object and should be explicitly requested using field selection. | [optional] 
**Name** | **string** | Display name of __SIPHealthCheck__. | 
**Port** | Pointer to **int64** | Optional. Destination port for the SIP OPTIONS request. Defaults to _5060_ for both TCP and UDP transport. | [optional] 
**ResultCode** | Pointer to **int64** | Optional. Expected SIP response code, used with _CODE_IS_ and _CODE_IS_NOT_ result modes. Defaults to _200_. | [optional] 
**ResultMode** | Pointer to **string** | Optional. Defines how the SIP response code is evaluated. Supported values: _ANY_, _CODE_IS_, _CODE_IS_NOT_. Defaults to _CODE_IS_. | [optional] 
**RetryDown** | Pointer to **int64** | Optional. Retry down count. The value determines how many bad health checks in a row must be received by the onprem host from the DTC Server for treating the health check as failed. Defaults to _1_. | [optional] 
**RetryUp** | Pointer to **int64** | Optional. Retry up count. The value determines how many good health checks in a row must be received by the onprem host from the DTC Server for treating the health check as successful. Defaults to _1_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __SIPHealthCheck__ in JSON format. | [optional] 
**Timeout** | Pointer to **int64** | Optional. Timeout value in seconds. The health check waits for the specified number of seconds after sending a request. If it does not receive a response within the number of seconds, then the health check is considered as failed. Defaults to _10_. | [optional] 
**Transport** | Pointer to **string** | Optional. Transport protocol for the SIP OPTIONS request. Supported values: _TCP_, _UDP_. Defaults to _TCP_. | [optional] 

## Methods

### NewSIPHealthCheck

`func NewSIPHealthCheck(name string, ) *SIPHealthCheck`

NewSIPHealthCheck instantiates a new SIPHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIPHealthCheckWithDefaults

`func NewSIPHealthCheckWithDefaults() *SIPHealthCheck`

NewSIPHealthCheckWithDefaults instantiates a new SIPHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SIPHealthCheck) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SIPHealthCheck) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SIPHealthCheck) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SIPHealthCheck) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *SIPHealthCheck) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SIPHealthCheck) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SIPHealthCheck) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SIPHealthCheck) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *SIPHealthCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SIPHealthCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SIPHealthCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SIPHealthCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *SIPHealthCheck) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SIPHealthCheck) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SIPHealthCheck) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *SIPHealthCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetadata

`func (o *SIPHealthCheck) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SIPHealthCheck) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SIPHealthCheck) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SIPHealthCheck) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SIPHealthCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SIPHealthCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SIPHealthCheck) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *SIPHealthCheck) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SIPHealthCheck) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SIPHealthCheck) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SIPHealthCheck) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetResultCode

`func (o *SIPHealthCheck) GetResultCode() int64`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *SIPHealthCheck) GetResultCodeOk() (*int64, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *SIPHealthCheck) SetResultCode(v int64)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *SIPHealthCheck) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMode

`func (o *SIPHealthCheck) GetResultMode() string`

GetResultMode returns the ResultMode field if non-nil, zero value otherwise.

### GetResultModeOk

`func (o *SIPHealthCheck) GetResultModeOk() (*string, bool)`

GetResultModeOk returns a tuple with the ResultMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMode

`func (o *SIPHealthCheck) SetResultMode(v string)`

SetResultMode sets ResultMode field to given value.

### HasResultMode

`func (o *SIPHealthCheck) HasResultMode() bool`

HasResultMode returns a boolean if a field has been set.

### GetRetryDown

`func (o *SIPHealthCheck) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *SIPHealthCheck) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *SIPHealthCheck) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *SIPHealthCheck) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *SIPHealthCheck) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *SIPHealthCheck) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *SIPHealthCheck) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *SIPHealthCheck) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTags

`func (o *SIPHealthCheck) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SIPHealthCheck) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SIPHealthCheck) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SIPHealthCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeout

`func (o *SIPHealthCheck) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SIPHealthCheck) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SIPHealthCheck) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SIPHealthCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTransport

`func (o *SIPHealthCheck) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *SIPHealthCheck) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *SIPHealthCheck) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *SIPHealthCheck) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


