# \ContactsApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCharactersCharacterIdContacts**](ContactsApi.md#DeleteCharactersCharacterIdContacts) | **Delete** /characters/{character_id}/contacts/ | Delete contacts


# **DeleteCharactersCharacterIdContacts**
> DeleteCharactersCharacterIdContacts(ctx, characterId, contactIds, optional)
Delete contacts

Bulk delete contacts  --- Alternate route: `/dev/characters/{character_id}/contacts/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
  **contactIds** | [**[]int32**](int32.md)| A list of contacts to delete | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **contactIds** | [**[]int32**](int32.md)| A list of contacts to delete | 
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

