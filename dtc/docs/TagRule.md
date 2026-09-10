# TagRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Required. Tag key to match against a source object&#39;s effective tags. | 
**Op** | Pointer to **string** | Optional. Match operator.  Supported values: - EQUALS: matches when the key exists and its value equals the configured value. - NOT_EQUALS: matches when the key exists and all values for that key differ   from the configured value.  A missing key does not satisfy either operator.  Defaults to _EQUALS_. | [optional] 
**Value** | **string** | Required. Tag value to match against a source object&#39;s effective tags. | 

## Methods

### NewTagRule

`func NewTagRule(key string, value string, ) *TagRule`

NewTagRule instantiates a new TagRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagRuleWithDefaults

`func NewTagRuleWithDefaults() *TagRule`

NewTagRuleWithDefaults instantiates a new TagRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TagRule) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TagRule) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TagRule) SetKey(v string)`

SetKey sets Key field to given value.


### GetOp

`func (o *TagRule) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *TagRule) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *TagRule) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *TagRule) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetValue

`func (o *TagRule) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagRule) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagRule) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


