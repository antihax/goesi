# \SearchApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdSearch**](SearchApi.md#GetCharactersCharacterIdSearch) | **Get** /characters/{character_id}/search/ | Search on a string
[**GetSearch**](SearchApi.md#GetSearch) | **Get** /search/ | Search on a string


# **GetCharactersCharacterIdSearch**
> GetCharactersCharacterIdSearchOk GetCharactersCharacterIdSearch(ctx, characterId, search, categories, optional)
Search on a string

Search for entities that match a given sub-string.  ---  Alternate route: `/latest/characters/{character_id}/search/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
  **search** | **string**| The string to search on | 
  **categories** | [**[]string**](string.md)| Type of entities to search for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **search** | **string**| The string to search on | 
 **categories** | [**[]string**](string.md)| Type of entities to search for | 
 **language** | **string**| Search locale | [default to en-us]
 **strict** | **bool**| Whether the search should be a strict match | [default to false]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetCharactersCharacterIdSearchOk**](get_characters_character_id_search_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSearch**
> GetSearchOk GetSearch(search, categories, optional)
Search on a string

Search for entities that match a given sub-string.  ---  Alternate route: `/dev/search/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **search** | **string**| The string to search on | 
  **categories** | [**[]string**](string.md)| Type of entities to search for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string**| The string to search on | 
 **categories** | [**[]string**](string.md)| Type of entities to search for | 
 **language** | **string**| Search locale | [default to en-us]
 **strict** | **bool**| Whether the search should be a strict match | [default to false]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetSearchOk**](get_search_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

