# \CorporationApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCorporationsCorporationId**](CorporationApi.md#GetCorporationsCorporationId) | **Get** /corporations/{corporation_id}/ | Get corporation information
[**GetCorporationsCorporationIdAlliancehistory**](CorporationApi.md#GetCorporationsCorporationIdAlliancehistory) | **Get** /corporations/{corporation_id}/alliancehistory/ | Get alliance history
[**GetCorporationsCorporationIdIcons**](CorporationApi.md#GetCorporationsCorporationIdIcons) | **Get** /corporations/{corporation_id}/icons/ | Get corporation icon
[**GetCorporationsCorporationIdRoles**](CorporationApi.md#GetCorporationsCorporationIdRoles) | **Get** /corporations/{corporation_id}/roles/ | Get corporation member roles
[**GetCorporationsNames**](CorporationApi.md#GetCorporationsNames) | **Get** /corporations/names/ | Get corporation names


# **GetCorporationsCorporationId**
> GetCorporationsCorporationIdOk GetCorporationsCorporationId(corporationId, optional)
Get corporation information

Public information about a corporation  ---  Alternate route: `/legacy/corporations/{corporation_id}/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **corporationId** | **int32**| An Eve corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An Eve corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetCorporationsCorporationIdOk**](get_corporations_corporation_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdAlliancehistory**
> []GetCorporationsCorporationIdAlliancehistory200Ok GetCorporationsCorporationIdAlliancehistory(corporationId, optional)
Get alliance history

Get a list of all the alliances a corporation has been a member of  ---  Alternate route: `/legacy/corporations/{corporation_id}/alliancehistory/`  Alternate route: `/latest/corporations/{corporation_id}/alliancehistory/`  Alternate route: `/dev/corporations/{corporation_id}/alliancehistory/`   ---  This route is cached for up to 3600 seconds

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

Get the icon urls for a corporation  ---  Alternate route: `/legacy/corporations/{corporation_id}/icons/`  Alternate route: `/latest/corporations/{corporation_id}/icons/`  Alternate route: `/dev/corporations/{corporation_id}/icons/`   ---  This route is cached for up to 3600 seconds

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

Return the roles of all members if the character has the personnel manager role or any grantable role.  ---  Alternate route: `/legacy/corporations/{corporation_id}/roles/`  Alternate route: `/latest/corporations/{corporation_id}/roles/`  Alternate route: `/dev/corporations/{corporation_id}/roles/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **corporationId** | **int32**| A corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| A corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
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

# **GetCorporationsNames**
> []GetCorporationsNames200Ok GetCorporationsNames(corporationIds, optional)
Get corporation names

Resolve a set of corporation IDs to corporation names  ---  Alternate route: `/legacy/corporations/names/`  Alternate route: `/latest/corporations/names/`  Alternate route: `/dev/corporations/names/`   ---  This route is cached for up to 3600 seconds

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

