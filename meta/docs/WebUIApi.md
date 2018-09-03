# \WebUIApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDev**](WebUIApi.md#GetDev) | **Get** /dev/ | SwaggerUI route (v2)
[**GetDiffVersionAVersionB**](WebUIApi.md#GetDiffVersionAVersionB) | **Get** /diff/{version_a}/{version_b}/ | Diff route
[**GetLatest**](WebUIApi.md#GetLatest) | **Get** /latest/ | SwaggerUI route (v2)
[**GetLegacy**](WebUIApi.md#GetLegacy) | **Get** /legacy/ | SwaggerUI route (v2)
[**GetUi**](WebUIApi.md#GetUi) | **Get** /ui/ | SwaggerUI route (v3)


# **GetDev**
> string GetDev(ctx, optional)
SwaggerUI route (v2)

ESI dev web UI. Note that this is a legacy (v2) SwaggerUI route, which will someday be removed. You can get mostly the same functionality by going to https://esi.evetech.net/ui/?version=dev

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

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **GetLatest**
> string GetLatest(ctx, optional)
SwaggerUI route (v2)

ESI latest web UI. Note that this is a legacy (v2) SwaggerUI route, which will someday be removed. You can get mostly the same functionality by going to https://esi.evetech.net/ui/?version=latest

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

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLegacy**
> string GetLegacy(ctx, optional)
SwaggerUI route (v2)

ESI legacy web UI. Note that this is a legacy (v2) SwaggerUI route, which will someday be removed. You can get mostly the same functionality by going to https://esi.evetech.net/ui/?version=legacy

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

