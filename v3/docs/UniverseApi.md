# \UniverseApi

All URIs are relative to *https://esi.tech.ccp.is/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniverseTypesTypeId**](UniverseApi.md#GetUniverseTypesTypeId) | **Get** /universe/types/{type_id}/ | Get type information


# **GetUniverseTypesTypeId**
> GetUniverseTypesTypeIdOk GetUniverseTypesTypeId(typeId, optional)
Get type information

Get information on a type  ---  Alternate route: `/dev/universe/types/{type_id}/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **typeId** | **int32**| An Eve item type ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeId** | **int32**| An Eve item type ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **language** | **string**| Language to use in the response | [default to en-us]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetUniverseTypesTypeIdOk**](get_universe_types_type_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

