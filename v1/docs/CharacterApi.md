# \CharacterApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdAgentsResearch**](CharacterApi.md#GetCharactersCharacterIdAgentsResearch) | **Get** /characters/{character_id}/agents_research/ | Get agents research
[**GetCharactersCharacterIdBlueprints**](CharacterApi.md#GetCharactersCharacterIdBlueprints) | **Get** /characters/{character_id}/blueprints/ | Get blueprints
[**GetCharactersCharacterIdChatChannels**](CharacterApi.md#GetCharactersCharacterIdChatChannels) | **Get** /characters/{character_id}/chat_channels/ | Get chat channels
[**GetCharactersCharacterIdCorporationhistory**](CharacterApi.md#GetCharactersCharacterIdCorporationhistory) | **Get** /characters/{character_id}/corporationhistory/ | Get corporation history
[**GetCharactersCharacterIdMedals**](CharacterApi.md#GetCharactersCharacterIdMedals) | **Get** /characters/{character_id}/medals/ | Get medals
[**GetCharactersCharacterIdPortrait**](CharacterApi.md#GetCharactersCharacterIdPortrait) | **Get** /characters/{character_id}/portrait/ | Get character portraits
[**GetCharactersCharacterIdRoles**](CharacterApi.md#GetCharactersCharacterIdRoles) | **Get** /characters/{character_id}/roles/ | Get character corporation roles
[**GetCharactersCharacterIdStandings**](CharacterApi.md#GetCharactersCharacterIdStandings) | **Get** /characters/{character_id}/standings/ | Get standings
[**GetCharactersNames**](CharacterApi.md#GetCharactersNames) | **Get** /characters/names/ | Get character names
[**PostCharactersAffiliation**](CharacterApi.md#PostCharactersAffiliation) | **Post** /characters/affiliation/ | Character affiliation


# **GetCharactersCharacterIdAgentsResearch**
> []GetCharactersCharacterIdAgentsResearch200Ok GetCharactersCharacterIdAgentsResearch(ctx, characterId, optional)
Get agents research

Return a list of agents research information for a character. The formula for finding the current research points with an agent is: currentPoints = remainderPoints + pointsPerDay * days(currentTime - researchStartDate)  --- Alternate route: `/legacy/characters/{character_id}/agents_research/`  Alternate route: `/latest/characters/{character_id}/agents_research/`  Alternate route: `/dev/characters/{character_id}/agents_research/`  --- This route is cached for up to 3600 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdAgentsResearch200Ok**](get_characters_character_id_agents_research_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdBlueprints**
> []GetCharactersCharacterIdBlueprints200Ok GetCharactersCharacterIdBlueprints(ctx, characterId, optional)
Get blueprints

Return a list of blueprints the character has  --- Alternate route: `/legacy/characters/{character_id}/blueprints/`  Alternate route: `/latest/characters/{character_id}/blueprints/`  Alternate route: `/dev/characters/{character_id}/blueprints/`  --- This route is cached for up to 3600 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdBlueprints200Ok**](get_characters_character_id_blueprints_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdChatChannels**
> []GetCharactersCharacterIdChatChannels200Ok GetCharactersCharacterIdChatChannels(ctx, characterId, optional)
Get chat channels

Return chat channels that a character is the owner or an operator of  --- Alternate route: `/legacy/characters/{character_id}/chat_channels/`  Alternate route: `/latest/characters/{character_id}/chat_channels/`  Alternate route: `/dev/characters/{character_id}/chat_channels/`  --- This route is cached for up to 300 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdChatChannels200Ok**](get_characters_character_id_chat_channels_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdCorporationhistory**
> []GetCharactersCharacterIdCorporationhistory200Ok GetCharactersCharacterIdCorporationhistory(characterId, optional)
Get corporation history

Get a list of all the corporations a character has been a member of  --- Alternate route: `/legacy/characters/{character_id}/corporationhistory/`  Alternate route: `/latest/characters/{character_id}/corporationhistory/`  Alternate route: `/dev/characters/{character_id}/corporationhistory/`  --- This route is cached for up to 3600 seconds

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
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdCorporationhistory200Ok**](get_characters_character_id_corporationhistory_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdMedals**
> []GetCharactersCharacterIdMedals200Ok GetCharactersCharacterIdMedals(ctx, characterId, optional)
Get medals

Return a list of medals the character has  --- Alternate route: `/legacy/characters/{character_id}/medals/`  Alternate route: `/latest/characters/{character_id}/medals/`  Alternate route: `/dev/characters/{character_id}/medals/`  --- This route is cached for up to 3600 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdMedals200Ok**](get_characters_character_id_medals_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdPortrait**
> GetCharactersCharacterIdPortraitOk GetCharactersCharacterIdPortrait(characterId, optional)
Get character portraits

Get portrait urls for a character  --- Alternate route: `/legacy/characters/{character_id}/portrait/`  --- This route is cached for up to 3600 seconds

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
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetCharactersCharacterIdPortraitOk**](get_characters_character_id_portrait_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdRoles**
> []string GetCharactersCharacterIdRoles(ctx, characterId, optional)
Get character corporation roles

Returns a character's corporation roles  --- Alternate route: `/legacy/characters/{character_id}/roles/`  Alternate route: `/latest/characters/{character_id}/roles/`  Alternate route: `/dev/characters/{character_id}/roles/`  --- This route is cached for up to 3600 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

**[]string**

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdStandings**
> []GetCharactersCharacterIdStandings200Ok GetCharactersCharacterIdStandings(ctx, characterId, optional)
Get standings

Return character standings from agents, NPC corporations, and factions  --- Alternate route: `/legacy/characters/{character_id}/standings/`  Alternate route: `/latest/characters/{character_id}/standings/`  Alternate route: `/dev/characters/{character_id}/standings/`  --- This route is cached for up to 3600 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdStandings200Ok**](get_characters_character_id_standings_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersNames**
> []GetCharactersNames200Ok GetCharactersNames(characterIds, optional)
Get character names

Resolve a set of character IDs to character names  --- Alternate route: `/legacy/characters/names/`  Alternate route: `/latest/characters/names/`  Alternate route: `/dev/characters/names/`  --- This route is cached for up to 3600 seconds

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
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersNames200Ok**](get_characters_names_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersAffiliation**
> []PostCharactersAffiliation200Ok PostCharactersAffiliation(characters, optional)
Character affiliation

Bulk lookup of character IDs to corporation, alliance and faction  --- Alternate route: `/legacy/characters/affiliation/`  Alternate route: `/latest/characters/affiliation/`  Alternate route: `/dev/characters/affiliation/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characters** | **[]int32**| The character IDs to fetch affiliations for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characters** | **[]int32**| The character IDs to fetch affiliations for | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]PostCharactersAffiliation200Ok**](post_characters_affiliation_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

