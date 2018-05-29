# \DogmaApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDogmaAttributes**](DogmaApi.md#GetDogmaAttributes) | **Get** /v1/dogma/attributes/ | Get attributes
[**GetDogmaAttributesAttributeId**](DogmaApi.md#GetDogmaAttributesAttributeId) | **Get** /v1/dogma/attributes/{attribute_id}/ | Get attribute information
[**GetDogmaDynamicItemsTypeIdItemId**](DogmaApi.md#GetDogmaDynamicItemsTypeIdItemId) | **Get** /v1/dogma/dynamic/items/{type_id}/{item_id}/ | Get dynamic item information
[**GetDogmaEffects**](DogmaApi.md#GetDogmaEffects) | **Get** /v1/dogma/effects/ | Get effects
[**GetDogmaEffectsEffectId**](DogmaApi.md#GetDogmaEffectsEffectId) | **Get** /v2/dogma/effects/{effect_id}/ | Get effect information


# **GetDogmaAttributes**
> []int32 GetDogmaAttributes(ctx, optional)
Get attributes

Get a list of dogma attribute ids  ---  This route expires daily at 11:05

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

# **GetDogmaAttributesAttributeId**
> GetDogmaAttributesAttributeIdOk GetDogmaAttributesAttributeId(ctx, attributeId, optional)
Get attribute information

Get information on a dogma attribute  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **attributeId** | **int32**| A dogma attribute ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeId** | **int32**| A dogma attribute ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetDogmaAttributesAttributeIdOk**](get_dogma_attributes_attribute_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDogmaDynamicItemsTypeIdItemId**
> GetDogmaDynamicItemsTypeIdItemIdOk GetDogmaDynamicItemsTypeIdItemId(ctx, itemId, typeId, optional)
Get dynamic item information

Returns info about a dynamic item resulting from mutation with a mutaplasmid.  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **itemId** | **int64**| item_id integer | 
  **typeId** | **int32**| type_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **int64**| item_id integer | 
 **typeId** | **int32**| type_id integer | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetDogmaDynamicItemsTypeIdItemIdOk**](get_dogma_dynamic_items_type_id_item_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDogmaEffects**
> []int32 GetDogmaEffects(ctx, optional)
Get effects

Get a list of dogma effect ids  ---  This route expires daily at 11:05

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

# **GetDogmaEffectsEffectId**
> GetDogmaEffectsEffectIdOk GetDogmaEffectsEffectId(ctx, effectId, optional)
Get effect information

Get information on a dogma effect  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **effectId** | **int32**| A dogma effect ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **effectId** | **int32**| A dogma effect ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetDogmaEffectsEffectIdOk**](get_dogma_effects_effect_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

