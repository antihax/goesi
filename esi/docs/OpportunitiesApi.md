# \OpportunitiesApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdOpportunities**](OpportunitiesApi.md#GetCharactersCharacterIdOpportunities) | **Get** /v1/characters/{character_id}/opportunities/ | Get a character&#39;s completed tasks
[**GetOpportunitiesGroups**](OpportunitiesApi.md#GetOpportunitiesGroups) | **Get** /v1/opportunities/groups/ | Get opportunities groups
[**GetOpportunitiesGroupsGroupId**](OpportunitiesApi.md#GetOpportunitiesGroupsGroupId) | **Get** /v1/opportunities/groups/{group_id}/ | Get opportunities group
[**GetOpportunitiesTasks**](OpportunitiesApi.md#GetOpportunitiesTasks) | **Get** /v1/opportunities/tasks/ | Get opportunities tasks
[**GetOpportunitiesTasksTaskId**](OpportunitiesApi.md#GetOpportunitiesTasksTaskId) | **Get** /v1/opportunities/tasks/{task_id}/ | Get opportunities task


# **GetCharactersCharacterIdOpportunities**
> []GetCharactersCharacterIdOpportunities200Ok GetCharactersCharacterIdOpportunities(ctx, characterId, optional)
Get a character's completed tasks

Return a list of tasks finished by a character  ---  This route is cached for up to 3600 seconds

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

[**[]GetCharactersCharacterIdOpportunities200Ok**](get_characters_character_id_opportunities_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesGroups**
> []int32 GetOpportunitiesGroups(ctx, optional)
Get opportunities groups

Return a list of opportunities groups  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesGroupsGroupId**
> GetOpportunitiesGroupsGroupIdOk GetOpportunitiesGroupsGroupId(ctx, groupId, optional)
Get opportunities group

Return information of an opportunities group  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **groupId** | **int32**| ID of an opportunities group | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32**| ID of an opportunities group | 
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**GetOpportunitiesGroupsGroupIdOk**](get_opportunities_groups_group_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesTasks**
> []int32 GetOpportunitiesTasks(ctx, optional)
Get opportunities tasks

Return a list of opportunities tasks  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesTasksTaskId**
> GetOpportunitiesTasksTaskIdOk GetOpportunitiesTasksTaskId(ctx, taskId, optional)
Get opportunities task

Return information of an opportunities task  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **taskId** | **int32**| ID of an opportunities task | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **int32**| ID of an opportunities task | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetOpportunitiesTasksTaskIdOk**](get_opportunities_tasks_task_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

