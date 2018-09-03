# GetStatusItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | ESI Endpoint cluster advertising this route | [default to null]
**Method** | **string** | Swagger defined method | [default to null]
**Route** | **string** | Swagger defined route, not including version prefix | [default to null]
**Status** | **string** | Vague route status. Green is good, yellow is degraded, meaning slow or potentially dropping requests. Red means most requests are not succeeding and/or are very slow (5s+) on average. | [default to null]
**Tags** | **[]string** | Swagger tags applicable to this route | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


