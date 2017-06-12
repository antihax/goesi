# \DogmaApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDogmaAttributes**](DogmaApi.md#GetDogmaAttributes) | **Get** /dogma/attributes/ | Get attributes
[**GetDogmaAttributesAttributeId**](DogmaApi.md#GetDogmaAttributesAttributeId) | **Get** /dogma/attributes/{attribute_id}/ | Get attribute information
[**GetDogmaEffects**](DogmaApi.md#GetDogmaEffects) | **Get** /dogma/effects/ | Get effects
[**GetDogmaEffectsEffectId**](DogmaApi.md#GetDogmaEffectsEffectId) | **Get** /dogma/effects/{effect_id}/ | Get effect information


# **GetDogmaAttributes**
> []int32 GetDogmaAttributes(optional)
Get attributes

Get a list of dogma attribute ids  ---  Alternate route: `/legacy/dogma/attributes/`  Alternate route: `/latest/dogma/attributes/`  Alternate route: `/dev/dogma/attributes/`   ---  This route expires daily at 11:05

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

# **GetDogmaAttributesAttributeId**
> GetDogmaAttributesAttributeIdOk GetDogmaAttributesAttributeId(attributeId, optional)
Get attribute information

Get information on a dogma attribute  ---  Alternate route: `/legacy/dogma/attributes/{attribute_id}/`  Alternate route: `/latest/dogma/attributes/{attribute_id}/`  Alternate route: `/dev/dogma/attributes/{attribute_id}/`   ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **attributeId** | **int32**| A dogma attribute ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeId** | **int32**| A dogma attribute ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetDogmaAttributesAttributeIdOk**](get_dogma_attributes_attribute_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDogmaEffects**
> []int32 GetDogmaEffects(optional)
Get effects

Get a list of dogma effect ids  ---  Alternate route: `/legacy/dogma/effects/`  Alternate route: `/latest/dogma/effects/`  Alternate route: `/dev/dogma/effects/`   ---  This route expires daily at 11:05

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

# **GetDogmaEffectsEffectId**
> GetDogmaEffectsEffectIdOk GetDogmaEffectsEffectId(effectId, optional)
Get effect information

Get information on a dogma effect  ---  Alternate route: `/legacy/dogma/effects/{effect_id}/`  Alternate route: `/latest/dogma/effects/{effect_id}/`   ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **effectId** | **int32**| A dogma effect ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **effectId** | **int32**| A dogma effect ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetDogmaEffectsEffectIdOk**](get_dogma_effects_effect_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

