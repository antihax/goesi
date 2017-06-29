# \FleetsApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFleetsFleetIdMembersMemberId**](FleetsApi.md#DeleteFleetsFleetIdMembersMemberId) | **Delete** /fleets/{fleet_id}/members/{member_id}/ | Kick fleet member
[**DeleteFleetsFleetIdSquadsSquadId**](FleetsApi.md#DeleteFleetsFleetIdSquadsSquadId) | **Delete** /fleets/{fleet_id}/squads/{squad_id}/ | Delete fleet squad
[**DeleteFleetsFleetIdWingsWingId**](FleetsApi.md#DeleteFleetsFleetIdWingsWingId) | **Delete** /fleets/{fleet_id}/wings/{wing_id}/ | Delete fleet wing
[**GetFleetsFleetId**](FleetsApi.md#GetFleetsFleetId) | **Get** /fleets/{fleet_id}/ | Get fleet information
[**GetFleetsFleetIdMembers**](FleetsApi.md#GetFleetsFleetIdMembers) | **Get** /fleets/{fleet_id}/members/ | Get fleet members
[**GetFleetsFleetIdWings**](FleetsApi.md#GetFleetsFleetIdWings) | **Get** /fleets/{fleet_id}/wings/ | Get fleet wings
[**PostFleetsFleetIdMembers**](FleetsApi.md#PostFleetsFleetIdMembers) | **Post** /fleets/{fleet_id}/members/ | Create fleet invitation
[**PostFleetsFleetIdWings**](FleetsApi.md#PostFleetsFleetIdWings) | **Post** /fleets/{fleet_id}/wings/ | Create fleet wing
[**PostFleetsFleetIdWingsWingIdSquads**](FleetsApi.md#PostFleetsFleetIdWingsWingIdSquads) | **Post** /fleets/{fleet_id}/wings/{wing_id}/squads/ | Create fleet squad
[**PutFleetsFleetId**](FleetsApi.md#PutFleetsFleetId) | **Put** /fleets/{fleet_id}/ | Update fleet
[**PutFleetsFleetIdMembersMemberId**](FleetsApi.md#PutFleetsFleetIdMembersMemberId) | **Put** /fleets/{fleet_id}/members/{member_id}/ | Move fleet member
[**PutFleetsFleetIdSquadsSquadId**](FleetsApi.md#PutFleetsFleetIdSquadsSquadId) | **Put** /fleets/{fleet_id}/squads/{squad_id}/ | Rename fleet squad
[**PutFleetsFleetIdWingsWingId**](FleetsApi.md#PutFleetsFleetIdWingsWingId) | **Put** /fleets/{fleet_id}/wings/{wing_id}/ | Rename fleet wing


# **DeleteFleetsFleetIdMembersMemberId**
> DeleteFleetsFleetIdMembersMemberId(ctx, fleetId, memberId, optional)
Kick fleet member

Kick a fleet member  --- Alternate route: `/legacy/fleets/{fleet_id}/members/{member_id}/`  Alternate route: `/latest/fleets/{fleet_id}/members/{member_id}/`  Alternate route: `/dev/fleets/{fleet_id}/members/{member_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
  **memberId** | **int32**| The character ID of a member in this fleet | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **memberId** | **int32**| The character ID of a member in this fleet | 
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

# **DeleteFleetsFleetIdSquadsSquadId**
> DeleteFleetsFleetIdSquadsSquadId(ctx, fleetId, squadId, optional)
Delete fleet squad

Delete a fleet squad, only empty squads can be deleted  --- Alternate route: `/legacy/fleets/{fleet_id}/squads/{squad_id}/`  Alternate route: `/latest/fleets/{fleet_id}/squads/{squad_id}/`  Alternate route: `/dev/fleets/{fleet_id}/squads/{squad_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
  **squadId** | **int64**| The squad to delete | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **squadId** | **int64**| The squad to delete | 
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

# **DeleteFleetsFleetIdWingsWingId**
> DeleteFleetsFleetIdWingsWingId(ctx, fleetId, wingId, optional)
Delete fleet wing

Delete a fleet wing, only empty wings can be deleted. The wing may contain squads, but the squads must be empty  --- Alternate route: `/legacy/fleets/{fleet_id}/wings/{wing_id}/`  Alternate route: `/latest/fleets/{fleet_id}/wings/{wing_id}/`  Alternate route: `/dev/fleets/{fleet_id}/wings/{wing_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
  **wingId** | **int64**| The wing to delete | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **wingId** | **int64**| The wing to delete | 
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

# **GetFleetsFleetId**
> GetFleetsFleetIdOk GetFleetsFleetId(ctx, fleetId, optional)
Get fleet information

Return details about a fleet  --- Alternate route: `/legacy/fleets/{fleet_id}/`  Alternate route: `/latest/fleets/{fleet_id}/`  Alternate route: `/dev/fleets/{fleet_id}/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetFleetsFleetIdOk**](get_fleets_fleet_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFleetsFleetIdMembers**
> []GetFleetsFleetIdMembers200Ok GetFleetsFleetIdMembers(ctx, fleetId, optional)
Get fleet members

Return information about fleet members  --- Alternate route: `/legacy/fleets/{fleet_id}/members/`  Alternate route: `/latest/fleets/{fleet_id}/members/`  Alternate route: `/dev/fleets/{fleet_id}/members/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **language** | **string**| Language to use in the response | [default to en-us]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetFleetsFleetIdMembers200Ok**](get_fleets_fleet_id_members_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFleetsFleetIdWings**
> []GetFleetsFleetIdWings200Ok GetFleetsFleetIdWings(ctx, fleetId, optional)
Get fleet wings

Return information about wings in a fleet  --- Alternate route: `/legacy/fleets/{fleet_id}/wings/`  Alternate route: `/latest/fleets/{fleet_id}/wings/`  Alternate route: `/dev/fleets/{fleet_id}/wings/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **language** | **string**| Language to use in the response | [default to en-us]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetFleetsFleetIdWings200Ok**](get_fleets_fleet_id_wings_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFleetsFleetIdMembers**
> PostFleetsFleetIdMembers(ctx, fleetId, invitation, optional)
Create fleet invitation

Invite a character into the fleet, if a character has a CSPA charge set, it is not possible to invite them to the fleet using ESI  --- Alternate route: `/legacy/fleets/{fleet_id}/members/`  Alternate route: `/latest/fleets/{fleet_id}/members/`  Alternate route: `/dev/fleets/{fleet_id}/members/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
  **invitation** | [**PostFleetsFleetIdMembersInvitation**](PostFleetsFleetIdMembersInvitation.md)| Details of the invitation | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **invitation** | [**PostFleetsFleetIdMembersInvitation**](PostFleetsFleetIdMembersInvitation.md)| Details of the invitation | 
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

# **PostFleetsFleetIdWings**
> PostFleetsFleetIdWingsCreated PostFleetsFleetIdWings(ctx, fleetId, optional)
Create fleet wing

Create a new wing in a fleet  --- Alternate route: `/legacy/fleets/{fleet_id}/wings/`  Alternate route: `/latest/fleets/{fleet_id}/wings/`  Alternate route: `/dev/fleets/{fleet_id}/wings/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**PostFleetsFleetIdWingsCreated**](post_fleets_fleet_id_wings_created.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFleetsFleetIdWingsWingIdSquads**
> PostFleetsFleetIdWingsWingIdSquadsCreated PostFleetsFleetIdWingsWingIdSquads(ctx, fleetId, wingId, optional)
Create fleet squad

Create a new squad in a fleet  --- Alternate route: `/legacy/fleets/{fleet_id}/wings/{wing_id}/squads/`  Alternate route: `/latest/fleets/{fleet_id}/wings/{wing_id}/squads/`  Alternate route: `/dev/fleets/{fleet_id}/wings/{wing_id}/squads/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
  **wingId** | **int64**| The wing_id to create squad in | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **wingId** | **int64**| The wing_id to create squad in | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**PostFleetsFleetIdWingsWingIdSquadsCreated**](post_fleets_fleet_id_wings_wing_id_squads_created.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFleetsFleetId**
> PutFleetsFleetId(ctx, fleetId, newSettings, optional)
Update fleet

Update settings about a fleet  --- Alternate route: `/legacy/fleets/{fleet_id}/`  Alternate route: `/latest/fleets/{fleet_id}/`  Alternate route: `/dev/fleets/{fleet_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
  **newSettings** | [**PutFleetsFleetIdNewSettings**](PutFleetsFleetIdNewSettings.md)| What to update for this fleet | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **newSettings** | [**PutFleetsFleetIdNewSettings**](PutFleetsFleetIdNewSettings.md)| What to update for this fleet | 
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

# **PutFleetsFleetIdMembersMemberId**
> PutFleetsFleetIdMembersMemberId(ctx, fleetId, memberId, movement, optional)
Move fleet member

Move a fleet member around  --- Alternate route: `/legacy/fleets/{fleet_id}/members/{member_id}/`  Alternate route: `/latest/fleets/{fleet_id}/members/{member_id}/`  Alternate route: `/dev/fleets/{fleet_id}/members/{member_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
  **memberId** | **int32**| The character ID of a member in this fleet | 
  **movement** | [**PutFleetsFleetIdMembersMemberIdMovement**](PutFleetsFleetIdMembersMemberIdMovement.md)| Details of the invitation | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **memberId** | **int32**| The character ID of a member in this fleet | 
 **movement** | [**PutFleetsFleetIdMembersMemberIdMovement**](PutFleetsFleetIdMembersMemberIdMovement.md)| Details of the invitation | 
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

# **PutFleetsFleetIdSquadsSquadId**
> PutFleetsFleetIdSquadsSquadId(ctx, fleetId, naming, squadId, optional)
Rename fleet squad

Rename a fleet squad  --- Alternate route: `/legacy/fleets/{fleet_id}/squads/{squad_id}/`  Alternate route: `/latest/fleets/{fleet_id}/squads/{squad_id}/`  Alternate route: `/dev/fleets/{fleet_id}/squads/{squad_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
  **naming** | [**PutFleetsFleetIdSquadsSquadIdNaming**](PutFleetsFleetIdSquadsSquadIdNaming.md)| New name of the squad | 
  **squadId** | **int64**| The squad to rename | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **naming** | [**PutFleetsFleetIdSquadsSquadIdNaming**](PutFleetsFleetIdSquadsSquadIdNaming.md)| New name of the squad | 
 **squadId** | **int64**| The squad to rename | 
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

# **PutFleetsFleetIdWingsWingId**
> PutFleetsFleetIdWingsWingId(ctx, fleetId, naming, wingId, optional)
Rename fleet wing

Rename a fleet wing  --- Alternate route: `/legacy/fleets/{fleet_id}/wings/{wing_id}/`  Alternate route: `/latest/fleets/{fleet_id}/wings/{wing_id}/`  Alternate route: `/dev/fleets/{fleet_id}/wings/{wing_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **fleetId** | **int64**| ID for a fleet | 
  **naming** | [**PutFleetsFleetIdWingsWingIdNaming**](PutFleetsFleetIdWingsWingIdNaming.md)| New name of the wing | 
  **wingId** | **int64**| The wing to rename | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **int64**| ID for a fleet | 
 **naming** | [**PutFleetsFleetIdWingsWingIdNaming**](PutFleetsFleetIdWingsWingIdNaming.md)| New name of the wing | 
 **wingId** | **int64**| The wing to rename | 
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

