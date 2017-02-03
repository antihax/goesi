# \SkillsApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdSkillqueue**](SkillsApi.md#GetCharactersCharacterIdSkillqueue) | **Get** /characters/{character_id}/skillqueue/ | Get character&#39;s skill queue
[**GetCharactersCharacterIdSkills**](SkillsApi.md#GetCharactersCharacterIdSkills) | **Get** /characters/{character_id}/skills/ | Get character skills


# **GetCharactersCharacterIdSkillqueue**
> []GetCharactersCharacterIdSkillqueue200Ok GetCharactersCharacterIdSkillqueue(ctx, characterId, optional)
Get character's skill queue

List the configured skill queue for the given character  ---  Alternate route: `/legacy/characters/{character_id}/skillqueue/`  Alternate route: `/latest/characters/{character_id}/skillqueue/`  Alternate route: `/dev/characters/{character_id}/skillqueue/`   ---  This route is cached for up to 120 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| Character id of the target character | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| Character id of the target character | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**[]GetCharactersCharacterIdSkillqueue200Ok**](get_characters_character_id_skillqueue_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdSkills**
> []GetCharactersCharacterIdSkills200Ok GetCharactersCharacterIdSkills(ctx, characterId, optional)
Get character skills

List all trained skills for the given character  ---  Alternate route: `/legacy/characters/{character_id}/skills/`   ---  This route is cached for up to 120 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**[]GetCharactersCharacterIdSkills200Ok**](get_characters_character_id_skills_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

