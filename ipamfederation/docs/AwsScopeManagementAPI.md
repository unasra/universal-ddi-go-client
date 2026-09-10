# AwsScopeManagementAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScopeManagement**](AwsScopeManagementAPI.md#ScopeManagement) | **Post** /federation/aws/scope/management | Configure management settings for AWS IPAM scopes



## ScopeManagement

> ScopeManagementResponse ScopeManagement(ctx).Body(body).Execute()

Configure management settings for AWS IPAM scopes



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
	body := *ipamfederation.NewScopeManagementObject("ScopeArn_example") // ScopeManagementObject | Scope management configuration details

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.AwsScopeManagementAPI.ScopeManagement(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsScopeManagementAPI.ScopeManagement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScopeManagement`: ScopeManagementResponse
	fmt.Fprintf(os.Stdout, "Response from `AwsScopeManagementAPI.ScopeManagement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `AwsScopeManagementAPIScopeManagementRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ScopeManagementObject**](ScopeManagementObject.md) | Scope management configuration details | 

### Return type

[**ScopeManagementResponse**](ScopeManagementResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

