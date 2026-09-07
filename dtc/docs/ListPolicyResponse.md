# ListPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Policy**](Policy.md) | List of __Policy__ objects. | [optional] 

## Methods

### NewListPolicyResponse

`func NewListPolicyResponse() *ListPolicyResponse`

NewListPolicyResponse instantiates a new ListPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPolicyResponseWithDefaults

`func NewListPolicyResponseWithDefaults() *ListPolicyResponse`

NewListPolicyResponseWithDefaults instantiates a new ListPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListPolicyResponse) GetResults() []Policy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListPolicyResponse) GetResultsOk() (*[]Policy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListPolicyResponse) SetResults(v []Policy)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListPolicyResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


