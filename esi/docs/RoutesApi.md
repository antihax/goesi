# \RoutesApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRouteOriginDestination**](RoutesApi.md#GetRouteOriginDestination) | **Get** /v1/route/{origin}/{destination}/ | Get route


# **GetRouteOriginDestination**
> []int32 GetRouteOriginDestination(ctx, destination, origin, optional)
Get route

Get the systems between origin and destination  ---  This route is cached for up to 86400 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **destination** | **int32**| destination solar system ID | 
  **origin** | **int32**| origin solar system ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destination** | **int32**| destination solar system ID | 
 **origin** | **int32**| origin solar system ID | 
 **avoid** | [**[]int32**](int32.md)| avoid solar system ID(s) | 
 **connections** | [**[][]int32**]([]int32.md)| connected solar system pairs | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **flag** | **string**| route security preference | [default to shortest]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

