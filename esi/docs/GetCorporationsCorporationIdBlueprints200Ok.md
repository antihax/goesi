# GetCorporationsCorporationIdBlueprints200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **int64** | Unique ID for this item. | [default to null]
**LocationFlag** | **string** | Type of the location_id | [default to null]
**LocationId** | **int64** | References a station, a ship or an item_id if this blueprint is located within a container. | [default to null]
**MaterialEfficiency** | **int32** | Material Efficiency Level of the blueprint. | [default to null]
**Quantity** | **int32** | A range of numbers with a minimum of -2 and no maximum value where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (e.g. no activities performed on them yet). | [default to null]
**Runs** | **int32** | Number of runs remaining if the blueprint is a copy, -1 if it is an original. | [default to null]
**TimeEfficiency** | **int32** | Time Efficiency Level of the blueprint. | [default to null]
**TypeId** | **int32** | type_id integer | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


