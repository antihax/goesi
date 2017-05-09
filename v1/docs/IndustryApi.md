# \IndustryApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdIndustryJobs**](IndustryApi.md#GetCharactersCharacterIdIndustryJobs) | **Get** /characters/{character_id}/industry/jobs/ | List character industry jobs
[**GetIndustryFacilities**](IndustryApi.md#GetIndustryFacilities) | **Get** /industry/facilities/ | List industry facilities
[**GetIndustrySystems**](IndustryApi.md#GetIndustrySystems) | **Get** /industry/systems/ | List solar system cost indices


# **GetCharactersCharacterIdIndustryJobs**
> []GetCharactersCharacterIdIndustryJobs200Ok GetCharactersCharacterIdIndustryJobs(ctx, characterId, optional)
List character industry jobs

List industry jobs placed by a character  ---  Alternate route: `/legacy/characters/{character_id}/industry/jobs/`  Alternate route: `/latest/characters/{character_id}/industry/jobs/`  Alternate route: `/dev/characters/{character_id}/industry/jobs/`   ---  This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **includeCompleted** | **bool**| Whether retrieve completed character industry jobs as well | 
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdIndustryJobs200Ok**](get_characters_character_id_industry_jobs_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIndustryFacilities**
> []GetIndustryFacilities200Ok GetIndustryFacilities(optional)
List industry facilities

Return a list of industry facilities  ---  Alternate route: `/legacy/industry/facilities/`  Alternate route: `/latest/industry/facilities/`  Alternate route: `/dev/industry/facilities/`   ---  This route is cached for up to 3600 seconds

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

[**[]GetIndustryFacilities200Ok**](get_industry_facilities_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIndustrySystems**
> []GetIndustrySystems200Ok GetIndustrySystems(optional)
List solar system cost indices

Return cost indices for solar systems  ---  Alternate route: `/legacy/industry/systems/`  Alternate route: `/latest/industry/systems/`  Alternate route: `/dev/industry/systems/`   ---  This route is cached for up to 3600 seconds

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

[**[]GetIndustrySystems200Ok**](get_industry_systems_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

