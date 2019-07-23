# \WebUIApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiffVersionAVersionB**](WebUIApi.md#GetDiffVersionAVersionB) | **Get** /diff/{version_a}/{version_b}/ | Diff route
[**GetUi**](WebUIApi.md#GetUi) | **Get** /ui/ | SwaggerUI route (v3)


# **GetDiffVersionAVersionB**
> string GetDiffVersionAVersionB(ctx, versionA, versionB, optional)
Diff route

Diff two ESI specs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **versionA** | **string**| Swagger spec version to compare with | 
  **versionB** | **string**| Swagger spec version to compare against | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versionA** | **string**| Swagger spec version to compare with | 
 **versionB** | **string**| Swagger spec version to compare against | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUi**
> string GetUi(ctx, optional)
SwaggerUI route (v3)

ESI web UI. This is an open source project. If you find ESI web UI specific bugs, please report them to https://github.com/esi/esi-swagger-ui/issues

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
 **version** | **string**| The Swagger spec version to display | [default to latest]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

