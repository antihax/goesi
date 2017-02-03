# \CharacterApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdCorporationhistory**](CharacterApi.md#GetCharactersCharacterIdCorporationhistory) | **Get** /characters/{character_id}/corporationhistory/ | Get corporation history
[**GetCharactersCharacterIdPortrait**](CharacterApi.md#GetCharactersCharacterIdPortrait) | **Get** /characters/{character_id}/portrait/ | Get character portraits
[**GetCharactersNames**](CharacterApi.md#GetCharactersNames) | **Get** /characters/names/ | Get character names


# **GetCharactersCharacterIdCorporationhistory**
> []GetCharactersCharacterIdCorporationhistory200Ok GetCharactersCharacterIdCorporationhistory(characterId, optional)
Get corporation history

Get a list of all the corporations a character has been a member of  ---  Alternate route: `/legacy/characters/{character_id}/corporationhistory/`  Alternate route: `/latest/characters/{character_id}/corporationhistory/`  Alternate route: `/dev/characters/{character_id}/corporationhistory/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**[]GetCharactersCharacterIdCorporationhistory200Ok**](get_characters_character_id_corporationhistory_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdPortrait**
> GetCharactersCharacterIdPortraitOk GetCharactersCharacterIdPortrait(characterId, optional)
Get character portraits

Get portrait urls for a character  ---  Alternate route: `/legacy/characters/{character_id}/portrait/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**GetCharactersCharacterIdPortraitOk**](get_characters_character_id_portrait_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersNames**
> []GetCharactersNames200Ok GetCharactersNames(characterIds, optional)
Get character names

Resolve a set of character IDs to character names  ---  Alternate route: `/legacy/characters/names/`  Alternate route: `/latest/characters/names/`  Alternate route: `/dev/characters/names/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterIds** | [**[]int64**](int64.md)| A comma separated list of character IDs | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterIds** | [**[]int64**](int64.md)| A comma separated list of character IDs | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**[]GetCharactersNames200Ok**](get_characters_names_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

