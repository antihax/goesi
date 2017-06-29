# \PlanetaryInteractionApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdPlanets**](PlanetaryInteractionApi.md#GetCharactersCharacterIdPlanets) | **Get** /characters/{character_id}/planets/ | Get colonies
[**GetCharactersCharacterIdPlanetsPlanetId**](PlanetaryInteractionApi.md#GetCharactersCharacterIdPlanetsPlanetId) | **Get** /characters/{character_id}/planets/{planet_id}/ | Get colony layout
[**GetUniverseSchematicsSchematicId**](PlanetaryInteractionApi.md#GetUniverseSchematicsSchematicId) | **Get** /universe/schematics/{schematic_id}/ | Get schematic information


# **GetCharactersCharacterIdPlanets**
> []GetCharactersCharacterIdPlanets200Ok GetCharactersCharacterIdPlanets(ctx, characterId, optional)
Get colonies

Returns a list of all planetary colonies owned by a character.  --- Alternate route: `/legacy/characters/{character_id}/planets/`  Alternate route: `/latest/characters/{character_id}/planets/`  Alternate route: `/dev/characters/{character_id}/planets/`  --- This route is cached for up to 600 seconds

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

[**[]GetCharactersCharacterIdPlanets200Ok**](get_characters_character_id_planets_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdPlanetsPlanetId**
> GetCharactersCharacterIdPlanetsPlanetIdOk GetCharactersCharacterIdPlanetsPlanetId(ctx, planetId, characterId, optional)
Get colony layout

Returns full details on the layout of a single planetary colony, including links, pins and routes. Note: Planetary information is only recalculated when the colony is viewed through the client. Information will not update until this criteria is met.  --- Alternate route: `/legacy/characters/{character_id}/planets/{planet_id}/`  --- This route is cached for up to 600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **planetId** | **int32**| Planet id of the target planet | 
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planetId** | **int32**| Planet id of the target planet | 
 **characterId** | **int32**| An EVE character ID | 
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

# **GetUniverseSchematicsSchematicId**
> GetUniverseSchematicsSchematicIdOk GetUniverseSchematicsSchematicId(schematicId, optional)
Get schematic information

Get information on a planetary factory schematic  --- Alternate route: `/legacy/universe/schematics/{schematic_id}/`  Alternate route: `/latest/universe/schematics/{schematic_id}/`  Alternate route: `/dev/universe/schematics/{schematic_id}/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **schematicId** | **int32**| A PI schematic ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schematicId** | **int32**| A PI schematic ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetUniverseSchematicsSchematicIdOk**](get_universe_schematics_schematic_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

