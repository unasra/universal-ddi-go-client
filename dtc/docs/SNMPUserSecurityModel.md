# SNMPUserSecurityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPassphrase** | Pointer to **string** | User passphrase for authentication. Ignored for __NoAuth__, otherwise mandatory. | [optional] 
**AuthProtocol** | Pointer to **string** | Authentication protocol.  Allowed values: * NoAuth * MD5 * SHA  Defaults to __NoAuth__. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**PrivacyPassphrase** | Pointer to **string** | User passphrase for privacy. Ignored for __NoPrivacy__, otherwise mandatory. | [optional] 
**PrivacyProtocol** | Pointer to **string** | Privacy protocol. Must be __NoPrivacy__ if auth_protocol set to __NoAuth__.  Allowed values: * NoPrivacy * DES * AES  Defaults to __NoPrivacy__. | [optional] 
**Username** | Pointer to **string** | User name with which to associate security information. | [optional] 

## Methods

### NewSNMPUserSecurityModel

`func NewSNMPUserSecurityModel() *SNMPUserSecurityModel`

NewSNMPUserSecurityModel instantiates a new SNMPUserSecurityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSNMPUserSecurityModelWithDefaults

`func NewSNMPUserSecurityModelWithDefaults() *SNMPUserSecurityModel`

NewSNMPUserSecurityModelWithDefaults instantiates a new SNMPUserSecurityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthPassphrase

`func (o *SNMPUserSecurityModel) GetAuthPassphrase() string`

GetAuthPassphrase returns the AuthPassphrase field if non-nil, zero value otherwise.

### GetAuthPassphraseOk

`func (o *SNMPUserSecurityModel) GetAuthPassphraseOk() (*string, bool)`

GetAuthPassphraseOk returns a tuple with the AuthPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassphrase

`func (o *SNMPUserSecurityModel) SetAuthPassphrase(v string)`

SetAuthPassphrase sets AuthPassphrase field to given value.

### HasAuthPassphrase

`func (o *SNMPUserSecurityModel) HasAuthPassphrase() bool`

HasAuthPassphrase returns a boolean if a field has been set.

### GetAuthProtocol

`func (o *SNMPUserSecurityModel) GetAuthProtocol() string`

GetAuthProtocol returns the AuthProtocol field if non-nil, zero value otherwise.

### GetAuthProtocolOk

`func (o *SNMPUserSecurityModel) GetAuthProtocolOk() (*string, bool)`

GetAuthProtocolOk returns a tuple with the AuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProtocol

`func (o *SNMPUserSecurityModel) SetAuthProtocol(v string)`

SetAuthProtocol sets AuthProtocol field to given value.

### HasAuthProtocol

`func (o *SNMPUserSecurityModel) HasAuthProtocol() bool`

HasAuthProtocol returns a boolean if a field has been set.

### GetId

`func (o *SNMPUserSecurityModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SNMPUserSecurityModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SNMPUserSecurityModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SNMPUserSecurityModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrivacyPassphrase

`func (o *SNMPUserSecurityModel) GetPrivacyPassphrase() string`

GetPrivacyPassphrase returns the PrivacyPassphrase field if non-nil, zero value otherwise.

### GetPrivacyPassphraseOk

`func (o *SNMPUserSecurityModel) GetPrivacyPassphraseOk() (*string, bool)`

GetPrivacyPassphraseOk returns a tuple with the PrivacyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassphrase

`func (o *SNMPUserSecurityModel) SetPrivacyPassphrase(v string)`

SetPrivacyPassphrase sets PrivacyPassphrase field to given value.

### HasPrivacyPassphrase

`func (o *SNMPUserSecurityModel) HasPrivacyPassphrase() bool`

HasPrivacyPassphrase returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *SNMPUserSecurityModel) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *SNMPUserSecurityModel) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *SNMPUserSecurityModel) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *SNMPUserSecurityModel) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.

### GetUsername

`func (o *SNMPUserSecurityModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SNMPUserSecurityModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SNMPUserSecurityModel) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SNMPUserSecurityModel) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


