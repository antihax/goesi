# \UserInterfaceApi

All URIs are relative to *https://esi.tech.ccp.is*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostUiAutopilotWaypoint**](UserInterfaceApi.md#PostUiAutopilotWaypoint) | **Post** /v2/ui/autopilot/waypoint/ | Set Autopilot Waypoint
[**PostUiOpenwindowContract**](UserInterfaceApi.md#PostUiOpenwindowContract) | **Post** /v1/ui/openwindow/contract/ | Open Contract Window
[**PostUiOpenwindowInformation**](UserInterfaceApi.md#PostUiOpenwindowInformation) | **Post** /v1/ui/openwindow/information/ | Open Information Window
[**PostUiOpenwindowMarketdetails**](UserInterfaceApi.md#PostUiOpenwindowMarketdetails) | **Post** /v1/ui/openwindow/marketdetails/ | Open Market Details
[**PostUiOpenwindowNewmail**](UserInterfaceApi.md#PostUiOpenwindowNewmail) | **Post** /v1/ui/openwindow/newmail/ | Open New Mail Window


# **PostUiAutopilotWaypoint**
> PostUiAutopilotWaypoint(ctx, addToBeginning, clearOtherWaypoints, destinationId, optional)
Set Autopilot Waypoint

Set a solar system as autopilot waypoint  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **addToBeginning** | **bool**| Whether this solar system should be added to the beginning of all waypoints | [default to false]
  **clearOtherWaypoints** | **bool**| Whether clean other waypoints beforing adding this one | [default to false]
  **destinationId** | **int64**| The destination to travel to, can be solar system, station or structure&#39;s id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addToBeginning** | **bool**| Whether this solar system should be added to the beginning of all waypoints | [default to false]
 **clearOtherWaypoints** | **bool**| Whether clean other waypoints beforing adding this one | [default to false]
 **destinationId** | **int64**| The destination to travel to, can be solar system, station or structure&#39;s id | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
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

# **PostUiOpenwindowContract**
> PostUiOpenwindowContract(ctx, contractId, optional)
Open Contract Window

Open the contract window inside the client  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **contractId** | **int32**| The contract to open | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **int32**| The contract to open | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
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

# **PostUiOpenwindowInformation**
> PostUiOpenwindowInformation(ctx, targetId, optional)
Open Information Window

Open the information window for a character, corporation or alliance inside the client  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **targetId** | **int32**| The target to open | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **int32**| The target to open | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
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

# **PostUiOpenwindowMarketdetails**
> PostUiOpenwindowMarketdetails(ctx, typeId, optional)
Open Market Details

Open the market details window for a specific typeID inside the client  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **typeId** | **int32**| The item type to open in market window | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeId** | **int32**| The item type to open in market window | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
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

# **PostUiOpenwindowNewmail**
> PostUiOpenwindowNewmail(ctx, newMail, optional)
Open New Mail Window

Open the New Mail window, according to settings from the request if applicable  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **newMail** | [**PostUiOpenwindowNewmailNewMail**](PostUiOpenwindowNewmailNewMail.md)| The details of mail to create | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newMail** | [**PostUiOpenwindowNewmailNewMail**](PostUiOpenwindowNewmailNewMail.md)| The details of mail to create | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
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

