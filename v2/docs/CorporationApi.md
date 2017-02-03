# \CorporationApi

All URIs are relative to *https://esi.tech.ccp.is/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCorporationsCorporationId**](CorporationApi.md#GetCorporationsCorporationId) | **Get** /corporations/{corporation_id}/ | Get corporation information
[**GetCorporationsCorporationIdMembers**](CorporationApi.md#GetCorporationsCorporationIdMembers) | **Get** /corporations/{corporation_id}/members/ | Get corporation members


# **GetCorporationsCorporationId**
> GetCorporationsCorporationIdOk GetCorporationsCorporationId(corporationId, optional)
Get corporation information

Public information about a corporation  ---  Alternate route: `/latest/corporations/{corporation_id}/`   ---  This route is cached for up to 3600 seconds

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

### Return type

[**GetCorporationsCorporationIdOk**](get_corporations_corporation_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdMembers**
> []GetCorporationsCorporationIdMembers200Ok GetCorporationsCorporationIdMembers(ctx, corporationId, optional)
Get corporation members

Read the current list of members if the calling character is a member.  ---  Alternate route: `/legacy/corporations/{corporation_id}/members/`  Alternate route: `/latest/corporations/{corporation_id}/members/`  Alternate route: `/dev/corporations/{corporation_id}/members/`   ---  This route is cached for up to 3600 seconds

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

### Return type

[**[]GetCorporationsCorporationIdMembers200Ok**](get_corporations_corporation_id_members_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

