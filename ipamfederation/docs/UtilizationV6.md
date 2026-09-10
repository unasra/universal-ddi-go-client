# UtilizationV6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**Integer128**](Integer128.md) | Total IPv6 addresses. | [optional] [readonly] 
**Used** | Pointer to [**Integer128**](Integer128.md) | Used IPv6 addresses. | [optional] [readonly] 

## Methods

### NewUtilizationV6

`func NewUtilizationV6() *UtilizationV6`

NewUtilizationV6 instantiates a new UtilizationV6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationV6WithDefaults

`func NewUtilizationV6WithDefaults() *UtilizationV6`

NewUtilizationV6WithDefaults instantiates a new UtilizationV6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *UtilizationV6) GetTotal() Integer128`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UtilizationV6) GetTotalOk() (*Integer128, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UtilizationV6) SetTotal(v Integer128)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UtilizationV6) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *UtilizationV6) GetUsed() Integer128`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *UtilizationV6) GetUsedOk() (*Integer128, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *UtilizationV6) SetUsed(v Integer128)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *UtilizationV6) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


