# \InsuranceApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInsurancePrices**](InsuranceApi.md#GetInsurancePrices) | **Get** /insurance/prices/ | List insurance levels


# **GetInsurancePrices**
> []GetInsurancePrices200Ok GetInsurancePrices(optional)
List insurance levels

Return available insurance levels for all ship types  ---  Alternate route: `/legacy/insurance/prices/`  Alternate route: `/latest/insurance/prices/`  Alternate route: `/dev/insurance/prices/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**[]GetInsurancePrices200Ok**](get_insurance_prices_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

