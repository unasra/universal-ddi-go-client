# NextAvailablePoolBlockAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNextAvailableBlocks**](NextAvailablePoolBlockAPI.md#CreateNextAvailableBlocks) | **Post** /federation/federated_pool/{id}/next_available_federated_block | Create the next available federated blocks within a pool.



## CreateNextAvailableBlocks

> CreateNextAvailableFederatedBlockResponse CreateNextAvailableBlocks(ctx, id).Body(body).Execute()

Create the next available federated blocks within a pool.



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
	resp, r, err := apiClient.NextAvailablePoolBlockAPI.CreateNextAvailableBlocks(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailablePoolBlockAPI.CreateNextAvailableBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAvailableBlocks`: CreateNextAvailableFederatedBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailablePoolBlockAPI.CreateNextAvailableBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailablePoolBlockAPICreateNextAvailableBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**NextAvailableBlockRequest**](NextAvailableBlockRequest.md) |  | 

### Return type

[**CreateNextAvailableFederatedBlockResponse**](CreateNextAvailableFederatedBlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

