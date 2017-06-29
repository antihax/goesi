# \CorporationApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCorporationsCorporationIdAlliancehistory**](CorporationApi.md#GetCorporationsCorporationIdAlliancehistory) | **Get** /corporations/{corporation_id}/alliancehistory/ | Get alliance history
[**GetCorporationsCorporationIdIcons**](CorporationApi.md#GetCorporationsCorporationIdIcons) | **Get** /corporations/{corporation_id}/icons/ | Get corporation icon
[**GetCorporationsCorporationIdRoles**](CorporationApi.md#GetCorporationsCorporationIdRoles) | **Get** /corporations/{corporation_id}/roles/ | Get corporation member roles
[**GetCorporationsCorporationIdStructures**](CorporationApi.md#GetCorporationsCorporationIdStructures) | **Get** /corporations/{corporation_id}/structures/ | Get corporation structures
[**GetCorporationsNames**](CorporationApi.md#GetCorporationsNames) | **Get** /corporations/names/ | Get corporation names
[**GetCorporationsNpccorps**](CorporationApi.md#GetCorporationsNpccorps) | **Get** /corporations/npccorps/ | Get npc corporations
[**PutCorporationsCorporationIdStructuresStructureId**](CorporationApi.md#PutCorporationsCorporationIdStructuresStructureId) | **Put** /corporations/{corporation_id}/structures/{structure_id}/ | Update structure vulnerability schedule


# **GetCorporationsCorporationIdAlliancehistory**
> []GetCorporationsCorporationIdAlliancehistory200Ok GetCorporationsCorporationIdAlliancehistory(corporationId, optional)
Get alliance history

Get a list of all the alliances a corporation has been a member of  --- Alternate route: `/legacy/corporations/{corporation_id}/alliancehistory/`  Alternate route: `/latest/corporations/{corporation_id}/alliancehistory/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsCorporationIdAlliancehistory200Ok**](get_corporations_corporation_id_alliancehistory_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdIcons**
> GetCorporationsCorporationIdIconsOk GetCorporationsCorporationIdIcons(corporationId, optional)
Get corporation icon

Get the icon urls for a corporation  --- Alternate route: `/legacy/corporations/{corporation_id}/icons/`  Alternate route: `/latest/corporations/{corporation_id}/icons/`  Alternate route: `/dev/corporations/{corporation_id}/icons/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetCorporationsCorporationIdIconsOk**](get_corporations_corporation_id_icons_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdRoles**
> []GetCorporationsCorporationIdRoles200Ok GetCorporationsCorporationIdRoles(ctx, corporationId, optional)
Get corporation member roles

Return the roles of all members if the character has the personnel manager role or any grantable role.  --- Alternate route: `/legacy/corporations/{corporation_id}/roles/`  Alternate route: `/latest/corporations/{corporation_id}/roles/`  Alternate route: `/dev/corporations/{corporation_id}/roles/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsCorporationIdRoles200Ok**](get_corporations_corporation_id_roles_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdStructures**
> []GetCorporationsCorporationIdStructures200Ok GetCorporationsCorporationIdStructures(ctx, corporationId, optional)
Get corporation structures

Get a list of corporation structures  --- Alternate route: `/legacy/corporations/{corporation_id}/structures/`  Alternate route: `/latest/corporations/{corporation_id}/structures/`  Alternate route: `/dev/corporations/{corporation_id}/structures/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **language** | **string**| Language to use in the response | [default to en-us]
 **page** | **float32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsCorporationIdStructures200Ok**](get_corporations_corporation_id_structures_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsNames**
> []GetCorporationsNames200Ok GetCorporationsNames(corporationIds, optional)
Get corporation names

Resolve a set of corporation IDs to corporation names  --- Alternate route: `/legacy/corporations/names/`  Alternate route: `/latest/corporations/names/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **corporationIds** | [**[]int64**](int64.md)| A comma separated list of corporation IDs | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationIds** | [**[]int64**](int64.md)| A comma separated list of corporation IDs | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsNames200Ok**](get_corporations_names_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsNpccorps**
> []int32 GetCorporationsNpccorps(optional)
Get npc corporations

Get a list of npc corporations  --- Alternate route: `/legacy/corporations/npccorps/`  Alternate route: `/latest/corporations/npccorps/`  Alternate route: `/dev/corporations/npccorps/`  --- This route expires daily at 11:05

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

# **PutCorporationsCorporationIdStructuresStructureId**
> PutCorporationsCorporationIdStructuresStructureId(ctx, newSchedule, structureId, corporationId, optional)
Update structure vulnerability schedule

Update the vulnerability window schedule of a corporation structure  --- Alternate route: `/legacy/corporations/{corporation_id}/structures/{structure_id}/`  Alternate route: `/latest/corporations/{corporation_id}/structures/{structure_id}/`  Alternate route: `/dev/corporations/{corporation_id}/structures/{structure_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **newSchedule** | [**[]PutCorporationsCorporationIdStructuresStructureIdNewSchedule**](put_corporations_corporation_id_structures_structure_id_new_schedule.md)| New vulnerability window schedule for the structure | 
  **structureId** | **int64**| A structure ID | 
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newSchedule** | [**[]PutCorporationsCorporationIdStructuresStructureIdNewSchedule**](put_corporations_corporation_id_structures_structure_id_new_schedule.md)| New vulnerability window schedule for the structure | 
 **structureId** | **int64**| A structure ID | 
 **corporationId** | **int32**| An EVE corporation ID | 
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

