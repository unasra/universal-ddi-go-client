# DesignatorService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServiceId** | **string** | The resource identifier. | 
**DnsServiceName** | Pointer to **string** | Display name of the __DNS Service__. Response-only; ignored on request. | [optional] [readonly] 

## Methods

### NewDesignatorService

`func NewDesignatorService(dnsServiceId string, ) *DesignatorService`

NewDesignatorService instantiates a new DesignatorService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesignatorServiceWithDefaults

`func NewDesignatorServiceWithDefaults() *DesignatorService`

NewDesignatorServiceWithDefaults instantiates a new DesignatorService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServiceId

`func (o *DesignatorService) GetDnsServiceId() string`

GetDnsServiceId returns the DnsServiceId field if non-nil, zero value otherwise.

### GetDnsServiceIdOk

`func (o *DesignatorService) GetDnsServiceIdOk() (*string, bool)`

GetDnsServiceIdOk returns a tuple with the DnsServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServiceId

`func (o *DesignatorService) SetDnsServiceId(v string)`

SetDnsServiceId sets DnsServiceId field to given value.


### GetDnsServiceName

`func (o *DesignatorService) GetDnsServiceName() string`

GetDnsServiceName returns the DnsServiceName field if non-nil, zero value otherwise.

### GetDnsServiceNameOk

`func (o *DesignatorService) GetDnsServiceNameOk() (*string, bool)`

GetDnsServiceNameOk returns a tuple with the DnsServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServiceName

`func (o *DesignatorService) SetDnsServiceName(v string)`

SetDnsServiceName sets DnsServiceName field to given value.

### HasDnsServiceName

`func (o *DesignatorService) HasDnsServiceName() bool`

HasDnsServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


