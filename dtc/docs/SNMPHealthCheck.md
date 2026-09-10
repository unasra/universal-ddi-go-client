# SNMPHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckList** | Pointer to [**[]SNMPHealthCheckEntryCheck**](SNMPHealthCheckEntryCheck.md) | List of specific checks for SNMP entries and their values in MIB hierarchy. Supported up to 15 checks. | [optional] 
**Comment** | Pointer to **string** | Optional. Comment for __SNMPHealthCheck__. | [optional] 
**Community** | Pointer to **string** | Optional. SNMP community string used for authentication. Mandatory for __v1__ and __v2c__ versions, ignored for __v3__.  Defaults to __public__. | [optional] 
**ContextEngineId** | Pointer to **string** | Optional. Uniquely identifies an SNMP entity that may realize an instance of a context with a particular context name.  Format is an arbitrary string that can contain from 10 to 64 hexadecimal digits (5 to 32 octet numbers).  Ignored for __v1__ and __v2c__ versions. | [optional] 
**ContextName** | Pointer to **string** | Optional. Name of administratively unique context for __v3__ version. Ignored for __v1__ and __v2c__ versions. | [optional] 
**Disabled** | Pointer to **bool** | Optional. Flag which enables/disables __SNMPHealthCheck__. Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Interval** | Pointer to **int64** | Optional. Interval value in seconds. The health check runs only for the specified interval and it is measured from the beginning of the previous check cycle. Defaults to _15_. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | Output only. __SNMPHealthCheck__ metadata. Defaults to empty object and should be explicitly requested using field selection. | [optional] 
**Name** | **string** | Display name of __SNMPHealthCheck__. | 
**Port** | Pointer to **int64** | Optional. Destination UDP port of __SNMPHealthCheck__. Defaults to _161_. | [optional] 
**RetryDown** | Pointer to **int64** | Optional. Retry down count. The value determines how many bad health checks in a row must be received by the onprem host from the DTC Server for treating the health check as failed. Defaults to _1_. | [optional] 
**RetryUp** | Pointer to **int64** | Optional. Retry up count. The value determines how many good health checks in a row must be received by the onprem host from the DTC Server for treating the health check as successful. Defaults to _1_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __SNMPHealthCheck__ in JSON format. | [optional] 
**Timeout** | Pointer to **int64** | Optional. Timeout value in seconds. The health check waits for the specified number of seconds after sending a request. If it does not receive a response within the number of seconds, then the health check is considered as failed. Defaults to _10_. | [optional] 
**UserSecurityModel** | Pointer to **string** | The resource identifier. | [optional] 
**Version** | **string** | SNMP version.  Allowed values: * v1  - version 1 * v2c - version 2 community * v3  - version 3 | 

## Methods

### NewSNMPHealthCheck

`func NewSNMPHealthCheck(name string, version string, ) *SNMPHealthCheck`

NewSNMPHealthCheck instantiates a new SNMPHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSNMPHealthCheckWithDefaults

`func NewSNMPHealthCheckWithDefaults() *SNMPHealthCheck`

NewSNMPHealthCheckWithDefaults instantiates a new SNMPHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckList

`func (o *SNMPHealthCheck) GetCheckList() []SNMPHealthCheckEntryCheck`

GetCheckList returns the CheckList field if non-nil, zero value otherwise.

### GetCheckListOk

`func (o *SNMPHealthCheck) GetCheckListOk() (*[]SNMPHealthCheckEntryCheck, bool)`

GetCheckListOk returns a tuple with the CheckList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckList

`func (o *SNMPHealthCheck) SetCheckList(v []SNMPHealthCheckEntryCheck)`

SetCheckList sets CheckList field to given value.

### HasCheckList

`func (o *SNMPHealthCheck) HasCheckList() bool`

HasCheckList returns a boolean if a field has been set.

### GetComment

`func (o *SNMPHealthCheck) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SNMPHealthCheck) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SNMPHealthCheck) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SNMPHealthCheck) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCommunity

`func (o *SNMPHealthCheck) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *SNMPHealthCheck) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *SNMPHealthCheck) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *SNMPHealthCheck) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetContextEngineId

`func (o *SNMPHealthCheck) GetContextEngineId() string`

GetContextEngineId returns the ContextEngineId field if non-nil, zero value otherwise.

### GetContextEngineIdOk

`func (o *SNMPHealthCheck) GetContextEngineIdOk() (*string, bool)`

GetContextEngineIdOk returns a tuple with the ContextEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextEngineId

`func (o *SNMPHealthCheck) SetContextEngineId(v string)`

SetContextEngineId sets ContextEngineId field to given value.

### HasContextEngineId

`func (o *SNMPHealthCheck) HasContextEngineId() bool`

HasContextEngineId returns a boolean if a field has been set.

### GetContextName

`func (o *SNMPHealthCheck) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *SNMPHealthCheck) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *SNMPHealthCheck) SetContextName(v string)`

SetContextName sets ContextName field to given value.

### HasContextName

`func (o *SNMPHealthCheck) HasContextName() bool`

HasContextName returns a boolean if a field has been set.

### GetDisabled

`func (o *SNMPHealthCheck) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SNMPHealthCheck) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SNMPHealthCheck) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SNMPHealthCheck) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *SNMPHealthCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SNMPHealthCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SNMPHealthCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SNMPHealthCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *SNMPHealthCheck) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SNMPHealthCheck) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SNMPHealthCheck) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *SNMPHealthCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetadata

`func (o *SNMPHealthCheck) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SNMPHealthCheck) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SNMPHealthCheck) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SNMPHealthCheck) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SNMPHealthCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SNMPHealthCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SNMPHealthCheck) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *SNMPHealthCheck) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SNMPHealthCheck) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SNMPHealthCheck) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SNMPHealthCheck) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRetryDown

`func (o *SNMPHealthCheck) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *SNMPHealthCheck) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *SNMPHealthCheck) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *SNMPHealthCheck) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *SNMPHealthCheck) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *SNMPHealthCheck) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *SNMPHealthCheck) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *SNMPHealthCheck) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTags

`func (o *SNMPHealthCheck) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SNMPHealthCheck) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SNMPHealthCheck) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SNMPHealthCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeout

`func (o *SNMPHealthCheck) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SNMPHealthCheck) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SNMPHealthCheck) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SNMPHealthCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUserSecurityModel

`func (o *SNMPHealthCheck) GetUserSecurityModel() string`

GetUserSecurityModel returns the UserSecurityModel field if non-nil, zero value otherwise.

### GetUserSecurityModelOk

`func (o *SNMPHealthCheck) GetUserSecurityModelOk() (*string, bool)`

GetUserSecurityModelOk returns a tuple with the UserSecurityModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSecurityModel

`func (o *SNMPHealthCheck) SetUserSecurityModel(v string)`

SetUserSecurityModel sets UserSecurityModel field to given value.

### HasUserSecurityModel

`func (o *SNMPHealthCheck) HasUserSecurityModel() bool`

HasUserSecurityModel returns a boolean if a field has been set.

### GetVersion

`func (o *SNMPHealthCheck) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SNMPHealthCheck) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SNMPHealthCheck) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


