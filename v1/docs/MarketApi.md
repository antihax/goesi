# \MarketApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketsPrices**](MarketApi.md#GetMarketsPrices) | **Get** /markets/prices/ | List market prices
[**GetMarketsRegionIdHistory**](MarketApi.md#GetMarketsRegionIdHistory) | **Get** /markets/{region_id}/history/ | List historical market statistics in a region
[**GetMarketsRegionIdOrders**](MarketApi.md#GetMarketsRegionIdOrders) | **Get** /markets/{region_id}/orders/ | List orders in a region
[**GetMarketsStructuresStructureId**](MarketApi.md#GetMarketsStructuresStructureId) | **Get** /markets/structures/{structure_id}/ | List orders in a structure


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

### Return type

[**[]GetMarketsRegionIdHistory200Ok**](get_markets_region_id_history_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketsRegionIdOrders**
> []GetMarketsRegionIdOrders200Ok GetMarketsRegionIdOrders(regionId, orderType, optional)
List orders in a region

Return a list of orders in a region  ---  Alternate route: `/legacy/markets/{region_id}/orders/`  Alternate route: `/latest/markets/{region_id}/orders/`  Alternate route: `/dev/markets/{region_id}/orders/`   ---  This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **regionId** | **int32**| Return orders in this region | 
  **orderType** | **string**| Filter buy/sell orders, return all orders by default. If you query without type_id, we always return both buy and sell orders.  | [default to all]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **int32**| Return orders in this region | 
 **orderType** | **string**| Filter buy/sell orders, return all orders by default. If you query without type_id, we always return both buy and sell orders.  | [default to all]
 **typeId** | **int32**| Return orders only for this type | 
 **page** | **int32**| Which page to query, only used for querying without type_id. Starting at 1  | [default to 1]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

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
 **page** | **int32**| Which page to query, starting at 1 | [default to 1]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**[]GetMarketsStructuresStructureId200Ok**](get_markets_structures_structure_id_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

