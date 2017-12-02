# \IndustryApi

All URIs are relative to *https://esi.tech.ccp.is*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdIndustryJobs**](IndustryApi.md#GetCharactersCharacterIdIndustryJobs) | **Get** /v1/characters/{character_id}/industry/jobs/ | List character industry jobs
[**GetCharactersCharacterIdMining**](IndustryApi.md#GetCharactersCharacterIdMining) | **Get** /v1/characters/{character_id}/mining/ | Character mining ledger
[**GetCorporationCorporationIdMiningExtractions**](IndustryApi.md#GetCorporationCorporationIdMiningExtractions) | **Get** /v1/corporation/{corporation_id}/mining/extractions/ | Moon extraction timers
[**GetCorporationCorporationIdMiningObservers**](IndustryApi.md#GetCorporationCorporationIdMiningObservers) | **Get** /v1/corporation/{corporation_id}/mining/observers/ | Corporation mining observers
[**GetCorporationCorporationIdMiningObserversObserverId**](IndustryApi.md#GetCorporationCorporationIdMiningObserversObserverId) | **Get** /v1/corporation/{corporation_id}/mining/observers/{observer_id}/ | Observed corporation mining
[**GetCorporationsCorporationIdIndustryJobs**](IndustryApi.md#GetCorporationsCorporationIdIndustryJobs) | **Get** /v1/corporations/{corporation_id}/industry/jobs/ | List corporation industry jobs
[**GetIndustryFacilities**](IndustryApi.md#GetIndustryFacilities) | **Get** /v1/industry/facilities/ | List industry facilities
[**GetIndustrySystems**](IndustryApi.md#GetIndustrySystems) | **Get** /v1/industry/systems/ | List solar system cost indices


# **GetCharactersCharacterIdIndustryJobs**
> []GetCharactersCharacterIdIndustryJobs200Ok GetCharactersCharacterIdIndustryJobs(ctx, characterId, optional)
List character industry jobs

List industry jobs placed by a character  ---  This route is cached for up to 300 seconds

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
 **includeCompleted** | **bool**| Whether retrieve completed character industry jobs as well | 
 **token** | **string**| Access token to use if unable to set a header | 
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

# **GetCharactersCharacterIdMining**
> []GetCharactersCharacterIdMining200Ok GetCharactersCharacterIdMining(ctx, characterId, optional)
Character mining ledger

Paginated record of all mining done by a character for the past 30 days  ---  This route is cached for up to 600 seconds

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

[**[]GetCharactersCharacterIdMining200Ok**](get_characters_character_id_mining_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationCorporationIdMiningExtractions**
> []GetCorporationCorporationIdMiningExtractions200Ok GetCorporationCorporationIdMiningExtractions(ctx, corporationId, optional)
Moon extraction timers

Extraction timers for all moon chunks being extracted by refineries belonging to a corporation.  ---  This route is cached for up to 1800 seconds  --- Requires one of the following EVE corporation role(s): Structure_manager

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
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationCorporationIdMiningExtractions200Ok**](get_corporation_corporation_id_mining_extractions_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationCorporationIdMiningObservers**
> []GetCorporationCorporationIdMiningObservers200Ok GetCorporationCorporationIdMiningObservers(ctx, corporationId, optional)
Corporation mining observers

Paginated list of all entities capable of observing and recording mining for a corporation  ---  This route is cached for up to 3600 seconds  --- Requires one of the following EVE corporation role(s): Accountant

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

[**[]GetCorporationCorporationIdMiningObservers200Ok**](get_corporation_corporation_id_mining_observers_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationCorporationIdMiningObserversObserverId**
> []GetCorporationCorporationIdMiningObserversObserverId200Ok GetCorporationCorporationIdMiningObserversObserverId(ctx, corporationId, observerId, optional)
Observed corporation mining

Paginated record of all mining seen by an observer  ---  This route is cached for up to 3600 seconds  --- Requires one of the following EVE corporation role(s): Accountant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | pass through context (authentication, logging, tracing)
  **corporationId** | **int32**| An EVE corporation ID | 
  **observerId** | **int64**| A mining observer id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **observerId** | **int64**| A mining observer id | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationCorporationIdMiningObserversObserverId200Ok**](get_corporation_corporation_id_mining_observers_observer_id_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdIndustryJobs**
> []GetCorporationsCorporationIdIndustryJobs200Ok GetCorporationsCorporationIdIndustryJobs(ctx, corporationId, optional)
List corporation industry jobs

List industry jobs run by a corporation  ---  This route is cached for up to 300 seconds  --- Requires one of the following EVE corporation role(s): FactoryManager

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
 **includeCompleted** | **bool**| Whether retrieve completed industry jobs as well | [default to false]
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCorporationsCorporationIdIndustryJobs200Ok**](get_corporations_corporation_id_industry_jobs_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIndustryFacilities**
> []GetIndustryFacilities200Ok GetIndustryFacilities(ctx, optional)
List industry facilities

Return a list of industry facilities  ---  This route is cached for up to 3600 seconds

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
> []GetIndustrySystems200Ok GetIndustrySystems(ctx, optional)
List solar system cost indices

Return cost indices for solar systems  ---  This route is cached for up to 3600 seconds

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

