# GetCharactersCharacterIdBlueprints200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **int64** | Unique ID for this item. The ID of an item is stable if that item is not repackaged, stacked, detached from a stack, assembled, or otherwise altered. If an item is changed in one of these ways, then the ID will also change (see notes below). | [default to null]
**LocationFlag** | **string** | Indicates something about this item&#39;s storage location. The flag is used to differentiate between hangar divisions, drone bay, fitting location, and similar. | [default to null]
**LocationId** | **int64** | References a solar system, station or itemID if this blueprint is located within a container. If an itemID the Character - AssetList API must be queried to find the container using the itemID, from which the correct location of the Blueprint can be derived. | [default to null]
**MaterialEfficiency** | **int32** | Material Efficiency Level of the blueprint, can be any integer between 0 and 10. | [default to null]
**Quantity** | **int32** | Typically will be -1 or -2 designating a singleton item, where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (no activities performed on them yet). | [default to null]
**Runs** | **int32** | Number of runs remaining if the blueprint is a copy, -1 if it is an original. | [default to null]
**TimeEfficiency** | **int32** | Time Efficiency Level of the blueprint, can be any even integer between 0 and 20. | [default to null]
**TypeId** | **int32** | type_id integer | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


