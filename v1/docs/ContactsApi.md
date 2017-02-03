# \ContactsApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCharactersCharacterIdContacts**](ContactsApi.md#DeleteCharactersCharacterIdContacts) | **Delete** /characters/{character_id}/contacts/ | Delete contacts
[**GetCharactersCharacterIdContacts**](ContactsApi.md#GetCharactersCharacterIdContacts) | **Get** /characters/{character_id}/contacts/ | Get contacts
[**GetCharactersCharacterIdContactsLabels**](ContactsApi.md#GetCharactersCharacterIdContactsLabels) | **Get** /characters/{character_id}/contacts/labels/ | Get contact labels
[**PostCharactersCharacterIdContacts**](ContactsApi.md#PostCharactersCharacterIdContacts) | **Post** /characters/{character_id}/contacts/ | Add contacts
[**PutCharactersCharacterIdContacts**](ContactsApi.md#PutCharactersCharacterIdContacts) | **Put** /characters/{character_id}/contacts/ | Edit contacts


# **DeleteCharactersCharacterIdContacts**
> DeleteCharactersCharacterIdContacts(ctx, characterId, contactIds, optional)
Delete contacts

Bulk delete contacts  ---  Alternate route: `/legacy/characters/{character_id}/contacts/`  Alternate route: `/latest/characters/{character_id}/contacts/`  Alternate route: `/dev/characters/{character_id}/contacts/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| ID for a character | 
  **contactIds** | **[]int32**| A list of contacts to edit | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| ID for a character | 
 **contactIds** | **[]int32**| A list of contacts to edit | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdContacts**
> []GetCharactersCharacterIdContacts200Ok GetCharactersCharacterIdContacts(ctx, characterId, optional)
Get contacts

Return contacts of a character  ---  Alternate route: `/legacy/characters/{character_id}/contacts/`  Alternate route: `/latest/characters/{character_id}/contacts/`  Alternate route: `/dev/characters/{character_id}/contacts/`   ---  This route is cached for up to 300 seconds

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
 **page** | **int32**| page integer | [default to 1]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

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

Return custom labels for contacts the character defined  ---  Alternate route: `/legacy/characters/{character_id}/contacts/labels/`  Alternate route: `/latest/characters/{character_id}/contacts/labels/`  Alternate route: `/dev/characters/{character_id}/contacts/labels/`   ---  This route is cached for up to 300 seconds

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

### Return type

[**[]GetCharactersCharacterIdContactsLabels200Ok**](get_characters_character_id_contacts_labels_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdContacts**
> []int32 PostCharactersCharacterIdContacts(ctx, characterId, standing, contactIds, optional)
Add contacts

Bulk add contacts with same settings  ---  Alternate route: `/legacy/characters/{character_id}/contacts/`  Alternate route: `/latest/characters/{character_id}/contacts/`  Alternate route: `/dev/characters/{character_id}/contacts/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| ID for a character | 
  **standing** | **float32**| Standing for the new contact | 
  **contactIds** | **[]int32**| A list of contacts to add | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| ID for a character | 
 **standing** | **float32**| Standing for the new contact | 
 **contactIds** | **[]int32**| A list of contacts to add | 
 **watched** | **bool**| Whether the new contact should be watched, note this is only effective on characters | [default to false]
 **labelId** | **int64**| Add a custom label to the new contact | [default to 0]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

**[]int32**

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCharactersCharacterIdContacts**
> PutCharactersCharacterIdContacts(ctx, characterId, standing, contactIds, optional)
Edit contacts

Bulk edit contacts with same settings  ---  Alternate route: `/legacy/characters/{character_id}/contacts/`  Alternate route: `/latest/characters/{character_id}/contacts/`  Alternate route: `/dev/characters/{character_id}/contacts/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| ID for a character | 
  **standing** | **float32**| Standing for the contact | 
  **contactIds** | **[]int32**| A list of contacts to edit | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| ID for a character | 
 **standing** | **float32**| Standing for the contact | 
 **contactIds** | **[]int32**| A list of contacts to edit | 
 **watched** | **bool**| Whether the contact should be watched, note this is only effective on characters | [default to false]
 **labelId** | **int64**| Add a custom label to the contact, use 0 for clearing label | [default to 0]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

