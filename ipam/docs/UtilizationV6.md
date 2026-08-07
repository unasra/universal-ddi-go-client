# UtilizationV6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abandoned** | Pointer to [**Integer128**](Integer128.md) | The number of IPV6 addresses in the scope of the object which are in the abandoned state (issued by a DHCP server and then declined by the client). | [optional] [readonly] 
**Dynamic** | Pointer to [**Integer128**](Integer128.md) | The number of IPV6 addresses handed out by DHCP in the scope of the object. This includes all leased addresses, fixed addresses that are defined but not currently leased and abandoned leases. | [optional] [readonly] 
**Static** | Pointer to [**Integer128**](Integer128.md) | The number of defined IPV6 addresses such as reservations or DNS records. It can be computed as _static_ &#x3D; _used_ - _dynamic_. | [optional] [readonly] 
**Total** | Pointer to [**Integer128**](Integer128.md) | The total number of IPV6 addresses available in the scope of the object. | [optional] [readonly] 
**Used** | Pointer to [**Integer128**](Integer128.md) | The number of IPV6 addresses used in the scope of the object. | [optional] [readonly] 

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

### GetAbandoned

`func (o *UtilizationV6) GetAbandoned() Integer128`

GetAbandoned returns the Abandoned field if non-nil, zero value otherwise.

### GetAbandonedOk

`func (o *UtilizationV6) GetAbandonedOk() (*Integer128, bool)`

GetAbandonedOk returns a tuple with the Abandoned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandoned

`func (o *UtilizationV6) SetAbandoned(v Integer128)`

SetAbandoned sets Abandoned field to given value.

### HasAbandoned

`func (o *UtilizationV6) HasAbandoned() bool`

HasAbandoned returns a boolean if a field has been set.

### GetDynamic

`func (o *UtilizationV6) GetDynamic() Integer128`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *UtilizationV6) GetDynamicOk() (*Integer128, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *UtilizationV6) SetDynamic(v Integer128)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *UtilizationV6) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetStatic

`func (o *UtilizationV6) GetStatic() Integer128`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *UtilizationV6) GetStaticOk() (*Integer128, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *UtilizationV6) SetStatic(v Integer128)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *UtilizationV6) HasStatic() bool`

HasStatic returns a boolean if a field has been set.

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


