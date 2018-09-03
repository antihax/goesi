# \MetaApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeaders**](MetaApi.md#GetHeaders) | **Get** /headers/ | Debug request headers
[**GetPing**](MetaApi.md#GetPing) | **Get** /ping | Ping route
[**GetStatus**](MetaApi.md#GetStatus) | **Get** /status.json | ESI health status
[**GetVerify**](MetaApi.md#GetVerify) | **Get** /verify/ | Verify access token
[**GetVersions**](MetaApi.md#GetVersions) | **Get** /versions/ | List versions


# **GetHeaders**
> map[string]string GetHeaders(ctx, )
Debug request headers

Echo the request headers for debugging purposes. Note that the 'Connection' header and any 'X-' headers are not included

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPing**
> string GetPing(ctx, )
Ping route

Ping the ESI routers

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatus**
> []GetStatusItem GetStatus(ctx, optional)
ESI health status

Provides a general health indicator per route and method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string**| The version of metrics to request. Note that alternate versions are grouped together | [default to latest]

### Return type

[**[]GetStatusItem**](get_status_item.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerify**
> GetVerifyOk GetVerify(ctx, optional)
Verify access token

Verify authorization tokens in ESI's auth cache

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
 **authorization** | **string**| Access token, in the format of \&quot;Bearer &lt;access token&gt;\&quot; | 

### Return type

[**GetVerifyOk**](get_verify_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersions**
> []string GetVersions(ctx, optional)
List versions

List all endpoint versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

