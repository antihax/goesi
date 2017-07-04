# \PlanetaryInteractionApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdPlanetsPlanetId**](PlanetaryInteractionApi.md#GetCharactersCharacterIdPlanetsPlanetId) | **Get** /characters/{character_id}/planets/{planet_id}/ | Get colony layout


# **GetCharactersCharacterIdPlanetsPlanetId**
> GetCharactersCharacterIdPlanetsPlanetIdOk GetCharactersCharacterIdPlanetsPlanetId(ctx, characterId, planetId, optional)
Get colony layout

Returns full details on the layout of a single planetary colony, including links, pins and routes. Note: Planetary information is only recalculated when the colony is viewed through the client. Information will not update until this criteria is met.  --- Alternate route: `/latest/characters/{character_id}/planets/{planet_id}/`  --- This route is cached for up to 600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
  **planetId** | **int32**| Planet id of the target planet | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **planetId** | **int32**| Planet id of the target planet | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetCharactersCharacterIdPlanetsPlanetIdOk**](get_characters_character_id_planets_planet_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

