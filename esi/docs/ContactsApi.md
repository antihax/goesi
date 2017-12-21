# \ContactsApi

All URIs are relative to *https://esi.tech.ccp.is*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCharactersCharacterIdContacts**](ContactsApi.md#DeleteCharactersCharacterIdContacts) | **Delete** /v2/characters/{character_id}/contacts/ | Delete contacts
[**GetAlliancesAllianceIdContacts**](ContactsApi.md#GetAlliancesAllianceIdContacts) | **Get** /v1/alliances/{alliance_id}/contacts/ | Get alliance contacts
[**GetCharactersCharacterIdContacts**](ContactsApi.md#GetCharactersCharacterIdContacts) | **Get** /v1/characters/{character_id}/contacts/ | Get contacts
[**GetCharactersCharacterIdContactsLabels**](ContactsApi.md#GetCharactersCharacterIdContactsLabels) | **Get** /v1/characters/{character_id}/contacts/labels/ | Get contact labels
[**GetCorporationsCorporationIdContacts**](ContactsApi.md#GetCorporationsCorporationIdContacts) | **Get** /v1/corporations/{corporation_id}/contacts/ | Get corporation contacts
[**PostCharactersCharacterIdContacts**](ContactsApi.md#PostCharactersCharacterIdContacts) | **Post** /v1/characters/{character_id}/contacts/ | Add contacts
[**PutCharactersCharacterIdContacts**](ContactsApi.md#PutCharactersCharacterIdContacts) | **Put** /v1/characters/{character_id}/contacts/ | Edit contacts


# **DeleteCharactersCharacterIdContacts**
> DeleteCharactersCharacterIdContacts(ctx, characterId, contactIds, optional)
Delete contacts

Bulk delete contacts  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
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

# **GetAlliancesAllianceIdContacts**
> []GetAlliancesAllianceIdContacts200Ok GetAlliancesAllianceIdContacts(ctx, allianceId, optional)
Get alliance contacts

Return contacts of an alliance  ---  This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **allianceId** | **int32**| An EVE alliance ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allianceId** | **int32**| An EVE alliance ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetAlliancesAllianceIdContacts200Ok**](get_alliances_alliance_id_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdContacts**
> []GetCharactersCharacterIdContacts200Ok GetCharactersCharacterIdContacts(ctx, characterId, optional)
Get contacts

Return contacts of a character  ---  This route is cached for up to 300 seconds

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
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdContacts200Ok**](get_characters_character_id_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdContactsLabels**
> []GetCharactersCharacterIdContactsLabels200Ok GetCharactersCharacterIdContactsLabels(ctx, characterId, optional)
Get contact labels

Return custom labels for contacts the character defined  ---  This route is cached for up to 300 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdContactsLabels200Ok**](get_characters_character_id_contacts_labels_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdContacts**
> []GetCorporationsCorporationIdContacts200Ok GetCorporationsCorporationIdContacts(ctx, corporationId, optional)
Get corporation contacts

Return contacts of a corporation  ---  This route is cached for up to 300 seconds

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
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsCorporationIdContacts200Ok**](get_corporations_corporation_id_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdContacts**
> []int32 PostCharactersCharacterIdContacts(ctx, characterId, contactIds, standing, optional)
Add contacts

Bulk add contacts with same settings  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **characterId** | **int32**| An EVE character ID | 
  **contactIds** | **[]int32**| A list of contacts to add | 
  **standing** | **float32**| Standing for the new contact | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **contactIds** | **[]int32**| A list of contacts to add | 
 **standing** | **float32**| Standing for the new contact | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **labelId** | **int64**| Add a custom label to the new contact | [default to 0]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **watched** | **bool**| Whether the new contact should be watched, note this is only effective on characters | [default to false]
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

**[]int32**

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCharactersCharacterIdContacts**
> PutCharactersCharacterIdContacts(ctx, characterId, contactIds, standing, optional)
Edit contacts

Bulk edit contacts with same settings  --- 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **characterId** | **int32**| An EVE character ID | 
  **contactIds** | **[]int32**| A list of contacts to edit | 
  **standing** | **float32**| Standing for the contact | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **contactIds** | **[]int32**| A list of contacts to edit | 
 **standing** | **float32**| Standing for the contact | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **labelId** | **int64**| Add a custom label to the contact, use 0 for clearing label | [default to 0]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **watched** | **bool**| Whether the contact should be watched, note this is only effective on characters | [default to false]
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

