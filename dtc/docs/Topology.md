# Topology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for __Topology__. | [optional] 
**Disabled** | Pointer to **bool** | Optional. Flag which enables/disables __Topology__.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | Output only. __Topology__ metadata.  Defaults to empty object and should be explicitly requested using field selection. | [optional] 
**Name** | **string** | Display name of __Topology__. | 
**Sources** | [**[]TopologySource**](TopologySource.md) | Required. List of __TopologySource__ objects with unique names. | 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __Topology__ in JSON format. | [optional] 

## Methods

### NewTopology

`func NewTopology(name string, sources []TopologySource, ) *Topology`

NewTopology instantiates a new Topology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyWithDefaults

`func NewTopologyWithDefaults() *Topology`

NewTopologyWithDefaults instantiates a new Topology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *Topology) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Topology) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Topology) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Topology) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *Topology) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Topology) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Topology) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Topology) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *Topology) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Topology) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Topology) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Topology) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *Topology) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Topology) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Topology) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Topology) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Topology) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Topology) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Topology) SetName(v string)`

SetName sets Name field to given value.


### GetSources

`func (o *Topology) GetSources() []TopologySource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Topology) GetSourcesOk() (*[]TopologySource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Topology) SetSources(v []TopologySource)`

SetSources sets Sources field to given value.


### GetTags

`func (o *Topology) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Topology) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Topology) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Topology) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


