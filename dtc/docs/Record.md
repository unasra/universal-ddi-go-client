# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsRdata** | Pointer to **string** | The DNS protocol textual representation of the record data. | [optional] [readonly] 
**Rdata** | **map[string]interface{}** | JSON representation of resource record data. | 
**Type** | **string** | Resource record type.  List of supported types: * _A_ (_TYPE1_) * _AAAA_ (_TYPE28_) * _CNAME_ (_TYPE5_) * _HTTPS_ (_TYPE65_) * _SRV_ (_TYPE33_) * _SVCB_ (_TYPE64_) | 

## Methods

### NewRecord

`func NewRecord(rdata map[string]interface{}, type_ string, ) *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsRdata

`func (o *Record) GetDnsRdata() string`

GetDnsRdata returns the DnsRdata field if non-nil, zero value otherwise.

### GetDnsRdataOk

`func (o *Record) GetDnsRdataOk() (*string, bool)`

GetDnsRdataOk returns a tuple with the DnsRdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRdata

`func (o *Record) SetDnsRdata(v string)`

SetDnsRdata sets DnsRdata field to given value.

### HasDnsRdata

`func (o *Record) HasDnsRdata() bool`

HasDnsRdata returns a boolean if a field has been set.

### GetRdata

`func (o *Record) GetRdata() map[string]interface{}`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *Record) GetRdataOk() (*map[string]interface{}, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *Record) SetRdata(v map[string]interface{})`

SetRdata sets Rdata field to given value.


### GetType

`func (o *Record) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Record) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Record) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


