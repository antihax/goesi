# \AllianceApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlliances**](AllianceApi.md#GetAlliances) | **Get** /alliances/ | List all alliances
[**GetAlliancesAllianceId**](AllianceApi.md#GetAlliancesAllianceId) | **Get** /alliances/{alliance_id}/ | Get alliance information
[**GetAlliancesAllianceIdCorporations**](AllianceApi.md#GetAlliancesAllianceIdCorporations) | **Get** /alliances/{alliance_id}/corporations/ | List alliance&#39;s corporations
[**GetAlliancesAllianceIdIcons**](AllianceApi.md#GetAlliancesAllianceIdIcons) | **Get** /alliances/{alliance_id}/icons/ | Get alliance icon
[**GetAlliancesNames**](AllianceApi.md#GetAlliancesNames) | **Get** /alliances/names/ | Get alliance names


# **GetAlliances**
> []int32 GetAlliances(optional)
List all alliances

List all active player alliances  --- Alternate route: `/latest/alliances/`  Alternate route: `/legacy/alliances/`  Alternate route: `/dev/alliances/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlliancesAllianceId**
> GetAlliancesAllianceIdOk GetAlliancesAllianceId(allianceId, optional)
Get alliance information

Public information about an alliance  --- Alternate route: `/dev/alliances/{alliance_id}/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **allianceId** | **int32**| An EVE alliance ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceId** | **int32**| An EVE alliance ID | 
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

# **GetAlliancesAllianceIdCorporations**
> []int32 GetAlliancesAllianceIdCorporations(allianceId, optional)
List alliance's corporations

List all current member corporations of an alliance  --- Alternate route: `/latest/alliances/{alliance_id}/corporations/`  Alternate route: `/legacy/alliances/{alliance_id}/corporations/`  Alternate route: `/dev/alliances/{alliance_id}/corporations/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **allianceId** | **int32**| An EVE alliance ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceId** | **int32**| An EVE alliance ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlliancesAllianceIdIcons**
> GetAlliancesAllianceIdIconsOk GetAlliancesAllianceIdIcons(allianceId, optional)
Get alliance icon

Get the icon urls for a alliance  --- Alternate route: `/legacy/alliances/{alliance_id}/icons/`  Alternate route: `/latest/alliances/{alliance_id}/icons/`  Alternate route: `/dev/alliances/{alliance_id}/icons/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **allianceId** | **int32**| An EVE alliance ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceId** | **int32**| An EVE alliance ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetAlliancesAllianceIdIconsOk**](get_alliances_alliance_id_icons_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlliancesNames**
> []GetAlliancesNames200Ok GetAlliancesNames(allianceIds, optional)
Get alliance names

Resolve a set of alliance IDs to alliance names  --- Alternate route: `/legacy/alliances/names/`  Alternate route: `/latest/alliances/names/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **allianceIds** | [**[]int64**](int64.md)| A comma separated list of alliance IDs | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceIds** | [**[]int64**](int64.md)| A comma separated list of alliance IDs | 
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

