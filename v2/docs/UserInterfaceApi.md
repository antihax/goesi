# \UserInterfaceApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostUiAutopilotWaypoint**](UserInterfaceApi.md#PostUiAutopilotWaypoint) | **Post** /ui/autopilot/waypoint/ | Set Autopilot Waypoint


# **PostUiAutopilotWaypoint**
> PostUiAutopilotWaypoint(ctx, destinationId, clearOtherWaypoints, addToBeginning, optional)
Set Autopilot Waypoint

Set a solar system as autopilot waypoint  ---  Alternate route: `/latest/ui/autopilot/waypoint/`  Alternate route: `/dev/ui/autopilot/waypoint/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **destinationId** | **int64**| The destination to travel to, can be solar system, station or structure&#39;s id | 
  **clearOtherWaypoints** | **bool**| Whether clean other waypoints beforing adding this one | [default to false]
  **addToBeginning** | **bool**| Whether this solar system should be added to the beginning of all waypoints | [default to false]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationId** | **int64**| The destination to travel to, can be solar system, station or structure&#39;s id | 
 **clearOtherWaypoints** | **bool**| Whether clean other waypoints beforing adding this one | [default to false]
 **addToBeginning** | **bool**| Whether this solar system should be added to the beginning of all waypoints | [default to false]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

