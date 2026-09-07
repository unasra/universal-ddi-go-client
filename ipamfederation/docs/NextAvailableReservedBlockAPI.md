# NextAvailableReservedBlockAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNextAvailableReservedBlocks**](NextAvailableReservedBlockAPI.md#CreateNextAvailableReservedBlocks) | **Post** /federation/federated_block/{id}/next_available_reserved_block | Retrieve the next available reserved block.
[**ListNextAvailableReservedBlocks**](NextAvailableReservedBlockAPI.md#ListNextAvailableReservedBlocks) | **Get** /federation/federated_block/{id}/next_available_reserved_block | List the next available reserved block.



## CreateNextAvailableReservedBlocks

> CreateNextAvailableReservedBlockResponse CreateNextAvailableReservedBlocks(ctx, id).Body(body).Execute()

Retrieve the next available reserved block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource
	body := *ipamfederation.NewNextAvailableBlockRequest() // NextAvailableBlockRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableReservedBlockAPI.CreateNextAvailableReservedBlocks(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableReservedBlockAPI.CreateNextAvailableReservedBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAvailableReservedBlocks`: CreateNextAvailableReservedBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableReservedBlockAPI.CreateNextAvailableReservedBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableReservedBlockAPICreateNextAvailableReservedBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**NextAvailableBlockRequest**](NextAvailableBlockRequest.md) |  | 

### Return type

[**CreateNextAvailableReservedBlockResponse**](CreateNextAvailableReservedBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNextAvailableReservedBlocks

> ListNextAvailableReservedBlockResponse ListNextAvailableReservedBlocks(ctx, id).Cidr(cidr).Count(count).Name(name).Comment(comment).Execute()

List the next available reserved block.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableReservedBlockAPI.ListNextAvailableReservedBlocks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableReservedBlockAPI.ListNextAvailableReservedBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNextAvailableReservedBlocks`: ListNextAvailableReservedBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableReservedBlockAPI.ListNextAvailableReservedBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableReservedBlockAPIListNextAvailableReservedBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**cidr** | **int64** | The CIDR of the federated block. This is required, if _address_ does not specify it in its input. | 
**count** | **int64** | The count of __Block__ required. If not provided, it will default to 1. | 
**name** | **string** | The name to be provided. | 
**comment** | **string** | The description for the _federation/federated_block_. May contain 0 to 1024 characters. Can include UTF-8. | 

### Return type

[**ListNextAvailableReservedBlockResponse**](ListNextAvailableReservedBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

