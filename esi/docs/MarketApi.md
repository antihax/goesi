# \MarketApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdOrders**](MarketApi.md#GetCharactersCharacterIdOrders) | **Get** /v2/characters/{character_id}/orders/ | List open orders from a character
[**GetCharactersCharacterIdOrdersHistory**](MarketApi.md#GetCharactersCharacterIdOrdersHistory) | **Get** /v1/characters/{character_id}/orders/history/ | List historical orders by a character
[**GetCorporationsCorporationIdOrders**](MarketApi.md#GetCorporationsCorporationIdOrders) | **Get** /v2/corporations/{corporation_id}/orders/ | List open orders from a corporation
[**GetCorporationsCorporationIdOrdersHistory**](MarketApi.md#GetCorporationsCorporationIdOrdersHistory) | **Get** /v1/corporations/{corporation_id}/orders/history/ | List historical orders from a corporation
[**GetMarketsGroups**](MarketApi.md#GetMarketsGroups) | **Get** /v1/markets/groups/ | Get item groups
[**GetMarketsGroupsMarketGroupId**](MarketApi.md#GetMarketsGroupsMarketGroupId) | **Get** /v1/markets/groups/{market_group_id}/ | Get item group information
[**GetMarketsPrices**](MarketApi.md#GetMarketsPrices) | **Get** /v1/markets/prices/ | List market prices
[**GetMarketsRegionIdHistory**](MarketApi.md#GetMarketsRegionIdHistory) | **Get** /v1/markets/{region_id}/history/ | List historical market statistics in a region
[**GetMarketsRegionIdOrders**](MarketApi.md#GetMarketsRegionIdOrders) | **Get** /v1/markets/{region_id}/orders/ | List orders in a region
[**GetMarketsRegionIdTypes**](MarketApi.md#GetMarketsRegionIdTypes) | **Get** /v1/markets/{region_id}/types/ | List type IDs relevant to a market
[**GetMarketsStructuresStructureId**](MarketApi.md#GetMarketsStructuresStructureId) | **Get** /v1/markets/structures/{structure_id}/ | List orders in a structure


# **GetCharactersCharacterIdOrders**
> []GetCharactersCharacterIdOrders200Ok GetCharactersCharacterIdOrders(ctx, characterId, optional)
List open orders from a character

List open market orders placed by a character  ---  This route is cached for up to 1200 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdOrders200Ok**](get_characters_character_id_orders_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdOrdersHistory**
> []GetCharactersCharacterIdOrdersHistory200Ok GetCharactersCharacterIdOrdersHistory(ctx, characterId, optional)
List historical orders by a character

List cancelled and expired market orders placed by a character up to 90 days in the past.  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdOrdersHistory200Ok**](get_characters_character_id_orders_history_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdOrders**
> []GetCorporationsCorporationIdOrders200Ok GetCorporationsCorporationIdOrders(ctx, corporationId, optional)
List open orders from a corporation

List open market orders placed on behalf of a corporation  ---  This route is cached for up to 1200 seconds  --- Requires one of the following EVE corporation role(s): Accountant, Trader  --- Warning: This route has an upgrade available.  --- [Diff of the upcoming changes](https://esi.evetech.net/diff/latest/dev/#GET-/corporations/{corporation_id}/orders/)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdOrders200Ok**](get_corporations_corporation_id_orders_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdOrdersHistory**
> []GetCorporationsCorporationIdOrdersHistory200Ok GetCorporationsCorporationIdOrdersHistory(ctx, corporationId, optional)
List historical orders from a corporation

List cancelled and expired market orders placed on behalf of a corporation up to 90 days in the past.  ---  This route is cached for up to 3600 seconds  --- Requires one of the following EVE corporation role(s): Accountant, Trader  --- Warning: This route has an upgrade available.  --- [Diff of the upcoming changes](https://esi.evetech.net/diff/latest/dev/#GET-/corporations/{corporation_id}/orders/history/)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdOrdersHistory200Ok**](get_corporations_corporation_id_orders_history_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsGroups**
> []int32 GetMarketsGroups(ctx, optional)
Get item groups

Get a list of item groups  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsGroupsMarketGroupId**
> GetMarketsGroupsMarketGroupIdOk GetMarketsGroupsMarketGroupId(ctx, marketGroupId, optional)
Get item group information

Get information on an item group  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **marketGroupId** | **int32**| An Eve item group ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketGroupId** | **int32**| An Eve item group ID | 
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**GetMarketsGroupsMarketGroupIdOk**](get_markets_groups_market_group_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsPrices**
> []GetMarketsPrices200Ok GetMarketsPrices(ctx, optional)
List market prices

Return a list of prices  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**[]GetMarketsPrices200Ok**](get_markets_prices_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsRegionIdHistory**
> []GetMarketsRegionIdHistory200Ok GetMarketsRegionIdHistory(ctx, regionId, typeId, optional)
List historical market statistics in a region

Return a list of historical market statistics for the specified type in a region  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
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
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**[]GetMarketsRegionIdHistory200Ok**](get_markets_region_id_history_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsRegionIdOrders**
> []GetMarketsRegionIdOrders200Ok GetMarketsRegionIdOrders(ctx, orderType, regionId, optional)
List orders in a region

Return a list of orders in a region  ---  This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
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
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]
 **typeId** | **int32**| Return orders only for this type | 

### Return type

[**[]GetMarketsRegionIdOrders200Ok**](get_markets_region_id_orders_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsRegionIdTypes**
> []int32 GetMarketsRegionIdTypes(ctx, regionId, optional)
List type IDs relevant to a market

Return a list of type IDs that have active orders in the region, for efficient market indexing.  ---  This route is cached for up to 600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **regionId** | **int32**| Return statistics in this region | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **int32**| Return statistics in this region | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsStructuresStructureId**
> []GetMarketsStructuresStructureId200Ok GetMarketsStructuresStructureId(ctx, structureId, optional)
List orders in a structure

Return all orders in a structure  ---  This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **structureId** | **int64**| Return orders in this structure | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **structureId** | **int64**| Return orders in this structure | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetMarketsStructuresStructureId200Ok**](get_markets_structures_structure_id_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

