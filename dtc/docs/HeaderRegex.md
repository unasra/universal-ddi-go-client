# HeaderRegex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | **string** | HTTP header name. | 
**Regex** | **string** | Regular expression to match against HTTP header value. | 

## Methods

### NewHeaderRegex

`func NewHeaderRegex(header string, regex string, ) *HeaderRegex`

NewHeaderRegex instantiates a new HeaderRegex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderRegexWithDefaults

`func NewHeaderRegexWithDefaults() *HeaderRegex`

NewHeaderRegexWithDefaults instantiates a new HeaderRegex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *HeaderRegex) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *HeaderRegex) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *HeaderRegex) SetHeader(v string)`

SetHeader sets Header field to given value.


### GetRegex

`func (o *HeaderRegex) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *HeaderRegex) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *HeaderRegex) SetRegex(v string)`

SetRegex sets Regex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


