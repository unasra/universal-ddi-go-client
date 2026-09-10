# MetadataResourceMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **map[string]string** | Structured data consisting of additional details of the configuration resource. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the configuration resource. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 

## Methods

### NewMetadataResourceMeta

`func NewMetadataResourceMeta() *MetadataResourceMeta`

NewMetadataResourceMeta instantiates a new MetadataResourceMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataResourceMetaWithDefaults

`func NewMetadataResourceMetaWithDefaults() *MetadataResourceMeta`

NewMetadataResourceMetaWithDefaults instantiates a new MetadataResourceMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *MetadataResourceMeta) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MetadataResourceMeta) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MetadataResourceMeta) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MetadataResourceMeta) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDisplayName

`func (o *MetadataResourceMeta) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MetadataResourceMeta) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MetadataResourceMeta) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MetadataResourceMeta) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *MetadataResourceMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataResourceMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataResourceMeta) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataResourceMeta) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


