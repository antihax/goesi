# \FittingsApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCharactersCharacterIdFittingsFittingId**](FittingsApi.md#DeleteCharactersCharacterIdFittingsFittingId) | **Delete** /characters/{character_id}/fittings/{fitting_id}/ | Delete fitting
[**GetCharactersCharacterIdFittings**](FittingsApi.md#GetCharactersCharacterIdFittings) | **Get** /characters/{character_id}/fittings/ | Get fittings
[**PostCharactersCharacterIdFittings**](FittingsApi.md#PostCharactersCharacterIdFittings) | **Post** /characters/{character_id}/fittings/ | Create fitting


# **DeleteCharactersCharacterIdFittingsFittingId**
> DeleteCharactersCharacterIdFittingsFittingId(ctx, fittingId, characterId, optional)
Delete fitting

Delete a fitting from a character  --- Alternate route: `/legacy/characters/{character_id}/fittings/{fitting_id}/`  Alternate route: `/latest/characters/{character_id}/fittings/{fitting_id}/`  Alternate route: `/dev/characters/{character_id}/fittings/{fitting_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fittingId** | **int32**| ID for a fitting of this character | 
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fittingId** | **int32**| ID for a fitting of this character | 
 **characterId** | **int32**| An EVE character ID | 
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

# **GetCharactersCharacterIdFittings**
> []GetCharactersCharacterIdFittings200Ok GetCharactersCharacterIdFittings(ctx, characterId, optional)
Get fittings

Return fittings of a character  --- Alternate route: `/legacy/characters/{character_id}/fittings/`  Alternate route: `/latest/characters/{character_id}/fittings/`  Alternate route: `/dev/characters/{character_id}/fittings/`  --- This route is cached for up to 300 seconds

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

[**[]GetCharactersCharacterIdFittings200Ok**](get_characters_character_id_fittings_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdFittings**
> PostCharactersCharacterIdFittingsCreated PostCharactersCharacterIdFittings(ctx, characterId, optional)
Create fitting

Save a new fitting for a character  --- Alternate route: `/legacy/characters/{character_id}/fittings/`  Alternate route: `/latest/characters/{character_id}/fittings/`  Alternate route: `/dev/characters/{character_id}/fittings/` 

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
 **fitting** | [**PostCharactersCharacterIdFittingsFitting**](PostCharactersCharacterIdFittingsFitting.md)| Details about the new fitting | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**PostCharactersCharacterIdFittingsCreated**](post_characters_character_id_fittings_created.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

