# AwsScopeTokenAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateToken**](AwsScopeTokenAPI.md#GenerateToken) | **Post** /federation/aws/scope/token | Generate AWS IPAM scope authentication token



## GenerateToken

> ManageScopeTokenResponse GenerateToken(ctx).Body(body).Execute()

Generate AWS IPAM scope authentication token



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
	body := *ipamfederation.NewManageScopeTokenRequest("AccountId_example", "ScopeArn_example") // ManageScopeTokenRequest | 

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.AwsScopeTokenAPI.GenerateToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsScopeTokenAPI.GenerateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateToken`: ManageScopeTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AwsScopeTokenAPI.GenerateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `AwsScopeTokenAPIGenerateTokenRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ManageScopeTokenRequest**](ManageScopeTokenRequest.md) |  | 

### Return type

[**ManageScopeTokenResponse**](ManageScopeTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

