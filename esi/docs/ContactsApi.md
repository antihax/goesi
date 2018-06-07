# \ContactsApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCharactersCharacterIdContacts**](ContactsApi.md#DeleteCharactersCharacterIdContacts) | **Delete** /v2/characters/{character_id}/contacts/ | Delete contacts
[**GetAlliancesAllianceIdContacts**](ContactsApi.md#GetAlliancesAllianceIdContacts) | **Get** /v2/alliances/{alliance_id}/contacts/ | Get alliance contacts
[**GetAlliancesAllianceIdContactsLabels**](ContactsApi.md#GetAlliancesAllianceIdContactsLabels) | **Get** /v1/alliances/{alliance_id}/contacts/labels/ | Get alliance contact labels
[**GetCharactersCharacterIdContacts**](ContactsApi.md#GetCharactersCharacterIdContacts) | **Get** /v2/characters/{character_id}/contacts/ | Get contacts
[**GetCharactersCharacterIdContactsLabels**](ContactsApi.md#GetCharactersCharacterIdContactsLabels) | **Get** /v1/characters/{character_id}/contacts/labels/ | Get contact labels
[**GetCorporationsCorporationIdContacts**](ContactsApi.md#GetCorporationsCorporationIdContacts) | **Get** /v2/corporations/{corporation_id}/contacts/ | Get corporation contacts
[**GetCorporationsCorporationIdContactsLabels**](ContactsApi.md#GetCorporationsCorporationIdContactsLabels) | **Get** /v1/corporations/{corporation_id}/contacts/labels/ | Get corporation contact labels
[**PostCharactersCharacterIdContacts**](ContactsApi.md#PostCharactersCharacterIdContacts) | **Post** /v2/characters/{character_id}/contacts/ | Add contacts
[**PutCharactersCharacterIdContacts**](ContactsApi.md#PutCharactersCharacterIdContacts) | **Put** /v2/characters/{character_id}/contacts/ | Edit contacts


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

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
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
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetAlliancesAllianceIdContacts200Ok**](get_alliances_alliance_id_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlliancesAllianceIdContactsLabels**
> []GetAlliancesAllianceIdContactsLabels200Ok GetAlliancesAllianceIdContactsLabels(ctx, allianceId, optional)
Get alliance contact labels

Return custom labels for an alliance's contacts  ---  This route is cached for up to 300 seconds

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
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetAlliancesAllianceIdContactsLabels200Ok**](get_alliances_alliance_id_contacts_labels_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
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
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdContacts200Ok**](get_characters_character_id_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdContactsLabels**
> []GetCharactersCharacterIdContactsLabels200Ok GetCharactersCharacterIdContactsLabels(ctx, characterId, optional)
Get contact labels

Return custom labels for a character's contacts  ---  This route is cached for up to 300 seconds

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

[**[]GetCharactersCharacterIdContactsLabels200Ok**](get_characters_character_id_contacts_labels_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
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
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdContacts200Ok**](get_corporations_corporation_id_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdContactsLabels**
> []GetCorporationsCorporationIdContactsLabels200Ok GetCorporationsCorporationIdContactsLabels(ctx, corporationId, optional)
Get corporation contact labels

Return custom labels for a corporation's contacts  ---  This route is cached for up to 300 seconds

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
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdContactsLabels200Ok**](get_corporations_corporation_id_contacts_labels_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
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
  **contactIds** | **[]int32**| A list of contacts | 
  **standing** | **float32**| Standing for the contact | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **contactIds** | **[]int32**| A list of contacts | 
 **standing** | **float32**| Standing for the contact | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **labelIds** | [**[]int64**](int64.md)| Add custom labels to the new contact | 
 **token** | **string**| Access token to use if unable to set a header | 
 **watched** | **bool**| Whether the contact should be watched, note this is only effective on characters | [default to false]

### Return type

**[]int32**

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
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
  **contactIds** | **[]int32**| A list of contacts | 
  **standing** | **float32**| Standing for the contact | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **contactIds** | **[]int32**| A list of contacts | 
 **standing** | **float32**| Standing for the contact | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **labelIds** | [**[]int64**](int64.md)| Add custom labels to the contact | 
 **token** | **string**| Access token to use if unable to set a header | 
 **watched** | **bool**| Whether the contact should be watched, note this is only effective on characters | [default to false]

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

