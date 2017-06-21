# \MarketApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdOrders**](MarketApi.md#GetCharactersCharacterIdOrders) | **Get** /characters/{character_id}/orders/ | List orders from a character
[**GetMarketsGroups**](MarketApi.md#GetMarketsGroups) | **Get** /markets/groups/ | Get item groups
[**GetMarketsGroupsMarketGroupId**](MarketApi.md#GetMarketsGroupsMarketGroupId) | **Get** /markets/groups/{market_group_id}/ | Get item group information
[**GetMarketsPrices**](MarketApi.md#GetMarketsPrices) | **Get** /markets/prices/ | List market prices
[**GetMarketsRegionIdHistory**](MarketApi.md#GetMarketsRegionIdHistory) | **Get** /markets/{region_id}/history/ | List historical market statistics in a region
[**GetMarketsRegionIdOrders**](MarketApi.md#GetMarketsRegionIdOrders) | **Get** /markets/{region_id}/orders/ | List orders in a region
[**GetMarketsStructuresStructureId**](MarketApi.md#GetMarketsStructuresStructureId) | **Get** /markets/structures/{structure_id}/ | List orders in a structure


# **GetCharactersCharacterIdOrders**
> []GetCharactersCharacterIdOrders200Ok GetCharactersCharacterIdOrders(ctx, characterId, optional)
List orders from a character

List market orders placed by a character  ---  Alternate route: `/legacy/characters/{character_id}/orders/`  Alternate route: `/latest/characters/{character_id}/orders/`  Alternate route: `/dev/characters/{character_id}/orders/`   ---  This route is cached for up to 3600 seconds

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
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdOrders200Ok**](get_characters_character_id_orders_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsGroups**
> []int32 GetMarketsGroups(optional)
Get item groups

Get a list of item groups  ---  Alternate route: `/legacy/markets/groups/`  Alternate route: `/latest/markets/groups/`  Alternate route: `/dev/markets/groups/`   ---  This route expires daily at 11:05

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

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsGroupsMarketGroupId**
> GetMarketsGroupsMarketGroupIdOk GetMarketsGroupsMarketGroupId(marketGroupId, optional)
Get item group information

Get information on an item group  ---  Alternate route: `/legacy/markets/groups/{market_group_id}/`  Alternate route: `/latest/markets/groups/{market_group_id}/`  Alternate route: `/dev/markets/groups/{market_group_id}/`   ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **marketGroupId** | **int32**| An Eve item group ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketGroupId** | **int32**| An Eve item group ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **language** | **string**| Language to use in the response | [default to en-us]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetMarketsGroupsMarketGroupIdOk**](get_markets_groups_market_group_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsPrices**
> []GetMarketsPrices200Ok GetMarketsPrices(optional)
List market prices

Return a list of prices  ---  Alternate route: `/legacy/markets/prices/`  Alternate route: `/latest/markets/prices/`  Alternate route: `/dev/markets/prices/`   ---  This route is cached for up to 3600 seconds

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

[**[]GetMarketsPrices200Ok**](get_markets_prices_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsRegionIdHistory**
> []GetMarketsRegionIdHistory200Ok GetMarketsRegionIdHistory(regionId, typeId, optional)
List historical market statistics in a region

Return a list of historical market statistics for the specified type in a region  ---  Alternate route: `/legacy/markets/{region_id}/history/`  Alternate route: `/latest/markets/{region_id}/history/`  Alternate route: `/dev/markets/{region_id}/history/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **regionId** | **int32**| Return statistics in this region | 
  **typeId** | **int32**| Return statistics for this type | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **int32**| Return statistics in this region | 
 **typeId** | **int32**| Return statistics for this type | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetMarketsRegionIdHistory200Ok**](get_markets_region_id_history_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsRegionIdOrders**
> []GetMarketsRegionIdOrders200Ok GetMarketsRegionIdOrders(orderType, regionId, optional)
List orders in a region

Return a list of orders in a region  ---  Alternate route: `/legacy/markets/{region_id}/orders/`  Alternate route: `/latest/markets/{region_id}/orders/`  Alternate route: `/dev/markets/{region_id}/orders/`   ---  This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **orderType** | **string**| Filter buy/sell orders, return all orders by default. If you query without type_id, we always return both buy and sell orders. | [default to all]
  **regionId** | **int32**| Return orders in this region | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderType** | **string**| Filter buy/sell orders, return all orders by default. If you query without type_id, we always return both buy and sell orders. | [default to all]
 **regionId** | **int32**| Return orders in this region | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **page** | **int32**| Which page to query, only used for querying without type_id. Starting at 1 | [default to 1]
 **typeId** | **int32**| Return orders only for this type | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetMarketsRegionIdOrders200Ok**](get_markets_region_id_orders_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsStructuresStructureId**
> []GetMarketsStructuresStructureId200Ok GetMarketsStructuresStructureId(ctx, structureId, optional)
List orders in a structure

Return all orders in a structure  ---  Alternate route: `/legacy/markets/structures/{structure_id}/`  Alternate route: `/latest/markets/structures/{structure_id}/`  Alternate route: `/dev/markets/structures/{structure_id}/`   ---  This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **structureId** | **int64**| Return orders in this structure | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **structureId** | **int64**| Return orders in this structure | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **page** | **int32**| Which page to query, starting at 1 | [default to 1]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetMarketsStructuresStructureId200Ok**](get_markets_structures_structure_id_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

