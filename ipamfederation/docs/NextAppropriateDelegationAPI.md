# NextAppropriateDelegationAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNextAppropriateDelegation**](NextAppropriateDelegationAPI.md#CreateNextAppropriateDelegation) | **Post** /federation/next_appropriate_delegation | Create Next Appropriate Delegation



## CreateNextAppropriateDelegation

> NextAppropriateDelegationResponse CreateNextAppropriateDelegation(ctx).Body(body).Execute()

Create Next Appropriate Delegation



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
	body := *ipamfederation.NewNextAppropriateDelegation() // NextAppropriateDelegation | The __NextAppropriateDelegation__ object.

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAppropriateDelegationAPI.CreateNextAppropriateDelegation(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAppropriateDelegationAPI.CreateNextAppropriateDelegation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNextAppropriateDelegation`: NextAppropriateDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAppropriateDelegationAPI.CreateNextAppropriateDelegation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NextAppropriateDelegationAPICreateNextAppropriateDelegationRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**NextAppropriateDelegation**](NextAppropriateDelegation.md) | The __NextAppropriateDelegation__ object. | 

### Return type

[**NextAppropriateDelegationResponse**](NextAppropriateDelegationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

