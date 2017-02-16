# \SovereigntyApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSovereigntyCampaigns**](SovereigntyApi.md#GetSovereigntyCampaigns) | **Get** /sovereignty/campaigns/ | List sovereignty campaigns
[**GetSovereigntyStructures**](SovereigntyApi.md#GetSovereigntyStructures) | **Get** /sovereignty/structures/ | List sovereignty structures


# **GetSovereigntyCampaigns**
> []GetSovereigntyCampaigns200Ok GetSovereigntyCampaigns(optional)
List sovereignty campaigns

Shows sovereignty data for campaigns.  ---  Alternate route: `/legacy/sovereignty/campaigns/`  Alternate route: `/latest/sovereignty/campaigns/`  Alternate route: `/dev/sovereignty/campaigns/`   ---  This route is cached for up to 5 seconds

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

[**[]GetSovereigntyCampaigns200Ok**](get_sovereignty_campaigns_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSovereigntyStructures**
> []GetSovereigntyStructures200Ok GetSovereigntyStructures(optional)
List sovereignty structures

Shows sovereignty data for structures.  ---  Alternate route: `/legacy/sovereignty/structures/`  Alternate route: `/latest/sovereignty/structures/`  Alternate route: `/dev/sovereignty/structures/`   ---  This route is cached for up to 120 seconds

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

[**[]GetSovereigntyStructures200Ok**](get_sovereignty_structures_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

