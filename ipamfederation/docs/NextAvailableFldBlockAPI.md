# NextAvailableFldBlockAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNextAvailableFLDBlocks**](NextAvailableFldBlockAPI.md#ListNextAvailableFLDBlocks) | **Get** /federation/list_next_available_fld | List the next available __ForwardLookingDelegation__ objects.



## ListNextAvailableFLDBlocks

> NextAvailableFLDResponse ListNextAvailableFLDBlocks(ctx).Cidr(cidr).Count(count).Name(name).Comment(comment).Protocol(protocol).Execute()

List the next available __ForwardLookingDelegation__ objects.



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

	apiClient := ipamfederation.NewAPIClient()
	resp, r, err := apiClient.NextAvailableFldBlockAPI.ListNextAvailableFLDBlocks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NextAvailableFldBlockAPI.ListNextAvailableFLDBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNextAvailableFLDBlocks`: NextAvailableFLDResponse
	fmt.Fprintf(os.Stdout, "Response from `NextAvailableFldBlockAPI.ListNextAvailableFLDBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NextAvailableFldBlockAPIListNextAvailableFLDBlocksRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**cidr** | **int64** | The CIDR of the __ForwardLookingDelegation__. This is required. | 
**count** | **int64** | The count of __ForwardLookingDelegation__ required. If not provided, it will default to 1. | 
**name** | **string** | The name to be provided. | 
**comment** | **string** | The description for the __ForwardLookingDelegation__. May contain 0 to 1024 characters. Can include UTF-8. | 
**protocol** | **string** | The version of the address (_ip4_ or _ip6_). If not present then it will default to _ip4_. | 

### Return type

[**NextAvailableFLDResponse**](NextAvailableFLDResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

