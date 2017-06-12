# \LoyaltyApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdLoyaltyPoints**](LoyaltyApi.md#GetCharactersCharacterIdLoyaltyPoints) | **Get** /characters/{character_id}/loyalty/points/ | Get loyalty points
[**GetLoyaltyStoresCorporationIdOffers**](LoyaltyApi.md#GetLoyaltyStoresCorporationIdOffers) | **Get** /loyalty/stores/{corporation_id}/offers/ | List loyalty store offers


# **GetCharactersCharacterIdLoyaltyPoints**
> []GetCharactersCharacterIdLoyaltyPoints200Ok GetCharactersCharacterIdLoyaltyPoints(ctx, characterId, optional)
Get loyalty points

Return a list of loyalty points for all corporations the character has worked for  ---  Alternate route: `/legacy/characters/{character_id}/loyalty/points/`  Alternate route: `/latest/characters/{character_id}/loyalty/points/`  Alternate route: `/dev/characters/{character_id}/loyalty/points/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| ID for a character | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| ID for a character | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdLoyaltyPoints200Ok**](get_characters_character_id_loyalty_points_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoyaltyStoresCorporationIdOffers**
> []GetLoyaltyStoresCorporationIdOffers200Ok GetLoyaltyStoresCorporationIdOffers(corporationId, optional)
List loyalty store offers

Return a list of offers from a specific corporation's loyalty store  ---  Alternate route: `/legacy/loyalty/stores/{corporation_id}/offers/`  Alternate route: `/latest/loyalty/stores/{corporation_id}/offers/`  Alternate route: `/dev/loyalty/stores/{corporation_id}/offers/`   ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **corporationId** | **int32**| ID of a corporation | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| ID of a corporation | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetLoyaltyStoresCorporationIdOffers200Ok**](get_loyalty_stores_corporation_id_offers_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

