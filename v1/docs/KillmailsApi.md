# \KillmailsApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdKillmailsRecent**](KillmailsApi.md#GetCharactersCharacterIdKillmailsRecent) | **Get** /characters/{character_id}/killmails/recent/ | List kills and losses
[**GetKillmailsKillmailIdKillmailHash**](KillmailsApi.md#GetKillmailsKillmailIdKillmailHash) | **Get** /killmails/{killmail_id}/{killmail_hash}/ | Get a single killmail


# **GetCharactersCharacterIdKillmailsRecent**
> []GetCharactersCharacterIdKillmailsRecent200Ok GetCharactersCharacterIdKillmailsRecent(ctx, characterId, optional)
List kills and losses

Return a list of character's recent kills and losses  --- Alternate route: `/legacy/characters/{character_id}/killmails/recent/`  Alternate route: `/latest/characters/{character_id}/killmails/recent/`  Alternate route: `/dev/characters/{character_id}/killmails/recent/`  --- This route is cached for up to 120 seconds

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
 **maxCount** | **int32**| How many killmails to return at maximum | [default to 50]
 **maxKillId** | **int32**| Only return killmails with ID smaller than this.  | 
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdKillmailsRecent200Ok**](get_characters_character_id_killmails_recent_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKillmailsKillmailIdKillmailHash**
> GetKillmailsKillmailIdKillmailHashOk GetKillmailsKillmailIdKillmailHash(killmailHash, killmailId, optional)
Get a single killmail

Return a single killmail from its ID and hash  --- Alternate route: `/legacy/killmails/{killmail_id}/{killmail_hash}/`  Alternate route: `/latest/killmails/{killmail_id}/{killmail_hash}/`  Alternate route: `/dev/killmails/{killmail_id}/{killmail_hash}/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **killmailHash** | **string**| The killmail hash for verification | 
  **killmailId** | **int32**| The killmail ID to be queried | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **killmailHash** | **string**| The killmail hash for verification | 
 **killmailId** | **int32**| The killmail ID to be queried | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetKillmailsKillmailIdKillmailHashOk**](get_killmails_killmail_id_killmail_hash_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

