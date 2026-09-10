# ScopeManagementObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether to enable (true) or disable (false) scope management. Defaults to false if not specified. | [optional] 
**ScopeArn** | **string** | AWS IPAM scope ARN to be set as externally managed | 

## Methods

### NewScopeManagementObject

`func NewScopeManagementObject(scopeArn string, ) *ScopeManagementObject`

NewScopeManagementObject instantiates a new ScopeManagementObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeManagementObjectWithDefaults

`func NewScopeManagementObjectWithDefaults() *ScopeManagementObject`

NewScopeManagementObjectWithDefaults instantiates a new ScopeManagementObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ScopeManagementObject) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ScopeManagementObject) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ScopeManagementObject) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ScopeManagementObject) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetScopeArn

`func (o *ScopeManagementObject) GetScopeArn() string`

GetScopeArn returns the ScopeArn field if non-nil, zero value otherwise.

### GetScopeArnOk

`func (o *ScopeManagementObject) GetScopeArnOk() (*string, bool)`

GetScopeArnOk returns a tuple with the ScopeArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeArn

`func (o *ScopeManagementObject) SetScopeArn(v string)`

SetScopeArn sets ScopeArn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


