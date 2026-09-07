# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | IP Address of the __Server__. Must be set to a valid IP address if __endpoint_type__ is set to __address__. Alternatively, it can be left blank. | [optional] 
**AutoCreateResponseRecords** | Pointer to **bool** | Optional. If the flag is enabled, A, AAAA or CNAME __Record__ is automatically generated.  Defaults to _false_. | [optional] 
**Comment** | Pointer to **string** | Optional. Comment for __Server__. | [optional] 
**Disabled** | Pointer to **bool** | Optional. Flag which enables/disables __Server__.  Defaults to _false_. | [optional] 
**EndpointType** | Pointer to **string** | The endpoint type configured for the __Server__. Can be IP Address or FQDN. The values of both fields __address__ and __fqdn__ are preserved and are not mutually exclusive, and the __endpoint_type__ defines which one to use.  Allowed values: * address * fqdn  Defaults to __address__. | [optional] 
**Fqdn** | Pointer to **string** | Fully Qualified Domain name of the __Server__. Must be set to a valid FQDN if __endpoint_type__ is set to __fqdn__. Alternatively, it can be left blank. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) | Output only. __Server__ metadata. Defaults to empty object and should be explicitly requested using field selection. | [optional] 
**Name** | **string** | Display name of __Server__. | 
**Records** | Pointer to [**[]Record**](Record.md) | Optional. List of __Records__ of the __Server__. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Optional. The tags for __Server__ in JSON format. | [optional] 

## Methods

### NewServer

`func NewServer(name string, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Server) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Server) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Server) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Server) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAutoCreateResponseRecords

`func (o *Server) GetAutoCreateResponseRecords() bool`

GetAutoCreateResponseRecords returns the AutoCreateResponseRecords field if non-nil, zero value otherwise.

### GetAutoCreateResponseRecordsOk

`func (o *Server) GetAutoCreateResponseRecordsOk() (*bool, bool)`

GetAutoCreateResponseRecordsOk returns a tuple with the AutoCreateResponseRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateResponseRecords

`func (o *Server) SetAutoCreateResponseRecords(v bool)`

SetAutoCreateResponseRecords sets AutoCreateResponseRecords field to given value.

### HasAutoCreateResponseRecords

`func (o *Server) HasAutoCreateResponseRecords() bool`

HasAutoCreateResponseRecords returns a boolean if a field has been set.

### GetComment

`func (o *Server) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Server) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Server) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Server) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *Server) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Server) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Server) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Server) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEndpointType

`func (o *Server) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *Server) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *Server) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *Server) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetFqdn

`func (o *Server) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Server) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Server) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Server) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetId

`func (o *Server) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Server) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *Server) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Server) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Server) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Server) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Server) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Server) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Server) SetName(v string)`

SetName sets Name field to given value.


### GetRecords

`func (o *Server) GetRecords() []Record`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *Server) GetRecordsOk() (*[]Record, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *Server) SetRecords(v []Record)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *Server) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTags

`func (o *Server) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Server) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Server) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Server) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


