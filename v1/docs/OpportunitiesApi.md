# \OpportunitiesApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdOpportunities**](OpportunitiesApi.md#GetCharactersCharacterIdOpportunities) | **Get** /characters/{character_id}/opportunities/ | Get a character&#39;s completed tasks
[**GetOpportunitiesGroups**](OpportunitiesApi.md#GetOpportunitiesGroups) | **Get** /opportunities/groups/ | Get opportunities groups
[**GetOpportunitiesGroupsGroupId**](OpportunitiesApi.md#GetOpportunitiesGroupsGroupId) | **Get** /opportunities/groups/{group_id}/ | Get opportunities group
[**GetOpportunitiesTasks**](OpportunitiesApi.md#GetOpportunitiesTasks) | **Get** /opportunities/tasks/ | Get opportunities tasks
[**GetOpportunitiesTasksTaskId**](OpportunitiesApi.md#GetOpportunitiesTasksTaskId) | **Get** /opportunities/tasks/{task_id}/ | Get opportunities task


# **GetCharactersCharacterIdOpportunities**
> []GetCharactersCharacterIdOpportunities200Ok GetCharactersCharacterIdOpportunities(ctx, characterId, optional)
Get a character's completed tasks

Return a list of tasks finished by a character  ---  Alternate route: `/legacy/characters/{character_id}/opportunities/`  Alternate route: `/latest/characters/{character_id}/opportunities/`  Alternate route: `/dev/characters/{character_id}/opportunities/`   ---  This route is cached for up to 3600 seconds

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
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdOpportunities200Ok**](get_characters_character_id_opportunities_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesGroups**
> []int32 GetOpportunitiesGroups(optional)
Get opportunities groups

Return a list of opportunities groups  ---  Alternate route: `/legacy/opportunities/groups/`  Alternate route: `/latest/opportunities/groups/`  Alternate route: `/dev/opportunities/groups/`   ---  This route is cached for up to 3600 seconds

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

# **GetOpportunitiesGroupsGroupId**
> GetOpportunitiesGroupsGroupIdOk GetOpportunitiesGroupsGroupId(groupId, optional)
Get opportunities group

Return information of an opportunities group  ---  Alternate route: `/legacy/opportunities/groups/{group_id}/`  Alternate route: `/latest/opportunities/groups/{group_id}/`  Alternate route: `/dev/opportunities/groups/{group_id}/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **groupId** | **int32**| ID of an opportunities group | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32**| ID of an opportunities group | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **language** | **string**| Language to use in the response | [default to en-us]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetOpportunitiesGroupsGroupIdOk**](get_opportunities_groups_group_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesTasks**
> []int32 GetOpportunitiesTasks(optional)
Get opportunities tasks

Return a list of opportunities tasks  ---  Alternate route: `/legacy/opportunities/tasks/`  Alternate route: `/latest/opportunities/tasks/`  Alternate route: `/dev/opportunities/tasks/`   ---  This route is cached for up to 3600 seconds

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

# **GetOpportunitiesTasksTaskId**
> GetOpportunitiesTasksTaskIdOk GetOpportunitiesTasksTaskId(taskId, optional)
Get opportunities task

Return information of an opportunities task  ---  Alternate route: `/legacy/opportunities/tasks/{task_id}/`  Alternate route: `/latest/opportunities/tasks/{task_id}/`  Alternate route: `/dev/opportunities/tasks/{task_id}/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **taskId** | **int32**| ID of an opportunities task | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **int32**| ID of an opportunities task | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetOpportunitiesTasksTaskIdOk**](get_opportunities_tasks_task_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

