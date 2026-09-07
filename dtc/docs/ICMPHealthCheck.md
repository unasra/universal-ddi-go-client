# ICMPHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for __ICMPHealthCheck__. | [optional] 
**Disabled** | Pointer to **bool** | Optional. Flag which enables/disables __ICMPHealthCheck__. Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Interval** | Pointer to **int64** | Optional. Interval value in seconds. The health check runs only for the specified interval and it is measured from the beginning of the previous check cycle. Defaults to _15_. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | Output only. __ICMPHealthCheck__ metadata. Defaults to empty object and should be explicitly requested using field selection. | [optional] 
**Name** | **string** | Display name of __ICMPHealthCheck__. | 
**RetryDown** | Pointer to **int64** | Optional. Retry down count. The value determines how many bad health checks in a row must be received by the onprem host from the DTC Server for treating the health check as failed. Defaults to _1_. | [optional] 
**RetryUp** | Pointer to **int64** | Optional. Retry up count. The value determines how many good health checks in a row must be received by the onprem host from the DTC Server for treating the health check as successful. Defaults to _1_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __ICMPHealthCheck__ in JSON format. | [optional] 
**Timeout** | Pointer to **int64** | Optional. Timeout value in seconds. The health check waits for the specified number of seconds after sending a request. If it does not receive a response within the number of seconds, then the health check is considered as failed. Defaults to _10_. | [optional] 

## Methods

### NewICMPHealthCheck

`func NewICMPHealthCheck(name string, ) *ICMPHealthCheck`

NewICMPHealthCheck instantiates a new ICMPHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewICMPHealthCheckWithDefaults

`func NewICMPHealthCheckWithDefaults() *ICMPHealthCheck`

NewICMPHealthCheckWithDefaults instantiates a new ICMPHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ICMPHealthCheck) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ICMPHealthCheck) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ICMPHealthCheck) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ICMPHealthCheck) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *ICMPHealthCheck) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ICMPHealthCheck) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ICMPHealthCheck) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ICMPHealthCheck) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *ICMPHealthCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ICMPHealthCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ICMPHealthCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ICMPHealthCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *ICMPHealthCheck) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ICMPHealthCheck) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ICMPHealthCheck) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ICMPHealthCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetadata

`func (o *ICMPHealthCheck) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ICMPHealthCheck) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ICMPHealthCheck) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ICMPHealthCheck) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ICMPHealthCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ICMPHealthCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ICMPHealthCheck) SetName(v string)`

SetName sets Name field to given value.


### GetRetryDown

`func (o *ICMPHealthCheck) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *ICMPHealthCheck) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *ICMPHealthCheck) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *ICMPHealthCheck) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *ICMPHealthCheck) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *ICMPHealthCheck) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *ICMPHealthCheck) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *ICMPHealthCheck) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTags

`func (o *ICMPHealthCheck) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ICMPHealthCheck) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ICMPHealthCheck) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ICMPHealthCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeout

`func (o *ICMPHealthCheck) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ICMPHealthCheck) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ICMPHealthCheck) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ICMPHealthCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


