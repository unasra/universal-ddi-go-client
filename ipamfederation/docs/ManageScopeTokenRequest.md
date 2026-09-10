# ManageScopeTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The Infoblox account ID for which the token is generated. | 
**ScopeArn** | **string** |  | 
**Ttl** | Pointer to **int64** | Optional token expiration duration in seconds. Defaults to 300 (5 minutes) if not specified or set to 0. Maximum allowed: 86400 (24 hours), Minimum: 0 (uses default). | [optional] 

## Methods

### NewManageScopeTokenRequest

`func NewManageScopeTokenRequest(accountId string, scopeArn string, ) *ManageScopeTokenRequest`

NewManageScopeTokenRequest instantiates a new ManageScopeTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageScopeTokenRequestWithDefaults

`func NewManageScopeTokenRequestWithDefaults() *ManageScopeTokenRequest`

NewManageScopeTokenRequestWithDefaults instantiates a new ManageScopeTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ManageScopeTokenRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ManageScopeTokenRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ManageScopeTokenRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetScopeArn

`func (o *ManageScopeTokenRequest) GetScopeArn() string`

GetScopeArn returns the ScopeArn field if non-nil, zero value otherwise.

### GetScopeArnOk

`func (o *ManageScopeTokenRequest) GetScopeArnOk() (*string, bool)`

GetScopeArnOk returns a tuple with the ScopeArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeArn

`func (o *ManageScopeTokenRequest) SetScopeArn(v string)`

SetScopeArn sets ScopeArn field to given value.


### GetTtl

`func (o *ManageScopeTokenRequest) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ManageScopeTokenRequest) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ManageScopeTokenRequest) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ManageScopeTokenRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


