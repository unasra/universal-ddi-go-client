# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedBy** | Pointer to [**[]MetadataResourceMeta**](MetadataResourceMeta.md) | List of structs representing a limited view on configuration objects that use a resource the metadata is provided for. | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedBy

`func (o *Metadata) GetUsedBy() []MetadataResourceMeta`

GetUsedBy returns the UsedBy field if non-nil, zero value otherwise.

### GetUsedByOk

`func (o *Metadata) GetUsedByOk() (*[]MetadataResourceMeta, bool)`

GetUsedByOk returns a tuple with the UsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBy

`func (o *Metadata) SetUsedBy(v []MetadataResourceMeta)`

SetUsedBy sets UsedBy field to given value.

### HasUsedBy

`func (o *Metadata) HasUsedBy() bool`

HasUsedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


