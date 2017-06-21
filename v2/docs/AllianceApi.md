# \AllianceApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlliancesAllianceId**](AllianceApi.md#GetAlliancesAllianceId) | **Get** /alliances/{alliance_id}/ | Get alliance information
[**GetAlliancesNames**](AllianceApi.md#GetAlliancesNames) | **Get** /alliances/names/ | Get alliance names


# **GetAlliancesAllianceId**
> GetAlliancesAllianceIdOk GetAlliancesAllianceId(allianceId, optional)
Get alliance information

Public information about an alliance  ---  Alternate route: `/latest/alliances/{alliance_id}/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **allianceId** | **int32**| An Eve alliance ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceId** | **int32**| An Eve alliance ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetAlliancesAllianceIdOk**](get_alliances_alliance_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlliancesNames**
> []GetAlliancesNames200Ok GetAlliancesNames(allianceIds, optional)
Get alliance names

Resolve a set of alliance IDs to alliance names  ---  Alternate route: `/dev/alliances/names/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **allianceIds** | [**[]int32**](int32.md)| A comma separated list of alliance IDs | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceIds** | [**[]int32**](int32.md)| A comma separated list of alliance IDs | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetAlliancesNames200Ok**](get_alliances_names_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

