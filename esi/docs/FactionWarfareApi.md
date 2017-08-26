# \FactionWarfareApi

All URIs are relative to *https://esi.tech.ccp.is*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFwLeaderboards**](FactionWarfareApi.md#GetFwLeaderboards) | **Get** /v1/fw/leaderboards/ | List of the top factions in faction warfare
[**GetFwLeaderboardsCharacters**](FactionWarfareApi.md#GetFwLeaderboardsCharacters) | **Get** /v1/fw/leaderboards/characters/ | List of the top pilots in faction warfare
[**GetFwLeaderboardsCorporations**](FactionWarfareApi.md#GetFwLeaderboardsCorporations) | **Get** /v1/fw/leaderboards/corporations/ | List of the top corporations in faction warfare
[**GetFwStats**](FactionWarfareApi.md#GetFwStats) | **Get** /v1/fw/stats/ | An overview of statistics about factions involved in faction warfare
[**GetFwSystems**](FactionWarfareApi.md#GetFwSystems) | **Get** /v1/fw/systems/ | Ownership of faction warfare systems
[**GetFwWars**](FactionWarfareApi.md#GetFwWars) | **Get** /v1/fw/wars/ | Data about which NPC factions are at war


# **GetFwLeaderboards**
> GetFwLeaderboardsOk GetFwLeaderboards(optional)
List of the top factions in faction warfare

Top 4 leaderboard of factions for kills and victory points separated by total, last week and yesterday.  ---  This route expires daily at 11:05

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

[**GetFwLeaderboardsOk**](get_fw_leaderboards_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFwLeaderboardsCharacters**
> GetFwLeaderboardsCharactersOk GetFwLeaderboardsCharacters(optional)
List of the top pilots in faction warfare

Top 100 leaderboard of pilots for kills and victory points separated by total, last week and yesterday.  ---  This route expires daily at 11:05

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

[**GetFwLeaderboardsCharactersOk**](get_fw_leaderboards_characters_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFwLeaderboardsCorporations**
> GetFwLeaderboardsCorporationsOk GetFwLeaderboardsCorporations(optional)
List of the top corporations in faction warfare

Top 10 leaderboard of corporations for kills and victory points separated by total, last week and yesterday.  ---  This route expires daily at 11:05

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

[**GetFwLeaderboardsCorporationsOk**](get_fw_leaderboards_corporations_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFwStats**
> []GetFwStats200Ok GetFwStats(optional)
An overview of statistics about factions involved in faction warfare

Statistical overviews of factions involved in faction warfare  ---  This route expires daily at 11:05

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

[**[]GetFwStats200Ok**](get_fw_stats_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFwSystems**
> []GetFwSystems200Ok GetFwSystems(optional)
Ownership of faction warfare systems

An overview of the current ownership of faction warfare solar systems  ---  This route is cached for up to 3600 seconds

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

[**[]GetFwSystems200Ok**](get_fw_systems_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFwWars**
> []GetFwWars200Ok GetFwWars(optional)
Data about which NPC factions are at war

Data about which NPC factions are at war  ---  This route expires daily at 11:05

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

[**[]GetFwWars200Ok**](get_fw_wars_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

