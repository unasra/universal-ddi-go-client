# FederatedRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationV4** | Pointer to [**Allocation**](Allocation.md) | The aggregate of all Federated Blocks within the Realm. | [optional] 
**Comment** | Pointer to **string** | The description of the federated realm. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | The metadata for the federated realm in JSON format. | [optional] 
**Name** | **string** | The name of the federated realm. May contain 1 to 256 characters; can include UTF-8. | 
**Provider** | Pointer to [**ProviderType**](ProviderType.md) | The provider type this realm is associated with. | [optional] [default to PROVIDERTYPE_NIOS_X]
**Region** | Pointer to **string** | The region where the realm is located. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the federated realm in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Utilization** | Pointer to **int64** | The percentage of Federated Realm utilization. | [optional] 
**UtilizationV6** | Pointer to [**UtilizationV6**](UtilizationV6.md) | The IPv6 utilization metrics for the federated realm. | [optional] [readonly] 

## Methods

### NewFederatedRealm

`func NewFederatedRealm(name string, ) *FederatedRealm`

NewFederatedRealm instantiates a new FederatedRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederatedRealmWithDefaults

`func NewFederatedRealmWithDefaults() *FederatedRealm`

NewFederatedRealmWithDefaults instantiates a new FederatedRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationV4

`func (o *FederatedRealm) GetAllocationV4() Allocation`

GetAllocationV4 returns the AllocationV4 field if non-nil, zero value otherwise.

### GetAllocationV4Ok

`func (o *FederatedRealm) GetAllocationV4Ok() (*Allocation, bool)`

GetAllocationV4Ok returns a tuple with the AllocationV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationV4

`func (o *FederatedRealm) SetAllocationV4(v Allocation)`

SetAllocationV4 sets AllocationV4 field to given value.

### HasAllocationV4

`func (o *FederatedRealm) HasAllocationV4() bool`

HasAllocationV4 returns a boolean if a field has been set.

### GetComment

`func (o *FederatedRealm) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FederatedRealm) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FederatedRealm) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FederatedRealm) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FederatedRealm) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FederatedRealm) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FederatedRealm) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FederatedRealm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *FederatedRealm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederatedRealm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederatedRealm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FederatedRealm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FederatedRealm) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FederatedRealm) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FederatedRealm) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FederatedRealm) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *FederatedRealm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederatedRealm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederatedRealm) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *FederatedRealm) GetProvider() ProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *FederatedRealm) GetProviderOk() (*ProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *FederatedRealm) SetProvider(v ProviderType)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *FederatedRealm) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *FederatedRealm) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FederatedRealm) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FederatedRealm) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FederatedRealm) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetTags

`func (o *FederatedRealm) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FederatedRealm) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FederatedRealm) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FederatedRealm) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FederatedRealm) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FederatedRealm) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FederatedRealm) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FederatedRealm) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUtilization

`func (o *FederatedRealm) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *FederatedRealm) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *FederatedRealm) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *FederatedRealm) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationV6

`func (o *FederatedRealm) GetUtilizationV6() UtilizationV6`

GetUtilizationV6 returns the UtilizationV6 field if non-nil, zero value otherwise.

### GetUtilizationV6Ok

`func (o *FederatedRealm) GetUtilizationV6Ok() (*UtilizationV6, bool)`

GetUtilizationV6Ok returns a tuple with the UtilizationV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationV6

`func (o *FederatedRealm) SetUtilizationV6(v UtilizationV6)`

SetUtilizationV6 sets UtilizationV6 field to given value.

### HasUtilizationV6

`func (o *FederatedRealm) HasUtilizationV6() bool`

HasUtilizationV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


