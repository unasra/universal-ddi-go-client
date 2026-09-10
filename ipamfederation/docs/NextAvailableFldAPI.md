# NextAvailableFldAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNextAvailableFLD**](NextAvailableFldAPI.md#CreateNextAvailableFLD) | **Post** /federation/federated_block/{id}/create_next_available_fld | Create the next available __ForwardLookingDelegation__ objects within a specific __FederatedBlock__. Use this method to create the next available __ForwardLookingDelegation__ objects under the specified parent __FederatedBlock__.
[**CreateNextAvailableFLDBlocks**](NextAvailableFldAPI.md#CreateNextAvailableFLDBlocks) | **Post** /federation/create_next_available_fld | Create the next available __ForwardLookingDelegation__ objects.
[**CreateNextAvailableFLDForPool**](NextAvailableFldAPI.md#CreateNextAvailableFLDForPool) | **Post** /federation/federated_pool/{id}/create_next_available_fld | Create the next available __ForwardLookingDelegation__ objects within a pool.



## CreateNextAvailableFLD

> NextAvailableFLDResponse CreateNextAvailableFLD(ctx, id).Body(body).Execute()

Create the next available __ForwardLookingDelegation__ objects within a specific __FederatedBlock__. Use this method to create the next available __ForwardLookingDelegation__ objects under the specified parent __FederatedBlock__.

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
	body := *ipamfederation.NewCreateNextAvailableFLDRequestForBlock(int64(123), "Id_example") // CreateNextAvailableFLDRequestForBlock | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableFldAPI.CreateNextAvailableFLD(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableFldAPI.CreateNextAvailableFLD``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAvailableFLD`: NextAvailableFLDResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableFldAPI.CreateNextAvailableFLD`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableFldAPICreateNextAvailableFLDRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**CreateNextAvailableFLDRequestForBlock**](CreateNextAvailableFLDRequestForBlock.md) |  | 

### Return type

[**NextAvailableFLDResponse**](NextAvailableFLDResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNextAvailableFLDBlocks

> NextAvailableFLDResponse CreateNextAvailableFLDBlocks(ctx).Body(body).Execute()

Create the next available __ForwardLookingDelegation__ objects.



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
	body := *ipamfederation.NewNextAvailableFLDRequest(int64(123)) // NextAvailableFLDRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableFldAPI.CreateNextAvailableFLDBlocks(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableFldAPI.CreateNextAvailableFLDBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAvailableFLDBlocks`: NextAvailableFLDResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableFldAPI.CreateNextAvailableFLDBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableFldAPICreateNextAvailableFLDBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**NextAvailableFLDRequest**](NextAvailableFLDRequest.md) |  | 

### Return type

[**NextAvailableFLDResponse**](NextAvailableFLDResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNextAvailableFLDForPool

> NextAvailableFLDResponse CreateNextAvailableFLDForPool(ctx, id).Body(body).Execute()

Create the next available __ForwardLookingDelegation__ objects within a pool.



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
	body := *ipamfederation.NewNextAvailableFLDPoolRequest(int64(123), "Id_example") // NextAvailableFLDPoolRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableFldAPI.CreateNextAvailableFLDForPool(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableFldAPI.CreateNextAvailableFLDForPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAvailableFLDForPool`: NextAvailableFLDResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableFldAPI.CreateNextAvailableFLDForPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableFldAPICreateNextAvailableFLDForPoolRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**NextAvailableFLDPoolRequest**](NextAvailableFLDPoolRequest.md) |  | 

### Return type

[**NextAvailableFLDResponse**](NextAvailableFLDResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

