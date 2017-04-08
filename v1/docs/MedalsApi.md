# \MedalsApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostFanfestToast**](MedalsApi.md#PostFanfestToast) | **Post** /fanfest/toast/ | Nothing to see here, move along


# **PostFanfestToast**
> PostFanfestToast(ctx, details, optional)
Nothing to see here, move along

These are not the endpoints you are looking for  ---  Alternate route: `/legacy/fanfest/toast/`  Alternate route: `/latest/fanfest/toast/`  Alternate route: `/dev/fanfest/toast/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **details** | [**PostFanfestToastDetails**](PostFanfestToastDetails.md)| Award details | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **details** | [**PostFanfestToastDetails**](PostFanfestToastDetails.md)| Award details | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

