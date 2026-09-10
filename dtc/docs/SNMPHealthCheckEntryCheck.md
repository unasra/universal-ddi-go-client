# SNMPHealthCheckEntryCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for __EntryCheck__. | [optional] 
**MaxValue** | Pointer to **string** | Optional. Expected max value of an entry to check against. Used for __in__ operator only, otherwise ignored. | [optional] 
**Name** | **string** | Name is a dotted-decimal number that defines the location of the entry in the universal MIB tree. | 
**Operator** | **string** | Operator defines operation to perform on an entry value.  Allowed values: * any - any value must be present * eq  - entry value must be equal to check&#39;s __value__. * leq - entry value must less or equal to check&#39;s __value__. * geq - entry value must be great or equal to check&#39;s __value__. * in  - entry value must be greater or equal than __value__ and less or equal than __max_value__.  Operator __in__ is supported only for __integer__ types. | 
**Type** | **string** | Type defines type of an entry value.  Allowed values: * string * integer  String type does not support __in__ operator. | 
**Value** | Pointer to **string** | Optional. Expected value of an entry to check against. Ignored for __any__ operator. | [optional] 

## Methods

### NewSNMPHealthCheckEntryCheck

`func NewSNMPHealthCheckEntryCheck(name string, operator string, type_ string, ) *SNMPHealthCheckEntryCheck`

NewSNMPHealthCheckEntryCheck instantiates a new SNMPHealthCheckEntryCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSNMPHealthCheckEntryCheckWithDefaults

`func NewSNMPHealthCheckEntryCheckWithDefaults() *SNMPHealthCheckEntryCheck`

NewSNMPHealthCheckEntryCheckWithDefaults instantiates a new SNMPHealthCheckEntryCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SNMPHealthCheckEntryCheck) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SNMPHealthCheckEntryCheck) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SNMPHealthCheckEntryCheck) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SNMPHealthCheckEntryCheck) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetMaxValue

`func (o *SNMPHealthCheckEntryCheck) GetMaxValue() string`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *SNMPHealthCheckEntryCheck) GetMaxValueOk() (*string, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *SNMPHealthCheckEntryCheck) SetMaxValue(v string)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *SNMPHealthCheckEntryCheck) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetName

`func (o *SNMPHealthCheckEntryCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SNMPHealthCheckEntryCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SNMPHealthCheckEntryCheck) SetName(v string)`

SetName sets Name field to given value.


### GetOperator

`func (o *SNMPHealthCheckEntryCheck) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SNMPHealthCheckEntryCheck) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SNMPHealthCheckEntryCheck) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetType

`func (o *SNMPHealthCheckEntryCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SNMPHealthCheckEntryCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SNMPHealthCheckEntryCheck) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *SNMPHealthCheckEntryCheck) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SNMPHealthCheckEntryCheck) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SNMPHealthCheckEntryCheck) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SNMPHealthCheckEntryCheck) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


