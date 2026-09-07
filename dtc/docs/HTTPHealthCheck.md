# HTTPHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckResponseBody** | Pointer to **bool** | Optional. Flag which enables checking of the HTTP response body content. Defaults to _false_. | [optional] 
**CheckResponseBodyNegative** | Pointer to **bool** | Optional. Flag which changes the meaning of the regex match result. If set to _true_, the response is valid if regular expression matches not found. Defaults to _false_.  The flag is currently not supported. | [optional] 
**CheckResponseBodyRegex** | Pointer to **string** | Optional. Regular expression to search for a string in the HTTP response body. Error if empty while _check_response_body_ is _true_. Defaults to empty. | [optional] 
**CheckResponseHeader** | Pointer to **bool** | Optional. Flag which enables checking of the HTTP response header(s) content. Defaults to _false_. | [optional] 
**CheckResponseHeaderNegative** | Pointer to **bool** | Optional. Flag which changes the meaning of the header regexes match result. If set to _true_, neither expression matches must be found in their respective headers for the headers to be considered valid. Defaults to _false_. | [optional] 
**CheckResponseHeaderRegexes** | Pointer to [**[]HeaderRegex**](HeaderRegex.md) | Optional. List of (header, regular expression) pairs. All expression matches must be found in their respective headers for the headers to be considered valid. Error if empty while _check_response_header_ is _true_. Defaults to empty. | [optional] 
**Codes** | Pointer to **string** | Optional. Response Status Codes meaning the health check is successful. If empty, any code means success. Individual codes and code ranges are supported, ex. \&quot;102,105-107,109-110,120\&quot;. | [optional] 
**Comment** | Pointer to **string** | Optional. Comment for __HTTPHealthCheck__. | [optional] 
**Disabled** | Pointer to **bool** | Optional. Flag which enables/disables __HTTPHealthCheck__. Defaults to _false_. | [optional] 
**Https** | Pointer to **bool** | Optional. Flag which enables Hypertext Transfer Protocol Secure (HTTPS) in a health check. Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Interval** | Pointer to **int64** | Optional. Interval value in seconds. The health check runs only for the specified interval and it is measured from the beginning of the previous check cycle. Defaults to _15_. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | Output only. __HTTPHealthCheck__ metadata. Defaults to empty object and should be explicitly requested using field selection. | [optional] 
**Name** | **string** | Display name of __HTTPHealthCheck__. | 
**Port** | **int64** | Destination TCP port of __HTTPHealthCheck__. | 
**Request** | Pointer to **string** | HTTP request in a text format, it consists of HTTP method, request target, HTTP headers, request body. | [optional] 
**RetryDown** | Pointer to **int64** | Optional. Retry down count. The value determines how many bad health checks in a row must be received by the onprem host from the DTC Server for treating the health check as failed. Defaults to _1_. | [optional] 
**RetryUp** | Pointer to **int64** | Optional. Retry up count. The value determines how many good health checks in a row must be received by the onprem host from the DTC Server for treating the health check as successful. Defaults to _1_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __HTTPHealthCheck__ in JSON format. | [optional] 
**Timeout** | Pointer to **int64** | Optional. Timeout value in seconds. The health check waits for the specified number of seconds after sending a request. If it does not receive a response within the number of seconds, then the health check is considered as failed. Defaults to _10_. | [optional] 

## Methods

### NewHTTPHealthCheck

`func NewHTTPHealthCheck(name string, port int64, ) *HTTPHealthCheck`

NewHTTPHealthCheck instantiates a new HTTPHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPHealthCheckWithDefaults

`func NewHTTPHealthCheckWithDefaults() *HTTPHealthCheck`

NewHTTPHealthCheckWithDefaults instantiates a new HTTPHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckResponseBody

`func (o *HTTPHealthCheck) GetCheckResponseBody() bool`

GetCheckResponseBody returns the CheckResponseBody field if non-nil, zero value otherwise.

### GetCheckResponseBodyOk

`func (o *HTTPHealthCheck) GetCheckResponseBodyOk() (*bool, bool)`

GetCheckResponseBodyOk returns a tuple with the CheckResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponseBody

`func (o *HTTPHealthCheck) SetCheckResponseBody(v bool)`

SetCheckResponseBody sets CheckResponseBody field to given value.

### HasCheckResponseBody

`func (o *HTTPHealthCheck) HasCheckResponseBody() bool`

HasCheckResponseBody returns a boolean if a field has been set.

### GetCheckResponseBodyNegative

`func (o *HTTPHealthCheck) GetCheckResponseBodyNegative() bool`

GetCheckResponseBodyNegative returns the CheckResponseBodyNegative field if non-nil, zero value otherwise.

### GetCheckResponseBodyNegativeOk

`func (o *HTTPHealthCheck) GetCheckResponseBodyNegativeOk() (*bool, bool)`

GetCheckResponseBodyNegativeOk returns a tuple with the CheckResponseBodyNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponseBodyNegative

`func (o *HTTPHealthCheck) SetCheckResponseBodyNegative(v bool)`

SetCheckResponseBodyNegative sets CheckResponseBodyNegative field to given value.

### HasCheckResponseBodyNegative

`func (o *HTTPHealthCheck) HasCheckResponseBodyNegative() bool`

HasCheckResponseBodyNegative returns a boolean if a field has been set.

### GetCheckResponseBodyRegex

`func (o *HTTPHealthCheck) GetCheckResponseBodyRegex() string`

GetCheckResponseBodyRegex returns the CheckResponseBodyRegex field if non-nil, zero value otherwise.

### GetCheckResponseBodyRegexOk

`func (o *HTTPHealthCheck) GetCheckResponseBodyRegexOk() (*string, bool)`

GetCheckResponseBodyRegexOk returns a tuple with the CheckResponseBodyRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponseBodyRegex

`func (o *HTTPHealthCheck) SetCheckResponseBodyRegex(v string)`

SetCheckResponseBodyRegex sets CheckResponseBodyRegex field to given value.

### HasCheckResponseBodyRegex

`func (o *HTTPHealthCheck) HasCheckResponseBodyRegex() bool`

HasCheckResponseBodyRegex returns a boolean if a field has been set.

### GetCheckResponseHeader

`func (o *HTTPHealthCheck) GetCheckResponseHeader() bool`

GetCheckResponseHeader returns the CheckResponseHeader field if non-nil, zero value otherwise.

### GetCheckResponseHeaderOk

`func (o *HTTPHealthCheck) GetCheckResponseHeaderOk() (*bool, bool)`

GetCheckResponseHeaderOk returns a tuple with the CheckResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponseHeader

`func (o *HTTPHealthCheck) SetCheckResponseHeader(v bool)`

SetCheckResponseHeader sets CheckResponseHeader field to given value.

### HasCheckResponseHeader

`func (o *HTTPHealthCheck) HasCheckResponseHeader() bool`

HasCheckResponseHeader returns a boolean if a field has been set.

### GetCheckResponseHeaderNegative

`func (o *HTTPHealthCheck) GetCheckResponseHeaderNegative() bool`

GetCheckResponseHeaderNegative returns the CheckResponseHeaderNegative field if non-nil, zero value otherwise.

### GetCheckResponseHeaderNegativeOk

`func (o *HTTPHealthCheck) GetCheckResponseHeaderNegativeOk() (*bool, bool)`

GetCheckResponseHeaderNegativeOk returns a tuple with the CheckResponseHeaderNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponseHeaderNegative

`func (o *HTTPHealthCheck) SetCheckResponseHeaderNegative(v bool)`

SetCheckResponseHeaderNegative sets CheckResponseHeaderNegative field to given value.

### HasCheckResponseHeaderNegative

`func (o *HTTPHealthCheck) HasCheckResponseHeaderNegative() bool`

HasCheckResponseHeaderNegative returns a boolean if a field has been set.

### GetCheckResponseHeaderRegexes

`func (o *HTTPHealthCheck) GetCheckResponseHeaderRegexes() []HeaderRegex`

GetCheckResponseHeaderRegexes returns the CheckResponseHeaderRegexes field if non-nil, zero value otherwise.

### GetCheckResponseHeaderRegexesOk

`func (o *HTTPHealthCheck) GetCheckResponseHeaderRegexesOk() (*[]HeaderRegex, bool)`

GetCheckResponseHeaderRegexesOk returns a tuple with the CheckResponseHeaderRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponseHeaderRegexes

`func (o *HTTPHealthCheck) SetCheckResponseHeaderRegexes(v []HeaderRegex)`

SetCheckResponseHeaderRegexes sets CheckResponseHeaderRegexes field to given value.

### HasCheckResponseHeaderRegexes

`func (o *HTTPHealthCheck) HasCheckResponseHeaderRegexes() bool`

HasCheckResponseHeaderRegexes returns a boolean if a field has been set.

### GetCodes

`func (o *HTTPHealthCheck) GetCodes() string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *HTTPHealthCheck) GetCodesOk() (*string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *HTTPHealthCheck) SetCodes(v string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *HTTPHealthCheck) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetComment

`func (o *HTTPHealthCheck) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *HTTPHealthCheck) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *HTTPHealthCheck) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *HTTPHealthCheck) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *HTTPHealthCheck) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *HTTPHealthCheck) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *HTTPHealthCheck) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *HTTPHealthCheck) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetHttps

`func (o *HTTPHealthCheck) GetHttps() bool`

GetHttps returns the Https field if non-nil, zero value otherwise.

### GetHttpsOk

`func (o *HTTPHealthCheck) GetHttpsOk() (*bool, bool)`

GetHttpsOk returns a tuple with the Https field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttps

`func (o *HTTPHealthCheck) SetHttps(v bool)`

SetHttps sets Https field to given value.

### HasHttps

`func (o *HTTPHealthCheck) HasHttps() bool`

HasHttps returns a boolean if a field has been set.

### GetId

`func (o *HTTPHealthCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HTTPHealthCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HTTPHealthCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HTTPHealthCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *HTTPHealthCheck) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *HTTPHealthCheck) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *HTTPHealthCheck) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *HTTPHealthCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetadata

`func (o *HTTPHealthCheck) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HTTPHealthCheck) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HTTPHealthCheck) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HTTPHealthCheck) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *HTTPHealthCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HTTPHealthCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HTTPHealthCheck) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *HTTPHealthCheck) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HTTPHealthCheck) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HTTPHealthCheck) SetPort(v int64)`

SetPort sets Port field to given value.


### GetRequest

`func (o *HTTPHealthCheck) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *HTTPHealthCheck) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *HTTPHealthCheck) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *HTTPHealthCheck) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetRetryDown

`func (o *HTTPHealthCheck) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *HTTPHealthCheck) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *HTTPHealthCheck) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *HTTPHealthCheck) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *HTTPHealthCheck) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *HTTPHealthCheck) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *HTTPHealthCheck) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *HTTPHealthCheck) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTags

`func (o *HTTPHealthCheck) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HTTPHealthCheck) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HTTPHealthCheck) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *HTTPHealthCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeout

`func (o *HTTPHealthCheck) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HTTPHealthCheck) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HTTPHealthCheck) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HTTPHealthCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


